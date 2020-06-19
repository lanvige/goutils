package utils

import (
	"encoding/hex"

	uuid "github.com/google/uuid"
)

// UUIDStringGen UUIDStringGen
// b9be9a09-7117-4bdb-9c6d-737269c86480
func UUIDStringGen() string {
	return UUIDV4StringGen()
}

// UUIDV1StringGen UUIDV1StringGen
// b9be9a09-7117-4bdb-9c6d-737269c86480
func UUIDV1StringGen() string {
	uuid, _ := uuid.NewUUID()
	return uuid.String()
}

// UUIDV1HexGen UUIDV1HexGen
// 5106b5e58ee44f74a5d49e779dcf7f57
func UUIDV1HexGen() string {
	uuid, _ := uuid.NewUUID()
	return UUIDHex(uuid)
}

// UUIDV4StringGen UUIDV4StringGen
// b9be9a09-7117-4bdb-9c6d-737269c86480
func UUIDV4StringGen() string {
	return uuid.New().String()
}

// UUIDV4HexGen UUIDV4HexGen
// 5106b5e58ee44f74a5d49e779dcf7f57
func UUIDV4HexGen() string {
	uuid := uuid.New()
	return UUIDHex(uuid)
}

// UUIDHex UUIDHex
func UUIDHex(uuid uuid.UUID) string {
	hexuuid := make([]byte, 32)
	hex.Encode(hexuuid, uuid[:])

	abc, _ := uuid.MarshalBinary()
	return string(abc)
}

// UUIDParseFromString UUIDParseFromString
func UUIDParseFromString(value string) (uuid.UUID, error) {
	uid, err := uuid.Parse(value)

	if nil != err {
		return uuid.Nil, err
	}

	return uid, nil
}

// UUIDHexGenFromString 使用字符串生成固定的 UUID
func UUIDHexGenFromString(value string) string {
	return UUIDHex(UUIDGenFromString(value))
}

// UUIDStringGenFromString 使用字符串生成固定的 UUID
func UUIDStringGenFromString(value string) string {
	return UUIDGenFromString(value).String()
}

// UUIDGenFromString 使用字符串生成固定的 UUID
func UUIDGenFromString(value string) uuid.UUID {
	sum := []byte(value)

	return uuid.NewMD5(uuid.Nil, sum)
}
