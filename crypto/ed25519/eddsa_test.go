package crypto

import (
	"encoding/base64"
	"encoding/hex"
	"testing"

	applog "github.com/lanvige/goutils/logger"
)

// PKCS8 256
var Pubkey = `-----BEGIN 公钥-----
MDwwDQYJKoZIhvcNAQEBBQADKwAwKAIhAKd0PbKjNEFOMmYRHJ28ucRTELcEj/H3YFqcIreFTtmJAgMBAAE=
-----END 公钥-----
`

// PKCS8 256 私钥
var Pirvatekey = `-----BEGIN 私钥-----
MIHBAgEAMA0GCSqGSIb3DQEBAQUABIGsMIGpAgEAAiEAp3Q9sqM0QU4yZhEcnby5x
FMQtwSP8fdgWpwit4VO2YkCAwEAAQIgPSbBBUR+Z77jvsxO8/egqs8j8ueUBEb2Pa
36RXsloTUCEQDM7w8RrHnoJgqZqpbU7zyLAhEA0S5PP16b5KDwWSeI/zzguwIQKpY
nKIISF0cIfuRvUbKhfwIRAJjaHxxuNLonBBoRsqDXylsCDz30xZWMCKys0CuQJZQc
/g==
-----END 私钥-----
`

// func Test_SetPublicKey(t *testing.T) {
// 	if err := Ed25519.SetPubKey(Pubkey); err != nil {
// 		t.Error(err)
// 	}
// }

// func Test_SetPrivateKey(t *testing.T) {
// 	if err := Ed25519.SetPriKey(Pirvatekey); err != nil {
// 		t.Error(err)
// 	}
// }

func TestRSA256PKCS8KeyGen(t *testing.T) {
	priKey, pubKey, _ := GenerateKey()

	// 私钥 64，公钥 32 byte
	// 所以 hex 是 128

	prikeybase64 := base64.StdEncoding.EncodeToString([]byte(*priKey))
	prikeyhex := hex.EncodeToString([]byte(*priKey))

	applog.Warn(prikeybase64)
	applog.Warn(prikeyhex)

	pubkeybase64 := base64.StdEncoding.EncodeToString([]byte(*pubKey))
	pubkeyhex := hex.EncodeToString([]byte(*pubKey))

	applog.Warn(pubkeybase64)
	applog.Warn(pubkeyhex)

	abc, _ := priKey.Sign()([]byte("dc95794a52749100bac61f8195e30162264141af4fea26e5446fa88858ed86aadc6ea14248af80db95b81b97e436e37107f7d2cdfb875401da6e598547e72713"))

}

// dc95794a52749100bac61f8195e30162264141af4fea26e5446fa88858ed86aadc6ea14248af80db95b81b97e436e37107f7d2cdfb875401da6e598547e72713
// dc6ea14248af80db95b81b97e436e37107f7d2cdfb875401da6e598547e72713

func TestRSA256PKCS8Key(t *testing.T) {
	Ed25519.SetPriKeyHex("dc95794a52749100bac61f8195e30162264141af4fea26e5446fa88858ed86aadc6ea14248af80db95b81b97e436e37107f7d2cdfb875401da6e598547e72713")

	abc, _ := Ed25519.SignWithSha1Hex([]byte("dc95794a52749100bac61f8195e30162264141af4fea26e5446fa88858ed86aadc6ea14248af80db95b81b97e436e37107f7d2cdfb875401da6e598547e72713"))

	applog.Warn(abc)
	// 00b21ea7e6c92a3330d0887436fa1e63d0227b38f197ee870dcaa41da563d9f29e153aecb135c3cb7dd6971d3390039306d192f987d1b5653fdf647c9f88790e
}

// // 私钥加密公钥解密
// func Test_PriSignPubVer(t *testing.T) {
// 	applog.Info("1")
// 	if err := Ed25519.SetPubKey(Pubkey); err != nil {
// 		applog.Info("1")
// 		return
// 		t.Error(err)
// 	}

// 	applog.Info("2")
// 	if err := Ed25519.SetPriKey(Pirvatekey); err != nil {
// 		applog.Info("2")

// 		t.Error(err)
// 	}

// 	applog.Info("3")
// 	prienctypt, err := Ed25519.SignWithSha1Base64([]byte(`hello world`))
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	applog.Info(prienctypt)

// 	// base64content := base64.StdEncoding.EncodeToString(prienctypt)
// 	// applog.Info(base64content)

// 	// err2 := Ed25519.VerySignWithSha1Base64(`hello world`, prienctypt)
// 	// if err2 != nil {
// 	// 	t.Error(err2)
// 	// }

// 	applog.Info("success")
// }
