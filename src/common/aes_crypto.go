package common

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
)

var Key = "QPMsNI1NaYsKnevjvmurTVbfc3IDITI4"
var Vi = "zNiMeSzszmP3JYkY"

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

func StringAesEncrypt(data, key, iv string) (result string, e error) {
	aesData, e := AesEncrypt([]byte(data), []byte(key), []byte(iv))

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


// https://github.com/fkfk/mysqlcrypto/blob/master/mysqlcrypto.go
func MysqlAESEncrypt(src []byte, key []byte) (encrypted []byte) {
	cipher, _ := aes.NewCipher(generateKey(key))
	length := (len(src) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, src)
	pad := byte(len(plain) - len(src))
	for i := len(src); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted = make([]byte, len(plain))

	for bs, be := 0, cipher.BlockSize(); bs <= len(src); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted
}

func MysqlAESDecrypt(encrypted []byte, key []byte) (decrypted []byte) {
	cipher, _ := aes.NewCipher(generateKey(key))
	decrypted = make([]byte, len(encrypted))
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return decrypted[:trim]
}

func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}