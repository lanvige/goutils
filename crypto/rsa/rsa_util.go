package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

// ==== GenRsaKey Operations ===== //
// ==== GenRsaKey Operations ===== //

// 公钥转换 字符串 -> 对象
func getPubKey(publickey []byte) (*rsa.PublicKey, error) {
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

	return pub.(*rsa.PublicKey), err
}

// 私钥转换 字符串 -> 对象
func getPriKey(privartekey []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(privartekey)
	if block == nil {
		return nil, errors.New("get private key error")
	}

	pri, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err == nil {
		return pri, nil
	}

	pri2, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return pri2.(*rsa.PrivateKey), nil
}

// ==== GenRsaKey Operations ===== //
// ==== GenRsaKey Operations ===== //

// GenerateRSAKey GenerateRSAKey
func GenerateRSAKey(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	// Gen Private key
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)

	if err != nil {
		return nil, nil, err
	}

	publicKey := &privateKey.PublicKey

	return privateKey, publicKey, nil
}

// DumpPriKeyPKCS1Base64 DumpPriKeyPKCS1Base64
func DumpPriKeyPKCS1Base64(privateKey *rsa.PrivateKey) (string, error) {
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)

	keybase64 := base64.StdEncoding.EncodeToString(derStream)

	return keybase64, nil
}

// DumpPubPKCS1Base64 DumpPubPKCS1Base64
func DumpPubPKCS1Base64(publickey *rsa.PublicKey) (string, error) {
	derPkix, err := x509.MarshalPKIXPublicKey(publickey)
	if err != nil {
		return "", err
	}

	keybase64 := base64.StdEncoding.EncodeToString(derPkix)

	return keybase64, nil
}

// DumpPriKeyPKCS8Base64 DumpPriKeyPKCS8Base64
func DumpPriKeyPKCS8Base64(privateKey *rsa.PrivateKey) (string, error) {
	derStream, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		return "", err
	}

	keybase64 := base64.StdEncoding.EncodeToString(derStream)

	return keybase64, nil
}

// DumpPubPKCS8Base64 DumpPubPKCS8Base64
func DumpPubPKCS8Base64(publickey *rsa.PublicKey) (string, error) {
	derPkix, err := x509.MarshalPKIXPublicKey(publickey)
	if err != nil {
		return "", err
	}

	keybase64 := base64.StdEncoding.EncodeToString(derPkix)

	return keybase64, nil
}
