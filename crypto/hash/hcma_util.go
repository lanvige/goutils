package hash

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
)

// ==== HS1 Operations ===== //
// ==== HS/HMAC Operations ===== //

// HmacSha1 HMAC-SHA1
func HmacSha1(data, key []byte) []byte {
	h := hmac.New(sha1.New, []byte(key))
	h.Write([]byte(data))

	return h.Sum(nil)
}

// HS1Hex HMAC-SHA1
func HS1Hex(data, key []byte) string {
	cipherStr := HmacSha1(data, key)
	return hex.EncodeToString(cipherStr)
}

// HS1Base64 HMAC-SHA1
func HS1Base64(data, key []byte) string {
	cipherStr := HmacSha1(data, key)
	return base64.StdEncoding.EncodeToString(cipherStr)
}

// ==== HS256 Operations ===== //
// ==== HS2/HMAC Operations ===== //

// HmacSha256 HMAC-SHA256
func HmacSha256(data, key []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(data)

	return h.Sum(nil)
}

// HS256Hex HMAC-SHA256
func HS256Hex(data, key []byte) string {
	cipherStr := HmacSha256(data, key)

	return hex.EncodeToString(cipherStr)
}

// HS256Base64 HMAC-SHA256
func HS256Base64(data, key []byte) string {
	cipherStr := HmacSha256(data, key)
	return base64.StdEncoding.EncodeToString(cipherStr)
}

// ==== HS384 Operations ===== //
// ==== HS2/HMAC Operations ===== //

// HmacSha384 HMAC-SHA384
func HmacSha384(data, key []byte) []byte {
	h := hmac.New(sha512.New384, key)
	h.Write(data)

	return h.Sum(nil)
}

// HS384Hex HMAC-SHA384
func HS384Hex(data, key []byte) string {
	cipherStr := HmacSha384(data, key)

	return hex.EncodeToString(cipherStr)
}

// HS384Base64 HMAC-SHA384
func HS384Base64(data, key []byte) string {
	cipherStr := HmacSha384(data, key)

	return base64.StdEncoding.EncodeToString(cipherStr)
}

// ==== HS512 Operations ===== //
// ==== HS2/HMAC Operations ===== //

// HmacSha512 HMAC-SHA512
func HmacSha512(data, key []byte) []byte {
	h := hmac.New(sha512.New, key)
	h.Write(data)

	return h.Sum(nil)
}

// HS512Hex HMAC-SHA512
func HS512Hex(data, key []byte) string {
	cipherStr := HmacSha512(data, key)

	return hex.EncodeToString(cipherStr)
}

// HS512Base64 HMAC-SHA512
func HS512Base64(data, key []byte) string {
	cipherStr := HmacSha512(data, key)

	return base64.StdEncoding.EncodeToString(cipherStr)
}
