package utils

import (
	"testing"

	applog "github.com/lanvige/goutils/logger"
)

func TestHexUUIDGen(t *testing.T) {
	code := UUIDV4StringGen()

	// be471f7b040043989ec108b60d2b48c7
	applog.Error(code)
}

func TestUUIDGen(t *testing.T) {
	code := UUIDV1Gen()

	// 20db2fb8-3ad6-4bdd-9504-84b07f833f6d
	applog.Error(code)
}

func TestSettlementUUIDGen(t *testing.T) {
	code := SettlementID("sss", "ssss")

	// 20db2fb8-3ad6-4bdd-9504-84b07f833f6d
	applog.Error(code)
}
