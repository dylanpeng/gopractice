package common

import (
	"crypto/aes"
	"encoding/base64"
)

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

func MysqlAESEncryptString(srcStr, keyStr string) (encryptedStr string) {
	enc := MysqlAESEncrypt([]byte(srcStr), []byte(keyStr))
	encryptedStr = Base64Encode(enc)
	return
}

func MysqlAESDecryptString(encryptedStr, keyStr string) (srcStr string, err error) {
	encrypted, err := Base64Decode(encryptedStr)

	if err != nil {
		return
	}

	srcStr = string(MysqlAESDecrypt(encrypted, []byte(keyStr)))
	return
}

// base64编码
func Base64Encode(data []byte) string {
	base64encodeBytes := base64.StdEncoding.EncodeToString(data)
	return base64encodeBytes
}

// base64解码
func Base64Decode(data string) ([]byte, error) {
	decodeBytes, err := base64.StdEncoding.DecodeString(data)
	return decodeBytes, err
}