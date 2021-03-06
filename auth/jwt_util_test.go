package authorization

import (
	"testing"
	"time"

	applog "github.com/lanvige/goutils/logger"

	jwtgo "github.com/dgrijalva/jwt-go"
)

// Define
// TestJWTPayload TestJWTPayload
type TestJWTPayload struct {
	jwtgo.StandardClaims
	// Provider string `json:"oap,omitempty"` //oap: OAuth Provider
	Mode      string `json:"mode,omitempty"`
	SessionID string `json:"sid,omitempty"`
	Sign      string `json:"sig,omitempty"`
	SignAlg   string `json:"sal,omitempty"`
}

func TestGenerateHSToken(t *testing.T) {
	myJWTClaims := TestJWTPayload{
		StandardClaims: jwtgo.StandardClaims{
			Issuer:    "30050",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
		// Provider: provider,
		SessionID: "84354f4578f91a",
	}

	code, _ := GenerateHS256Token("ihY.dNu", &myJWTClaims)

	applog.Info(code)
}

func TestGenerateRSA256Token(t *testing.T) {
	myJWTClaims := TestJWTPayload{
		StandardClaims: jwtgo.StandardClaims{
			Issuer:    "30001",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 168).Unix(),
		},
	}

	// cert := "MIIBVAIBADANBgkqhkiG9w0BAQEFAASCAT4wggE6AgEAAkEA1O/W1f3BTpMIQI9z42YJVA/+2Igq7KwIVIuiVzYw2qpClaWMnKBfzObso0CbP953lxiRHyQVDI5tdSs88vk8VwIDAQABAkAe+Ug51nCQ+mg/494KnOQMe0spuhzsUlXf3nPd1b8jf07FNMUtn26+af6ix0JS8ubxvsW4k0CAexbPgqZ8eKh5AiEA6v4dFH3qiwiO6ZJrwsjv/nGc6OdUUTKAEsn9ir2EzbMCIQDn+PjX3v74yExiAEPpS+e8WEWPed+UfQSEi4FR4UNszQIgSeogL3/jtFO+0A00RBJ8GqaV2lRZGyktS7upOmieSy0CIAEoTd2N4EnPfgoTJEOavFpFOnufyQzNnuMmawvEiviVAiEAzPiFl/AEkpDK6ag4oTSpj2uCSU8kbaMrDZdpU3ZE3xc="
	cert := "MIIBVAIBADANBgkqhkiG9w0BAQEFAASCAT4wggE6AgEAAkEA9zMlWyFXxbKtX4aiKFufSeoKbn8fi0WCcEEMwlVZ/JrOqD5fBSnIhg55HfxWSGz+/UErN7uOM8wLUiO/h52ctQIDAQABAkEA6lhZVKAyZmAyDfpDZgEd6+Nsyj3HeK+HrcVtRBnH+jgnLDzhUQOLXgmOlMcqY6xgb+N5/lujZyI0i2U+QskvrQIhAP5yvhRlXbKaZ3RUkIq5oKOOylsuftv5JCepKtQNjxJLAiEA+LUWOXeQzIQpzK1YZkLXuqqkzkG4zBPkZC3Gjo9wrP8CIHqQKp8IdKmjVlHiqf0ojQg7r51cOB8dUx0Hj2m1XkUhAiBkd3J04HHjy/e+sL+m/5V52HkewLnP5rE6Lquo5xTTLQIgZ6vK0a305ZhA9jJM0/8bsNIa90nW15G8brZuPDmW7sY="
	code, err := GenerateRSA256Token(cert, &myJWTClaims)

	if err != nil {
		applog.Error(err)
	}

	applog.Info(code)
}

// 生成测试用的 Server token
func TestGenerateRSA512Token(t *testing.T) {
	myJWTClaims := TestJWTPayload{
		StandardClaims: jwtgo.StandardClaims{
			Issuer:    "30001",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 168).Unix(),
		},
	}

	// cert := "MIIBVAIBADANBgkqhkiG9w0BAQEFAASCAT4wggE6AgEAAkEA1O/W1f3BTpMIQI9z42YJVA/+2Igq7KwIVIuiVzYw2qpClaWMnKBfzObso0CbP953lxiRHyQVDI5tdSs88vk8VwIDAQABAkAe+Ug51nCQ+mg/494KnOQMe0spuhzsUlXf3nPd1b8jf07FNMUtn26+af6ix0JS8ubxvsW4k0CAexbPgqZ8eKh5AiEA6v4dFH3qiwiO6ZJrwsjv/nGc6OdUUTKAEsn9ir2EzbMCIQDn+PjX3v74yExiAEPpS+e8WEWPed+UfQSEi4FR4UNszQIgSeogL3/jtFO+0A00RBJ8GqaV2lRZGyktS7upOmieSy0CIAEoTd2N4EnPfgoTJEOavFpFOnufyQzNnuMmawvEiviVAiEAzPiFl/AEkpDK6ag4oTSpj2uCSU8kbaMrDZdpU3ZE3xc="
	cert := "MIIBVQIBADANBgkqhkiG9w0BAQEFAASCAT8wggE7AgEAAkEAnfApRhGYeOAlVHLY4NSFO4ZuvxJvKTPeBh04UlcaAOr+scNdbyqbk8Vf5dMbt4isNmzcwxFAucrc4Kk9A+487QIDAQABAkBCgPogMxCM9XVClgmUwmhHpFfFxTFhnCGFWZB5u4uLjr7RcTUE32TD1kdfoPOpyvLO2FXe3elRfzhv/Zs8oUxRAiEAyOObvS2zKT5SyDpaCsnD4IPrPmcXWlJ+TNufHLcYVUMCIQDJRB3+FGDCV95SJqogv3joGr9SEzYyZ15IHS1ZBAbqDwIhAJltJHPepFfajpT34+8bW+tsNU2KpM5J9MY0chgXSpFLAiAYcTxPOfxSPceTJHzuwzIUxgLMHecsEd+iVGDIebO5QwIhAItZOhxgnE88y1DeJXNwp5VEGZ14H48nCQAgwd485rsj"
	code, err := GenerateRSA256Token(cert, &myJWTClaims)

	if err != nil {
		applog.Error(err)
	}

	applog.Info(code)
}

type AppJWTPayload struct {
	jwtgo.StandardClaims

	UID           string         `json:"uid,omitempty"`
	Mode          string         `json:"mode,omitempty"`
	Scheme        AuthSchemeEnum `json:"scheme,omitempty"`
	OAuthProvider string         `json:"oap,omitempty"` //oap: OAuth Provider
	SessionID     string         `json:"sid,omitempty"`
	Sign          string         `json:"sig,omitempty"`
	SignAlg       string         `json:"sal,omitempty"`
	Extra         string         `json:"extra,omitempty"` // 额外的字段，可以存放 json 等非标数据
}

func TestParseClaimsUnverified(t *testing.T) {
	a := AppJWTPayload{}
	tstring := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiIzMDAxNSIsImV4cCI6NjY0MTY5ODQzMywianRpIjoiZTY5NzdlNDMtYTg3Zi00NDgxLWFkOTAtM2E4NTMyYTU5OGU0IiwiaWF0IjoxNTkyODk5MDc1LCJpc3MiOiIzMDAwMiIsIm9hcCI6Im1peGluIiwic2lkIjoiMzhiMmZiODgwMmM0NGI1Nzk1MzZmYjMwZTFkYjU0N2QifQ.s_7R2jeau7p6UPuvDQnvDufeRfBEkMIS9jsUdHtpP9o"
	ParseClaimsUnverified(tstring, &a)
}
