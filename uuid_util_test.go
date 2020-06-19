package utils

import (
	"testing"

	applog "github.com/lanvige/goutils/logger"
)

func TestHex2UUIDGen(t *testing.T) {
	code := UUIDV1HexGen()

	// be471f7b040043989ec108b60d2b48c7
	applog.Error(code)
}

func TestHexUUIDGen(t *testing.T) {
	code := UUIDV1HexGen()

	// be471f7b040043989ec108b60d2b48c7
	applog.Error(code)
}

func TestUUIDGen(t *testing.T) {
	code := UUIDV1StringGen()

	// 20db2fb8-3ad6-4bdd-9504-84b07f833f6d
	applog.Error(code)
}

func TestSettlement1UUIDGen(t *testing.T) {
	code := UUIDHexGenFromString("sss")

	// 4d73a7c5-539e-3d3f-b685-41de3441be7d
	applog.Error(code)
}
