package authorization

import (
	"encoding/base64"
	"fmt"
	"testing"
	"time"

	applog "github.com/lanvige/goutils/logger"

	jwtgo "github.com/dgrijalva/jwt-go"
	goutils "github.com/lanvige/goutils"
)

// Define
// AppJWTPayload AppJWTPayload
type AppJWTPayload struct {
	jwtgo.StandardClaims
	// Provider string `json:"oap,omitempty"` //oap: OAuth Provider
	Mode      string `json:"mode,omitempty"`
	SessionID string `json:"sid,omitempty"`
	Sign      string `json:"sig,omitempty"`
	SignAlg   string `json:"sal,omitempty"`
}

// BASIC

func TestGenerateBasicToken(t *testing.T) {
	keyStr := goutils.UUIDV4StringGen()
	secretStr := goutils.UUIDV4StringGen()

	tokenStr := fmt.Sprintf("%s:%s", keyStr, secretStr)
	tokenB64Str := base64.StdEncoding.EncodeToString([]byte(tokenStr))

	applog.Info(keyStr)
	applog.Info(secretStr)
	applog.Info(tokenStr)
	applog.Info(tokenB64Str)
}

func TestGenerateHSToken(t *testing.T) {
	myJWTClaims := AppJWTPayload{
		StandardClaims: jwtgo.StandardClaims{
			Issuer:    "30001",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
		// Provider: provider,
	}

	code, _ := generateHS256Token("abcdddd", &myJWTClaims)

	applog.Info(code)
}

func TestGenerateRSA256Token(t *testing.T) {
	myJWTClaims := AppJWTPayload{
		StandardClaims: jwtgo.StandardClaims{
			Issuer:    "30001",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 168).Unix(),
		},
	}

	// cert := "MIIBVAIBADANBgkqhkiG9w0BAQEFAASCAT4wggE6AgEAAkEA1O/W1f3BTpMIQI9z42YJVA/+2Igq7KwIVIuiVzYw2qpClaWMnKBfzObso0CbP953lxiRHyQVDI5tdSs88vk8VwIDAQABAkAe+Ug51nCQ+mg/494KnOQMe0spuhzsUlXf3nPd1b8jf07FNMUtn26+af6ix0JS8ubxvsW4k0CAexbPgqZ8eKh5AiEA6v4dFH3qiwiO6ZJrwsjv/nGc6OdUUTKAEsn9ir2EzbMCIQDn+PjX3v74yExiAEPpS+e8WEWPed+UfQSEi4FR4UNszQIgSeogL3/jtFO+0A00RBJ8GqaV2lRZGyktS7upOmieSy0CIAEoTd2N4EnPfgoTJEOavFpFOnufyQzNnuMmawvEiviVAiEAzPiFl/AEkpDK6ag4oTSpj2uCSU8kbaMrDZdpU3ZE3xc="
	cert := "MIIBVAIBADANBgkqhkiG9w0BAQEFAASCAT4wggE6AgEAAkEA9zMlWyFXxbKtX4aiKFufSeoKbn8fi0WCcEEMwlVZ/JrOqD5fBSnIhg55HfxWSGz+/UErN7uOM8wLUiO/h52ctQIDAQABAkEA6lhZVKAyZmAyDfpDZgEd6+Nsyj3HeK+HrcVtRBnH+jgnLDzhUQOLXgmOlMcqY6xgb+N5/lujZyI0i2U+QskvrQIhAP5yvhRlXbKaZ3RUkIq5oKOOylsuftv5JCepKtQNjxJLAiEA+LUWOXeQzIQpzK1YZkLXuqqkzkG4zBPkZC3Gjo9wrP8CIHqQKp8IdKmjVlHiqf0ojQg7r51cOB8dUx0Hj2m1XkUhAiBkd3J04HHjy/e+sL+m/5V52HkewLnP5rE6Lquo5xTTLQIgZ6vK0a305ZhA9jJM0/8bsNIa90nW15G8brZuPDmW7sY="
	code, err := generateRSA256Token(cert, &myJWTClaims)

	if err != nil {
		applog.Error(err)
	}

	applog.Info(code)
}

// 生成测试用的 Server token
func TestGenerateRSA512Token(t *testing.T) {
	myJWTClaims := AppJWTPayload{
		StandardClaims: jwtgo.StandardClaims{
			Issuer:    "30001",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 168).Unix(),
		},
	}

	// cert := "MIIBVAIBADANBgkqhkiG9w0BAQEFAASCAT4wggE6AgEAAkEA1O/W1f3BTpMIQI9z42YJVA/+2Igq7KwIVIuiVzYw2qpClaWMnKBfzObso0CbP953lxiRHyQVDI5tdSs88vk8VwIDAQABAkAe+Ug51nCQ+mg/494KnOQMe0spuhzsUlXf3nPd1b8jf07FNMUtn26+af6ix0JS8ubxvsW4k0CAexbPgqZ8eKh5AiEA6v4dFH3qiwiO6ZJrwsjv/nGc6OdUUTKAEsn9ir2EzbMCIQDn+PjX3v74yExiAEPpS+e8WEWPed+UfQSEi4FR4UNszQIgSeogL3/jtFO+0A00RBJ8GqaV2lRZGyktS7upOmieSy0CIAEoTd2N4EnPfgoTJEOavFpFOnufyQzNnuMmawvEiviVAiEAzPiFl/AEkpDK6ag4oTSpj2uCSU8kbaMrDZdpU3ZE3xc="
	cert := "MIIBVQIBADANBgkqhkiG9w0BAQEFAASCAT8wggE7AgEAAkEAnfApRhGYeOAlVHLY4NSFO4ZuvxJvKTPeBh04UlcaAOr+scNdbyqbk8Vf5dMbt4isNmzcwxFAucrc4Kk9A+487QIDAQABAkBCgPogMxCM9XVClgmUwmhHpFfFxTFhnCGFWZB5u4uLjr7RcTUE32TD1kdfoPOpyvLO2FXe3elRfzhv/Zs8oUxRAiEAyOObvS2zKT5SyDpaCsnD4IPrPmcXWlJ+TNufHLcYVUMCIQDJRB3+FGDCV95SJqogv3joGr9SEzYyZ15IHS1ZBAbqDwIhAJltJHPepFfajpT34+8bW+tsNU2KpM5J9MY0chgXSpFLAiAYcTxPOfxSPceTJHzuwzIUxgLMHecsEd+iVGDIebO5QwIhAItZOhxgnE88y1DeJXNwp5VEGZ14H48nCQAgwd485rsj"
	code, err := generateRSA256Token(cert, &myJWTClaims)

	if err != nil {
		applog.Error(err)
	}

	applog.Info(code)
}
