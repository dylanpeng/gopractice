package common

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
)

func padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

func unPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}

func AesEncrypt(origData, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	origData = padding(origData, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, iv)
	cryptData := make([]byte, len(origData))
	blockMode.CryptBlocks(cryptData, origData)
	return cryptData, nil
}

func AesDecrypt(cryptData, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(cryptData))
	blockMode.CryptBlocks(origData, cryptData)
	origData = unPadding(origData)
	return origData, nil
}

func DataAesEncrypt(data interface{}, key, iv string) (result string, e error) {
	mData, e := json.Marshal(data)

	if e != nil {
		return
	}

	aesData, e := AesEncrypt(mData, []byte(key), []byte(iv))

	if e != nil {
		return
	}

	result = base64.StdEncoding.EncodeToString(aesData)

	return
}

func DataAesDecrypt(data string, key, iv string) (result string, e error) {
	decodeData, e := base64.StdEncoding.DecodeString(data)

	if e != nil {
		return
	}

	decResult, e := AesDecrypt(decodeData, []byte(key), []byte(iv))

	if e != nil {
		return
	}

	result = string(decResult)

	return
}