package crypto

import (
	"encoding/base64"
	"testing"

	applog "github.com/lanvige/goutils/logger"
)

var Pubkey = `-----BEGIN 公钥-----
MDwwDQYJKoZIhvcNAQEBBQADKwAwKAIhAMtetY4WKELOEX1S52FUtBMrio/qZRG0
YVdAvqBZzIyHAgMBAAE=
-----END 公钥-----
`

// PKCS1 私钥
var Pirvatekey = `-----BEGIN 私钥-----
MIGqAgEAAiEAy161jhYoQs4RfVLnYVS0EyuKj+plEbRhV0C+oFnMjIcCAwEAAQIg
PDm/6mQFaEM+Sq7SdrBu4vbzZpAE+hQfh395q3J0Z4ECEQDUTUpSrfGqltKFaI6Q
WXxtAhEA9TrFpy47xsMwiYOC9Y9sQwIQNBirTlKk5OmRy2723PE4mQIQL+gqYs2a
aesJ2WNyuwwoxwIRAIEzIqqjlfAZoF27H9dYB5w=
-----END 私钥-----
`

var Pubkey2 = `-----BEGIN 公钥-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAk+89V7vpOj1rG6bTAKYM
56qmFLwNCBVDJ3MltVVtxVUUByqc5b6u909MmmrLBqS//PWC6zc3wZzU1+ayh8xb
UAEZuA3EjlPHIaFIVIz04RaW10+1xnby/RQE23tDqsv9a2jv/axjE/27b62nzvCW
eItu1kNQ3MGdcuqKjke+LKhQ7nWPRCOd/ffVqSuRvG0YfUEkOz/6UpsPr6vrI331
hWRB4DlYy8qFUmDsyvvExe4NjZWblXCqkEXRRAhi2SQRCl3teGuIHtDUxCskRIDi
aMD+Qt2Yp+Vvbz6hUiqIWSIH1BoHJer/JOq2/O6X3cmuppU4AdVNgy8Bq236iXvr
MQIDAQAB
-----END 公钥-----
`

// PKCS1 私钥 512 位？
var Pirvatekey2 = `-----BEGIN 私钥-----
MIIEpAIBAAKCAQEAk+89V7vpOj1rG6bTAKYM56qmFLwNCBVDJ3MltVVtxVUUByqc
5b6u909MmmrLBqS//PWC6zc3wZzU1+ayh8xbUAEZuA3EjlPHIaFIVIz04RaW10+1
xnby/RQE23tDqsv9a2jv/axjE/27b62nzvCWeItu1kNQ3MGdcuqKjke+LKhQ7nWP
RCOd/ffVqSuRvG0YfUEkOz/6UpsPr6vrI331hWRB4DlYy8qFUmDsyvvExe4NjZWb
lXCqkEXRRAhi2SQRCl3teGuIHtDUxCskRIDiaMD+Qt2Yp+Vvbz6hUiqIWSIH1BoH
Jer/JOq2/O6X3cmuppU4AdVNgy8Bq236iXvrMQIDAQABAoIBAQCCbxZvHMfvCeg+
YUD5+W63dMcq0QPMdLLZPbWpxMEclH8sMm5UQ2SRueGY5UBNg0WkC/R64BzRIS6p
jkcrZQu95rp+heUgeM3C4SmdIwtmyzwEa8uiSY7Fhbkiq/Rly6aN5eB0kmJpZfa1
6S9kTszdTFNVp9TMUAo7IIE6IheT1x0WcX7aOWVqp9MDXBHV5T0Tvt8vFrPTldFg
IuK45t3tr83tDcx53uC8cL5Ui8leWQjPh4BgdhJ3/MGTDWg+LW2vlAb4x+aLcDJM
CH6Rcb1b8hs9iLTDkdVw9KirYQH5mbACXZyDEaqj1I2KamJIU2qDuTnKxNoc96HY
2XMuSndhAoGBAMPwJuPuZqioJfNyS99x++ZTcVVwGRAbEvTvh6jPSGA0k3cYKgWR
NnssMkHBzZa0p3/NmSwWc7LiL8whEFUDAp2ntvfPVJ19Xvm71gNUyCQ/hojqIAXy
tsNT1gBUTCMtFZmAkUsjqdM/hUnJMM9zH+w4lt5QM2y/YkCThoI65BVbAoGBAMFI
GsIbnJDNhVap7HfWcYmGOlWgEEEchG6Uq6Lbai9T8c7xMSFc6DQiNMmQUAlgDaMV
b6izPK4KGQaXMFt5h7hekZgkbxCKBd9xsLM72bWhM/nd/HkZdHQqrNAPFhY6/S8C
IjRnRfdhsjBIA8K73yiUCsQlHAauGfPzdHET8ktjAoGAQdxeZi1DapuirhMUN9Zr
kr8nkE1uz0AafiRpmC+cp2Hk05pWvapTAtIXTo0jWu38g3QLcYtWdqGa6WWPxNOP
NIkkcmXJjmqO2yjtRg9gevazdSAlhXpRPpTWkSPEt+o2oXNa40PomK54UhYDhyeu
akuXQsD4mCw4jXZJN0suUZMCgYAgzpBcKjulCH19fFI69RdIdJQqPIUFyEViT7Hi
bsPTTLham+3u78oqLzQukmRDcx5ddCIDzIicMfKVf8whertivAqSfHytnf/pMW8A
vUPy5G3iF5/nHj76CNRUbHsfQtv+wqnzoyPpHZgVQeQBhcoXJSm+qV3cdGjLU6OM
HgqeaQKBgQCnmL5SX7GSAeB0rSNugPp2GezAQj0H4OCc8kNrHK8RUvXIU9B2zKA2
z/QUKFb1gIGcKxYr+LqQ25/+TGvINjuf6P3fVkHL0U8jOG0IqpPJXO3Vl9B8ewWL
cFQVB/nQfmaMa4ChK0QEUe+Mqi++MwgYbRHx1lIOXEfUJO+PXrMekw==
-----END 私钥-----
`

