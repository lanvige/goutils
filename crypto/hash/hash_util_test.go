package hash

import (
	"testing"

	goutils "github.com/lanvige/goutils"
	applog "github.com/lanvige/goutils/logger"
)

func TestCRC32Encode(t *testing.T) {
	code1 := CRC32Encode("30001")

	applog.Error(code1)
}

func TestCRC32Hex(t *testing.T) {
	code1 := CRC32Hex("30001")

	applog.Error(code1)
}

func TestMD5ShortGen(t *testing.T) {
	rdxString := cryptolib.GetRandomString(100)
	code1 := MD5Hex(rdxString)
	code2 := MD5HexShort(rdxString)
	code3 := MD5Base64(rdxString)

	applog.Error(code1)
	applog.Error(code2)
	applog.Error(code3)
}

// b010 ff32 d6470135ad28f21bfae7c409
// 0000 0000 d6470135ad28f21b
// sBD/MtZHATWtKPIb+ufECQ==

func TestSHA1Hex(t *testing.T) {
	content := "CN-IDCARD-10112319970712004X"
	code1 := SHA1Hex(content)
	code2 := SHA1Base64(content)

	applog.Error(code1)
	applog.Error(code2)
}

func TestHS256Hmac(t *testing.T) {
	rdxString := goutils.GetRandomString(100)
	rdxString1 := []byte(rdxString)

	code1 := HS256Hex(rdxString1, rdxString1)
	code2 := HS256Base64(rdxString1, rdxString1)

	applog.Error(code1)
	applog.Error(code2)
}

// d6ef58e99de7eca1fc3a478019e70d40d3da7c29aef57c76bc48ab5c4a3c8053
// 1u9Y6Z3n7KH8OkeAGecNQNPafCmu9Xx2vEirXEo8gFM=

func TestSHA1CHID(t *testing.T) {

	// country := "CN"
	// cardType := "IDCARD"
	// idNumber := "10112319970712004X"

	// content := fmt.Sprintf("%s-%s-%s", country, cardType, idNumber)

	content := "CN-IDCARD-10112319970712004X"
	applog.Error(content)

	code1 := SHA1Hex(content)

	applog.Error(code1)
}
