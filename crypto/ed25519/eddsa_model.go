package crypto

import (
	"crypto/ed25519"
	"encoding/hex"
)

// Ed25519 Ed25519
var Ed25519 = &EdDSAEncrypt{}

// EdDSAEncrypt EdDSAEncrypt
type EdDSAEncrypt struct {
	pubKeyStr string             // 公钥字符串
	priKeyStr string             // 私钥字符串
	pubKey    ed25519.PublicKey  // 公钥
	priKey    ed25519.PrivateKey // 私钥
}

// // SetPubKey 设置公钥
// func (rsas *EdDSAEncrypt) SetPubKey(pubKeyStr string) (err error) {
// 	rsas.pubKeyStr = pubKeyStr
// 	rsas.pubKey, err = getPubKey([]byte(pubKeyStr))

// 	return err
// }

// // SetPriKey 设置私钥
// func (rsas *EdDSAEncrypt) SetPriKey(priKeyStr string) (err error) {
// 	rsas.priKeyStr = priKeyStr
// 	rsas.priKey, err = getPriKey([]byte(priKeyStr))

// 	return err
// }

// SetPriKeyHex 设置私钥
func (rsas *EdDSAEncrypt) SetPriKeyHex(priKeyStr string) (err error) {
	rsas.priKeyStr = priKeyStr

	abc, _ := hex.DecodeString(priKeyStr)
	rsas.priKey = abc

	return err
}

// // GetPrivatekey *rsa.PublicKey
// func (rsas *EdDSAEncrypt) GetPrivatekey() (*ed25519.PrivateKey, error) {
// 	if rsas.priKey == nil {
// 		return nil, errors.New(`PriKey not exist`)
// 	}

// 	return rsas.priKey, nil
// }

// // GetPublickey *rsa.PrivateKey
// func (rsas *EdDSAEncrypt) GetPublickey() (*ed25519.PublicKey, error) {
// 	if rsas.pubKey == nil {
// 		return nil, errors.New(`PubKey not exist`)
// 	}

// 	return rsas.pubKey, nil
// }
