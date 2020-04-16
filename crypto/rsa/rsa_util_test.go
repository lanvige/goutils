package crypto

import (
	"testing"

	applog "github.com/lanvige/goutils/logger"
)

// TestRSA256PKCS1KeyGen TestRSA256KeyGen
func TestRSA256PKCS1KeyGen(t *testing.T) {
	priKey, pubKey, _ := GenerateRSAKey(256)

	priKeyPem, _ := DumpPriKeyPKCS1Base64(priKey)
	pubKeyPem, _ := DumpPubPKCS1Base64(pubKey)

	applog.Info(priKeyPem)
	applog.Info(pubKeyPem)
}

func TestRSA256PKCS8KeyGen(t *testing.T) {
	priKey, pubKey, _ := GenerateRSAKey(256)

	priKeyPem, _ := DumpPriKeyPKCS8Base64(priKey)
	pubKeyPem, _ := DumpPubPKCS8Base64(pubKey)

	applog.Info(priKeyPem)
	applog.Warn(pubKeyPem)
}

// TestRSA512KeyGen TestRSA512KeyGen
func TestRSA512PKCS1KeyGen(t *testing.T) {
	priKey, pubKey, _ := GenerateRSAKey(512)

	priKeyPem, _ := DumpPriKeyPKCS1Base64(priKey)
	pubKeyPem, _ := DumpPubPKCS1Base64(pubKey)

	applog.Info(priKeyPem)
	applog.Warn(pubKeyPem)
}

func TestRSA512PKCS8KeyGen(t *testing.T) {
	priKey, pubKey, _ := GenerateRSAKey(512)

	priKeyPem, _ := DumpPriKeyPKCS8Base64(priKey)
	pubKeyPem, _ := DumpPubPKCS8Base64(pubKey)

	applog.Info(priKeyPem)
	applog.Warn(pubKeyPem)
}
