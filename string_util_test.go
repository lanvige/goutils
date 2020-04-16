package utils

import (
	"testing"

	applog "goutils/logger"
)

func TestRandomNumGen(t *testing.T) {
	code := GetRandomNum(10)

	applog.Error(code)
}

func TestRandomStringGen(t *testing.T) {
	code := GetRandomString(100)

	applog.Error(code)
}
