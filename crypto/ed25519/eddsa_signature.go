package crypto

import (
	"crypto/ed25519"
	"encoding/base64"
	"encoding/hex"
	"errors"
)

// ==== PrivateKey Signature Operations ===== //
// ==== Signature-SHA1 Operations ===== //

// SignWithSha1Hex RsaSignWithSha1Hex
func (rsas *EdDSAEncrypt) SignWithSha1Hex(data []byte) (string, error) {
	if rsas.priKey == nil {
		return "", errors.New(`Please set the private key in advance`)
	}

	signature := ed25519.Sign(rsas.priKey, data)

	out := hex.EncodeToString(signature)
	return out, nil
}

// SignWithSha1Base64 SignWithSha1Base64
func (rsas *EdDSAEncrypt) SignWithSha1Base64(data []byte) (string, error) {
	if rsas.priKey == nil {
		return "", errors.New(`Please set the private key in advance`)
	}

	signature := ed25519.Sign(rsas.priKey, data)

	out := base64.StdEncoding.EncodeToString(signature)
	return out, nil
}
