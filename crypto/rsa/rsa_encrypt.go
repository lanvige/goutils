package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"errors"
)

// ==== Public Encrypt Operations ===== //
// ==== Encrypt Operations ===== //

// PubKeyEncrypt 公钥加密
func (rsas *RSAEncrypt) PubKeyEncrypt(input []byte) ([]byte, error) {
	if rsas.pubKey == nil {
		return []byte(""), errors.New(`Please set the public key in advance`)
	}

	encryptedData, err := rsa.EncryptPKCS1v15(rand.Reader, rsas.pubKey, input)
	if err != nil {
		return []byte(""), err
	}

	return encryptedData, nil
}

// PriKeyDecrypt 私钥解密
func (rsas *RSAEncrypt) PriKeyDecrypt(input []byte) ([]byte, error) {
	if rsas.priKey == nil {
		return []byte(""), errors.New(`Please set the private key in advance`)
	}

	originalData, err := rsa.DecryptPKCS1v15(rand.Reader, rsas.priKey, input)
	if err != nil {
		return []byte(""), err
	}

	return originalData, err
}

// ==== Private Encrypt Operations ===== //
// ==== Encrypt Operations ===== //
