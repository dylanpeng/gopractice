package main

import (
	"bytes"
	"fmt"
	"github.com/skip2/go-qrcode"
	"image"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
)

func main() {
	backImg, err := GetImageFromFile("./Group1171278146.png")
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
