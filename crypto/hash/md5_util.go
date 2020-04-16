package hash

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

// ==== MD5 Operations ===== //
// ==== MD5 Operations ===== //

// MD5Encode MD5
func MD5Encode(data string) []byte {
	h := md5.New()
	h.Write([]byte(data))

	return h.Sum(nil)
}

// MD5Hex MD5Hex
func MD5Hex(str string) string {
	cipherStr := MD5Encode(str)

	return hex.EncodeToString(cipherStr)
}

// MD5Base64 MD5Base64
func MD5Base64(str string) string {
	cipherStr := MD5Encode(str)

	return base64.StdEncoding.EncodeToString(cipherStr)
}

// MD5HexShort str.substring(8, 24);
func MD5HexShort(str string) string {
	data := MD5Hex(str)

	return data[8:24]
}
