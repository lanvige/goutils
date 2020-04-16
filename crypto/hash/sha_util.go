package hash

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

// ==== SHA1 Operations ===== //
// ==== SHA1 Operations ===== //

// SHA1 SHA1
func SHA1(data string) []byte {
	sha1 := sha1.New()
	sha1.Write([]byte(data))

	return sha1.Sum([]byte(""))
}

// SHA1Hex SHA1Hex
func SHA1Hex(data string) string {
	cipherStr := SHA1(data)

	return hex.EncodeToString(cipherStr)
}

// SHA1Base64 SHA1Base64
func SHA1Base64(data string) string {
	cipherStr := SHA1(data)

	return base64.StdEncoding.EncodeToString(cipherStr)
}

// ==== SHA256 Operations ===== //
// ==== SHA2 Operations ===== //

// SHA256 SHA256
func SHA256(data string) []byte {
	sha256Value := sha256.New()
	sha256Value.Write([]byte(data))

	return sha256Value.Sum([]byte(""))
}

// SHA256Hex SHA256Hex
func SHA256Hex(data string) string {
	cipherStr := SHA256(data)

	return hex.EncodeToString(cipherStr)
}

// SHA256Base64 SHA256Base64
func SHA256Base64(data string) string {
	cipherStr := SHA256(data)

	return base64.StdEncoding.EncodeToString(cipherStr)
}

// ==== SHA512 Operations ===== //
// ==== SHA2 Operations ===== //

// SHA512 SHA512
func SHA512(data string) []byte {
	shaValue := sha256.New()
	shaValue.Write([]byte(data))

	return shaValue.Sum([]byte(""))
}

// SHA512Hex SHA512Hex
func SHA512Hex(data string) string {
	cipherStr := SHA512(data)

	return hex.EncodeToString(cipherStr)
}

// SHA512Base64 SHA512Base64
func SHA512Base64(data string) string {
	cipherStr := SHA512(data)

	return base64.StdEncoding.EncodeToString(cipherStr)
}
