package im

import (
	"bytes"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"sort"
	"strings"
)

// @title 创建RSA的公钥私钥
// @param bits 值是1024或2048
func GenRsaKey(bits int) (string, string, error) {
	priKey, err2 := rsa.GenerateKey(rand.Reader, bits)
	if err2 != nil {
		panic(err2)
	}
	derStream := x509.MarshalPKCS1PrivateKey(priKey)
	block := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: derStream,
	}
	prvKey := pem.EncodeToMemory(block)
	puKey := &priKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(puKey)
	if err != nil {
		panic(err)
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	pubKey := pem.EncodeToMemory(block)
	privateKey := string(prvKey)
	publicKey := string(pubKey)
	return privateKey, publicKey, nil
}

// @title RSA签名
// @param originalData 签名前的原始字符串
// @param privateKey RSA私钥
func SignBase64(originalData, privateKey string) (string, error) {
	block, _ := pem.Decode([]byte(privateKey))
	// priKey, parseErr := x509.ParsePKCS8PrivateKey(block.Bytes)
	priKey, parseErr := x509.ParsePKCS1PrivateKey(block.Bytes)
	if parseErr != nil {
		return "", errors.New("parse privete key error")
	}
	// sha256 加密方式，必须与 下面的 crypto.SHA256 对应
	hash := sha256.New()
	hash.Write([]byte(originalData))
	// signature, err := rsa.SignPSS(rand.Reader, priKey.(*rsa.PrivateKey), crypto.SHA256, hash.Sum(nil), nil)
	signature, err := rsa.SignPKCS1v15(rand.Reader, priKey, crypto.SHA256, hash.Sum(nil))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signature), err
}

// @title RSA验签
// @param originalData 签名前的原始字符串
// @param signData 签名串,Base64格式的签名串
// @param pubKey 公钥,需与加密时使用的私钥相对应
// @return true代表验签通过，false为不通过
func VerifySignWithBase64(originalData, signData, pubKey string) (bool, error) {
	sign, err := base64.StdEncoding.DecodeString(signData)
	if err != nil {
		return false, err
	}
	block, _ := pem.Decode([]byte(pubKey))
	pub, err1 := x509.ParsePKIXPublicKey(block.Bytes)
	if err1 != nil {
		return false, err1
	}
	// sha256 加密方式，必须与 下面的 crypto.SHA256 对应
	hash := sha256.New()
	hash.Write([]byte(originalData))
	verifyErr := rsa.VerifyPKCS1v15(pub.(*rsa.PublicKey), crypto.SHA256, hash.Sum(nil), sign)
	return verifyErr == nil, nil
}

// @title AES加密
// @param origDatastr 原始字符串
// @param keystring AES的密钥key
func AesEncrypt(origDatastr string, keystring string) (string, error) {
	key := []byte(keystring)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	//AES分组长度为128位，所以blockSize=16，单位字节
	blockSize := block.BlockSize()
	origData := PKCS5Padding([]byte(origDatastr), blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize]) //初始向量的长度必须等于块block的长度16字节
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	pass64 := base64.StdEncoding.EncodeToString(crypted)
	return pass64, nil
}

// @title AES解密
// @param pass64 密文,Base64格式的密文
// @param keystring AES的密钥key,需与加密时使用的密钥相同
func AesDecrypt(pass64 string, keystring string) (string, error) {
	key := []byte(keystring)
	crypted, err := base64.StdEncoding.DecodeString(pass64)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	//AES分组长度为128位，所以blockSize=16，单位字节
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize]) //初始向量的长度必须等于块block的长度16字节
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return string(origData), nil
}

// 填充明文
func PKCS5Padding(plaintext []byte, blockSize int) []byte {
	padding := blockSize - len(plaintext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(plaintext, padtext...)
}

// 去除填充数据
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// @title 按参数名ASCII码从小到大排序 (key1=value1&key2=value2&key3=value3)
// @param data 待拼接的map数据集
// @param splite 连接符 &
// @param exceptKeys 被排除的参数名，不参与排序及拼接
func JoinStringsInASCII(data map[string]string, splite string, exceptKeys ...string) string {
	var list []string
	m := make(map[string]int)
	if len(exceptKeys) > 0 {
		for _, except := range exceptKeys {
			m[except] = 1
		}
	}
	for k, value := range data {
		if _, ok := m[k]; ok {
			continue
		}
		list = append(list, fmt.Sprintf("%s=%s", k, value))
	}
	sort.Strings(list)
	return strings.Join(list, splite)
}

// @title 解析字符串到map集合中
// @param originData 字符串(key1=value1&key2=value2&key3=value3)
// @param splite 连接符 &
func ParseSringToMap(originData string, splite string) map[string]string {
	dataMap := make(map[string]string)
	if originData != "" {
		array := strings.Split(originData, splite) // &
		for _, v := range array {
			arraylist := strings.Split(v, "=") // &
			if len(arraylist) == 2 {
				dataMap[arraylist[0]] = arraylist[1]
			}
		}
	}
	return dataMap
}

// @title md5加密
// @param originData 字符串
func Md5Encrypt(originData string) string {
	sum := md5.Sum([]byte(originData))
	return fmt.Sprintf("%x", sum)
}
