package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"

	goutilscrypto "github.com/lanvige/goutils/crypto"
)

func expendKey(key []byte) []byte {
	for len(key) < 16 {
		key = append(key, key...)
	}

	return key[:16]
}

// Encrypt Encrypt
func Encrypt(data []byte, key, iv []byte) (string, error) {
	if data == nil || len(data) == 0 {
		return "", errors.New("Encrypt Params error")
	}
	if key == nil || len(key) == 0 {
		return "", errors.New("Encrypt Params error")
	}
	if iv == nil || len(iv) == 0 {
		return "", errors.New("Encrypt Params error")
	}

	key = expendKey(key)
	iv = expendKey(iv)

	ckey, err := aes.NewCipher(key)
	if nil != err {
		return "", err
	}

	encrypter := cipher.NewCBCEncrypter(ckey, iv)

	// PKCS7补码
	str := cryptolib.PKCS7Padding([]byte(data))
	out := make([]byte, len(str))

	encrypter.CryptBlocks(out, str)

	return base64.StdEncoding.EncodeToString(out), nil
}

// Decrypt Decrypt
func Decrypt(base64Str string, key, iv []byte) ([]byte, error) {
	key = expendKey(key)
	iv = expendKey(iv)

	ckey, err := aes.NewCipher(key)
	if nil != err {
		return nil, err
	}

	decrypter := cipher.NewCBCDecrypter(ckey, iv)

	base64In, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return nil, err
	}

	out := make([]byte, len(base64In))
	decrypter.CryptBlocks(out, base64In)

	// 去除 PKCS7 补码
	out = goutilscrypto.PKCS7Unpadding(out)
	if out == nil {
		return nil, nil
	}

	return out, nil
}
