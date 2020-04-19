package aes

import (
	"testing"

	applog "github.com/lanvige/goutils/logger"
)

// TestEncrypt TestEncrypt
func TestEncrypt(t *testing.T) {
	data := "1011212004X"
	key := "abcdxxx"
	iv := "14265d0a68"

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
	data := "TGer5j10ZHn0adiIH8Xjhw=="
	key := "abcdxxx"
	iv := "14265d0a68"

	// dataByte := []byte(data)
	keyByte := []byte(key)
	ivByte := []byte(iv)
	code, err := Decrypt(data, keyByte, ivByte)

	if nil != err {
		applog.Error(err)
	}

	applog.Error(string(code))
}
