package authorization

import (
	"encoding/base64"
	"fmt"
	"testing"

	applog "github.com/lanvige/goutils/logger"

	goutils "github.com/lanvige/goutils"
)

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

func TestGenerateBasicToken1(t *testing.T) {
	keyStr := "05912316696dba4f"
	secretStr := "759123ba4f"

	tokenStr := fmt.Sprintf("%s:%s", keyStr, secretStr)
	tokenB64Str := base64.StdEncoding.EncodeToString([]byte(tokenStr))

	applog.Info(keyStr)
	applog.Info(secretStr)
	applog.Info(tokenStr)
	applog.Info(tokenB64Str)
}
