package utils

import (
	"testing"

	applog "github.com/lanvige/goutils/logger"
)

func TestRandomNumGen(t *testing.T) {
	code := GenRandomNumString(10)

	applog.Error(code)
}

func TestRandomStringGen(t *testing.T) {
	code := GenRandomString(100)

	applog.Error(code)
}

func TestRad6StringGen(t *testing.T) {
	code := GenRandomSixDigital()

	applog.Error(code)
}
