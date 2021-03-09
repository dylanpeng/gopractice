package main

import (
	"encoding/base64"
	"fmt"
	"github.com/forgoer/openssl"
)

func main() {
	key := []byte("123456789012345678901234")
	encodeData, _ := openssl.Des3ECBEncrypt([]byte("aaasdfesef"), key, openssl.ZEROS_PADDING)
	fmt.Printf("加密:%s\n", encodeData)
	encryptBaseData := base64.StdEncoding.EncodeToString(encodeData)
	fmt.Printf("base64加密:%s\n", encryptBaseData)
	decodeBaseData, _ := base64.StdEncoding.DecodeString(encryptBaseData)
	fmt.Printf("base64解密:%s\n", decodeBaseData)
	decodeData, _ := openssl.Des3ECBDecrypt(decodeBaseData, key, openssl.PKCS7_PADDING)
	fmt.Printf("解密:%s\n", decodeData)
}
