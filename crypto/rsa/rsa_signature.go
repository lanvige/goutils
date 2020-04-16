package crypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
)

// ==== PrivateKey Signature Operations ===== //
// ==== Signature-SHA1 Operations ===== //

// SignWithSha1Hex RsaSignWithSha1Hex
func (rsas *RSAEncrypt) SignWithSha1Hex(data string) (string, error) {
	if rsas.priKey == nil {
		return "", errors.New(`Please set the private key in advance`)
	}

	h := sha1.New()
	h.Write([]byte([]byte(data)))
	hash := h.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, rsas.priKey, crypto.SHA1, hash[:])
	if err != nil {
		fmt.Printf("Error from signing: %s\n", err)
		return "", err
	}

	out := hex.EncodeToString(signature)
	return out, nil
}

// SignWithSha1Base64 SignWithSha1Base64
func (rsas *RSAEncrypt) SignWithSha1Base64(data string) (string, error) {
	if rsas.priKey == nil {
		return "", errors.New(`Please set the private key in advance`)
	}

	h := sha1.New()
	h.Write([]byte([]byte(data)))
	hash := h.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, rsas.priKey, crypto.SHA1, hash[:])
	if err != nil {
		fmt.Printf("Error from signing: %s\n", err)
		return "", err
	}

	out := base64.StdEncoding.EncodeToString(signature)
	return out, nil
}

// VerySignWithSha1Hex 验签
func (rsas *RSAEncrypt) VerySignWithSha1Hex(originalData, signData string) error {
	if rsas.priKey == nil {
		return errors.New(`Please set the private key in advance`)
	}

	sign, err := hex.DecodeString(signData)
	if err != nil {
		return err
	}

	h := sha1.New()
	h.Write([]byte([]byte(originalData)))
	hash := h.Sum(nil)

	return rsa.VerifyPKCS1v15(rsas.pubKey, crypto.SHA1, hash, sign)
}

// VerySignWithSha1Base64 验签
func (rsas *RSAEncrypt) VerySignWithSha1Base64(originalData, signData string) error {
	sign, err := base64.StdEncoding.DecodeString(signData)
	if err != nil {
		return err
	}

	h := sha1.New()
	h.Write([]byte([]byte(originalData)))
	hash := h.Sum(nil)

	return rsa.VerifyPKCS1v15(rsas.pubKey, crypto.SHA1, hash, sign)
}

// ==== PrivateKey Signature Operations ===== //
// ==== Signature-SHA256 Operations ===== //

// SignWithSha256Hex SignWithSha256Hex
func (rsas *RSAEncrypt) SignWithSha256Hex(data string) (string, error) {
	if rsas.priKey == nil {
		return "", errors.New(`Please set the private key in advance`)
	}

	h := sha256.New()
	h.Write([]byte([]byte(data)))
	hash := h.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, rsas.priKey, crypto.SHA256, hash[:])
	if err != nil {
		fmt.Printf("Error from signing: %s\n", err)
		return "", err
	}

	out := hex.EncodeToString(signature)
	return out, nil
}

// SignWithSha256Base64 SignWithSha256Base64
func (rsas *RSAEncrypt) SignWithSha256Base64(data string) (string, error) {
	if rsas.priKey == nil {
		return "", errors.New(`Please set the private key in advance`)
	}

	h := sha1.New()
	h.Write([]byte([]byte(data)))
	hash := h.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, rsas.priKey, crypto.SHA256, hash[:])
	if err != nil {
		fmt.Printf("Error from signing: %s\n", err)
		return "", err
	}

	out := base64.StdEncoding.EncodeToString(signature)
	return out, nil
}

// VerySignWithSha256Hex 验签
func (rsas *RSAEncrypt) VerySignWithSha256Hex(originalData, signData string) error {
	if rsas.priKey == nil {
		return errors.New(`Please set the private key in advance`)
	}

	sign, err := hex.DecodeString(signData)
	if err != nil {
		return err
	}

	h := sha256.New()
	h.Write([]byte([]byte(originalData)))
	hash := h.Sum(nil)

	return rsa.VerifyPKCS1v15(rsas.pubKey, crypto.SHA256, hash, sign)
}

// VerySignWithSha256Base64 验签
func (rsas *RSAEncrypt) VerySignWithSha256Base64(originalData, signData string) error {
	if rsas.priKey == nil {
		return errors.New(`Please set the private key in advance`)
	}

	sign, err := base64.StdEncoding.DecodeString(signData)
	if err != nil {
		return err
	}

	h := sha256.New()
	h.Write([]byte([]byte(originalData)))
	hash := h.Sum(nil)

	return rsa.VerifyPKCS1v15(rsas.pubKey, crypto.SHA256, hash, sign)
}

// ==== PrivateKey Signature Operations ===== //
// ==== Signature-SHA512 Operations ===== //

// SignWithSha512Hex SignWithSha512Hex
func (rsas *RSAEncrypt) SignWithSha512Hex(data string) (string, error) {
	if rsas.priKey == nil {
		return "", errors.New(`Please set the private key in advance`)
	}

	h := sha512.New()
	h.Write([]byte([]byte(data)))
	hash := h.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, rsas.priKey, crypto.SHA512, hash[:])
	if err != nil {
		fmt.Printf("Error from signing: %s\n", err)
		return "", err
	}

	out := hex.EncodeToString(signature)
	return out, nil
}

// SignWithSha512Base64 SignWithSha512Base64
func (rsas *RSAEncrypt) SignWithSha512Base64(data string) (string, error) {
	if rsas.priKey == nil {
		return "", errors.New(`Please set the private key in advance`)
	}

	h := sha512.New()
	h.Write([]byte([]byte(data)))
	hash := h.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, rsas.priKey, crypto.SHA512, hash[:])
	if err != nil {
		fmt.Printf("Error from signing: %s\n", err)
		return "", err
	}

	out := base64.StdEncoding.EncodeToString(signature)
	return out, nil
}

// VerySignWithSha512Hex 验签
func (rsas *RSAEncrypt) VerySignWithSha512Hex(originalData, signData string) error {
	if rsas.priKey == nil {
		return errors.New(`Please set the private key in advance`)
	}

	sign, err := hex.DecodeString(signData)
	if err != nil {
		return err
	}

	h := sha512.New()
	h.Write([]byte([]byte(originalData)))
	hash := h.Sum(nil)

	return rsa.VerifyPKCS1v15(rsas.pubKey, crypto.SHA512, hash, sign)
}

// VerySignWithSha512Base64 验签
func (rsas *RSAEncrypt) VerySignWithSha512Base64(originalData, signData string) error {
	if rsas.priKey == nil {
		return errors.New(`Please set the private key in advance`)
	}

	sign, err := base64.StdEncoding.DecodeString(signData)
	if err != nil {
		return err
	}

	h := sha512.New()
	h.Write([]byte([]byte(originalData)))
	hash := h.Sum(nil)

	return rsa.VerifyPKCS1v15(rsas.pubKey, crypto.SHA512, hash, sign)
}