func Test_SetPublicKey(t *testing.T) {
	if err := RSA.SetPubKey(Pubkey); err != nil {
		t.Error(err)
	}
}

func Test_SetPrivateKey(t *testing.T) {
	if err := RSA.SetPriKey(Pirvatekey); err != nil {
		t.Error(err)
	}
}

// 加载私钥，解密出公钥
func Test_GenPrivateKey(t *testing.T) {
	if err := RSA.SetPriKey(Pirvatekey); err != nil {
		t.Error(err)
	}

	priKey, err2 := RSA.GetPrivatekey()
	if err2 != nil {
		t.Error(err2)
	}

	abc := priKey.PublicKey

	abcd, _ := DumpPubPKCS1Base64(&abc)

	applog.Info(abcd)
}

// 公钥加密私钥解密
// 这个才是最常用的用法
func Test_PubENCTYPTPriDECRYPT(t *testing.T) {
	if err := RSA.SetPubKey(Pubkey); err != nil {
		t.Error(err)
	}
	if err := RSA.SetPriKey(Pirvatekey); err != nil {
		t.Error(err)
	}

	pubenctypt, err := RSA.PubKeyEncrypt([]byte(`hello world`))
	if err != nil {
		t.Error(err)
	}

	base64content := base64.StdEncoding.EncodeToString(pubenctypt)
	applog.Info(base64content)

	// thenewpubencrypt, _ := base64.StdEncoding.DecodeString(base64content)
	// pridecrypt, err := RSA.PriKeyDecrypt(thenewpubencrypt)

	pridecrypt, err := RSA.PriKeyDecrypt(pubenctypt)

	if err != nil {
		t.Error(err)
	}

	applog.Info(string(pridecrypt))

	if string(pridecrypt) != `hello world` {
		t.Error(`不符合预期`)
	}

	applog.Info("Success")
}

// 私钥加密公钥解密
// func Test_PriENCTYPTPubDECRYPT(t *testing.T) {
// 	if err := RSA.SetPubKey(Pubkey); err != nil {
// 		t.Error(err)
// 	}
// 	if err := RSA.SetPriKey(Pirvatekey); err != nil {
// 		t.Error(err)
// 	}

// 	prienctypt, err := RSA.PriKeyEncrypt([]byte(`hello world`))
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	base64content := base64.StdEncoding.EncodeToString(prienctypt)
// 	applog.Info(base64content)

// 	pubdecrypt, err := RSA.PubKeyDecrypt(prienctypt)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	if string(pubdecrypt) != `hello world` {
// 		t.Error(`不符合预期`)
// 	}

// 	applog.Info(string(pubdecrypt))
// }

// 私钥加密公钥解密
func Test_PriSignPubVer(t *testing.T) {
	if err := RSA.SetPubKey(Pubkey2); err != nil {
		t.Error(err)
	}
	if err := RSA.SetPriKey(Pirvatekey2); err != nil {
		t.Error(err)
	}

	prienctypt, err := RSA.SignWithSha1Base64(`hello world`)
	if err != nil {
		t.Error(err)
	}

	applog.Info(prienctypt)

	// base64content := base64.StdEncoding.EncodeToString(prienctypt)
	// applog.Info(base64content)

	err2 := RSA.VerySignWithSha1Base64(`hello world`, prienctypt)
	if err2 != nil {
		t.Error(err2)
	}

	applog.Info("success")
}
