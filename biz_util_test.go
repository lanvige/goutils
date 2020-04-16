package utils

import (
	"testing"

	applog "github.com/lanvige/goutils/logger"
)

func TestBIZIDGen(t *testing.T) {
	code := OrderBizIDGen()

	applog.Info(code)
}

func TestShortBIZIDGen(t *testing.T) {
	code := OrderShortBizIDGen()

	applog.Info(code)
}

func TestOriIDID(t *testing.T) {
	code := OrderOriIDRestore("order-2019080812282674502")

	applog.Info(code)
}

func TestUserBizIDGen(t *testing.T) {
	code := UserBizIDGen()

	applog.Info(code)
}
