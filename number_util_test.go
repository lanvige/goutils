package utils

import (
	"testing"

	applog "github.com/lanvige/logger"
)

func TestIsDigitPresent(t *testing.T) {
	code := IsDigitPresent(143, 3)

	applog.Error(code)
}

func TestParseInt32ToStr(t *testing.T) {
	code := ParseInt32ToStr(13)

	applog.Info(code)
}

func TestParseInt64ToStr(t *testing.T) {
	code := ParseInt64ToStr(41513515623443)

	applog.Info(code)
}

func TestParseIntToBytes(t *testing.T) {
	var abc int = 41156234
	code := ParseIntToBytes(abc)

	applog.Info(code)
}

func TestParseUint32ToBytes(t *testing.T) {
	var abc uint32 = 41156234
	code := ParseUint32ToBytes(abc)

	applog.Info(code)
}
