package crypto

import (
	"bytes"
)

// PKCS7Padding PKCS7补码
func PKCS7Padding(data []byte) []byte {
	blockSize := 16
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)

}

// PKCS7Unpadding 去除PKCS7的补码
func PKCS7Unpadding(data []byte) []byte {
	length := len(data)

	// 去掉最后一个字节 unpadding 次
	unpadding := int(data[length-1])
	if length <= unpadding {
		return nil
	}

	return data[:(length - unpadding)]
}
