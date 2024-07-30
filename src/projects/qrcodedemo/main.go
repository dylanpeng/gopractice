package main

import (
	"bytes"
	"fmt"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/skip2/go-qrcode"
	"golang.org/x/image/font"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

func main() {
	backImg, err := GetImageFromNet("https://static.awanptest.com/pint-intl-test/image-normal/20240730085002-ibvXU.png")
	if err != nil {
		fmt.Printf("GetImageFromFile fail. err:%s", err)
		return
	}
	qrcodeByte, err := qrcode.Encode("http://www.codesuger.com/", qrcode.Medium, 256)
	if err != nil {
		fmt.Printf("GetImageFromFile fail. err:%s", err)
		return
	}
	qrcodeBuffer := bytes.NewBuffer(qrcodeByte)
	qrImg, _, err := image.Decode(qrcodeBuffer)
	if err != nil {
		fmt.Printf("GetImageFromFile fail. err:%s", err)
		return
	}
	newImage, err := MergeImageNew(backImg, qrImg, 67, 67)
	if err != nil {
		fmt.Printf("GetImageFromFile fail. err:%s", err)
		return
	}
	f, err := os.Create("./newimage.png")
	if err != nil {
		fmt.Printf("Create fail. err:%s", err)
		return
	}
	err = png.Encode(f, newImage)
	if err != nil {
		fmt.Printf("Encode fail. err:%s", err)
		return
	}
	//qrcode.Encode("http://www.codesuger.com/", qrcode.Medium, 256)
	////qrcode.WriteFile("http://www.codesuger.com/", qrcode.Medium, 256, "./blog_qrcode.png")
}

// 图片拼接
func MergeImageNew(base image.Image, mask image.Image, paddingX int, paddingY int) (*image.RGBA, error) {
	baseSrcBounds := base.Bounds().Max

	maskSrcBounds := mask.Bounds().Max

	newWidth := baseSrcBounds.X
	newHeight := baseSrcBounds.Y

	maskWidth := maskSrcBounds.X
	maskHeight := maskSrcBounds.Y

	des := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight)) // 底板
	//首先将一个图片信息存入jpg
	draw.Draw(des, des.Bounds(), base, base.Bounds().Min, draw.Over)
	//将另外一张图片信息存入jpg
	draw.Draw(des, image.Rect(paddingX, newHeight-paddingY-maskHeight, (paddingX+maskWidth), (newHeight-paddingY)), mask, image.ZP, draw.Over)

	return des, nil
}

func GetImageFromFile(filePath string) (img image.Image, err error) {
	f1Src, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}
	defer f1Src.Close()

	buff := make([]byte, 512) // why 512 bytes ? see http://golang.org/pkg/net/http/#DetectContentType
	_, err = f1Src.Read(buff)

	if err != nil {
		return nil, err
	}

	filetype := http.DetectContentType(buff)

	fmt.Println(filetype)

	fSrc, err := os.Open(filePath)
	defer fSrc.Close()

	switch filetype {
	case "image/jpeg", "image/jpg":
		img, err = jpeg.Decode(fSrc)
		if err != nil {
			fmt.Println("jpeg error")
			return nil, err
		}

	case "image/gif":
		img, err = gif.Decode(fSrc)
		if err != nil {
			return nil, err
		}

	case "image/png":
		img, err = png.Decode(fSrc)
		if err != nil {
			return nil, err
		}
	default:
		return nil, err
	}
	return img, nil
}

func GetImageFromNet(url string) (image.Image, error) {
	res, err := http.Get(url)
	if err != nil || res.StatusCode != 200 {
		return nil, err
	}
	defer res.Body.Close()
	m, _, err := image.Decode(res.Body)
	return m, err
}

// 字体相关
type TextBrush struct {
	FontType  *truetype.Font
	FontSize  float64
	FontColor *image.Uniform
	TextWidth int
}

func NewTextBrush(FontFilePath string, FontSize float64, FontColor *image.Uniform, textWidth int) (*TextBrush, error) {
	fontFile, err := ioutil.ReadFile(FontFilePath)
	if err != nil {
		return nil, err
	}
	fontType, err := truetype.Parse(fontFile)
	if err != nil {
		return nil, err
	}
	if textWidth <= 0 {
		textWidth = 20
	}
	return &TextBrush{FontType: fontType, FontSize: FontSize, FontColor: FontColor, TextWidth: textWidth}, nil
}

// 图片插入文字
func (fb *TextBrush) DrawFontOnRGBA(rgba *image.RGBA, pt image.Point, content string) {

	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(fb.FontType)
	c.SetHinting(font.HintingFull)
	c.SetFontSize(fb.FontSize)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fb.FontColor)

	c.DrawString(content, freetype.Pt(pt.X, pt.Y))

}

func Image2RGBA(img image.Image) *image.RGBA {

	baseSrcBounds := img.Bounds().Max

	newWidth := baseSrcBounds.X
	newHeight := baseSrcBounds.Y

	des := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight)) // 底板
	//首先将一个图片信息存入jpg
	draw.Draw(des, des.Bounds(), img, img.Bounds().Min, draw.Over)

	return des
}

func SaveImage(targetPath string, m image.Image) error {
	fSave, err := os.Create(targetPath)
	if err != nil {
		return err
	}
	defer fSave.Close()

	err = jpeg.Encode(fSave, m, nil)

	if err != nil {
		return err
	}

	return nil
}

func TestTextBrush_DrawFontOnRGBA(t *testing.T) {
	textBrush, err := NewTextBrush("字体库ttf位置", 20, image.Black, 20)
	if err != nil {
		t.Log(err)
	}

	backgroud, err := GetImageFromFile("./resource/backgroud.jpg")
	if err != nil {
		t.Log(err)
	}
	des := Image2RGBA(backgroud)
	textBrush.DrawFontOnRGBA(des, image.Pt(10, 50), "世界你好")

	//调整颜色
	textBrush.FontColor = image.NewUniform(color.RGBA{
		R: 0x8E,
		G: 0xE5,
		B: 0xEE,
		A: 255,
	})

	textBrush.DrawFontOnRGBA(des, image.Pt(10, 80), "我是用Go拼上的文字")

	if err := SaveImage("./resource/text.png", des); err != nil {
		t.Log(err)
	}
}
