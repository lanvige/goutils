package aes

import (
	"testing"

	applog "github.com/lanvige/goutils/logger"
)

// TestEncrypt TestEncrypt
func TestEncrypt(t *testing.T) {
	data := "10112319970712004X"
	key := "-"
	iv := "14266e140b5d0a68"

	dataByte := []byte(data)
	keyByte := []byte(key)
	ivByte := []byte(iv)
	dataEncrypt, err := Encrypt(dataByte, keyByte, ivByte)

	if nil != err {
		applog.Error(err)
	}

	applog.Error(dataEncrypt)
}

// TestDecrypt TestDecrypt
func TestDecryptq(t *testing.T) {
	data := "ImK1fgeV7XWqigpOTs5fnpA7youa/hh373WalpEZKcY="
	key := "-"
	iv := "14266e140b5d0a68"

	// dataByte := []byte(data)
	keyByte := []byte(key)
	ivByte := []byte(iv)
	code, err := Decrypt(data, keyByte, ivByte)

	if nil != err {
		applog.Error(err)
	}

	applog.Error(string(code))
}
