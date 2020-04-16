package crypto

import (
	"crypto/rsa"
	"errors"
)

// RSA RSA
var RSA = &RSAEncrypt{}

// RSAEncrypt RSAEncrypt
type RSAEncrypt struct {
	pubKeyStr string          // 公钥字符串
	priKeyStr string          // 私钥字符串
	pubKey    *rsa.PublicKey  // 公钥
	priKey    *rsa.PrivateKey // 私钥
}

// SetPubKey 设置公钥
func (rsas *RSAEncrypt) SetPubKey(pubKeyStr string) (err error) {
	rsas.pubKeyStr = pubKeyStr
	rsas.pubKey, err = getPubKey([]byte(rsas.pubKeyStr))

	return err
}

// SetPriKey 设置私钥
func (rsas *RSAEncrypt) SetPriKey(priKeyStr string) (err error) {
	rsas.priKeyStr = priKeyStr
	rsas.priKey, err = getPriKey([]byte(rsas.priKeyStr))

	return err
}

// GetPrivatekey *rsa.PublicKey
func (rsas *RSAEncrypt) GetPrivatekey() (*rsa.PrivateKey, error) {
	if rsas.priKey == nil {
		return nil, errors.New(`PriKey not exist`)
	}

	return rsas.priKey, nil
}

// GetPublickey *rsa.PrivateKey
func (rsas *RSAEncrypt) GetPublickey() (*rsa.PublicKey, error) {
	if rsas.pubKey == nil {
		return nil, errors.New(`PubKey not exist`)
	}

	return rsas.pubKey, nil
}
