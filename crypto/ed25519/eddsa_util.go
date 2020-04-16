package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

// ==== ed25519 Key Operations ===== //
// ==== ed25519 Key Operations ===== //

// 公钥转换 字符串 -> 对象
func getPubKey(publickey []byte) (*ed25519.PublicKey, error) {
	// decode public key
	block, _ := pem.Decode(publickey)
	if block == nil {
		return nil, errors.New("get public key error")
	}

	// x509 parse public key
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return pub.(*ed25519.PublicKey), err
}

// 私钥转换 字符串 -> 对象
func getPriKey(privartekey []byte) (*ed25519.PrivateKey, error) {
	block, _ := pem.Decode(privartekey)
	if block == nil {
		return nil, errors.New("get private key error")
	}

	pri2, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return pri2.(*ed25519.PrivateKey), nil
}

// ==== GenRsaKey Operations ===== //
// ==== GenRsaKey Operations ===== //

// GenerateKey GenerateKey
func GenerateKey() (*ed25519.PrivateKey, *ed25519.PublicKey, error) {
	// Gen Private key
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)

	if err != nil {
		return nil, nil, err
	}

	return &privateKey, &publicKey, nil
}

// DumpPriKeyPKCS8Base64 DumpPriKeyPKCS8Base64
func DumpPriKeyPKCS8Base64(privateKey *ed25519.PrivateKey) (string, error) {
	derStream, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		return "", err
	}

	keybase64 := base64.StdEncoding.EncodeToString(derStream)

	return keybase64, nil
}

// DumpPubPKCS8Base64 DumpPubPKCS8Base64
func DumpPubPKCS8Base64(publickey *ed25519.PublicKey) (string, error) {
	derPkix, err := x509.MarshalPKIXPublicKey(publickey)
	if err != nil {
		return "", err
	}

	keybase64 := base64.StdEncoding.EncodeToString(derPkix)

	return keybase64, nil
}
