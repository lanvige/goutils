package utils

import (
	"bytes"
	"encoding/binary"
	"strconv"
)

// IsDigitPresent IsDigitPresent
func IsDigitPresent(number, d int) bool {
	// Breal loop if d is present as digit
	for number > 0 {
		if number%10 == d {
			break
		}

		number = number / 10
	}

	// If loop broke
	return (number > 0)
}

// IsDigitPresentUint64 IsDigitPresentUint64
func IsDigitPresentUint64(number uint64, d uint64) bool {
	// Breal loop if d is present as digit
	for number > 0 {
		if number%10 == d {
			break
		}

		number = number / 10
	}

	// If loop broke
	return (number > 0)
}

// ==== CONVERT String Operations ===== //
// ==== CONVERT Operations ===== //

// ParseInt32ToStr ParseInt32ToStr
func ParseInt32ToStr(value int32) string {
	i64value := int64(value)
	return strconv.FormatInt(i64value, 10)
}

// ParseInt64ToStr ParseInt64ToStr
func ParseInt64ToStr(value int64) string {
	return strconv.FormatInt(value, 10)
}

// ParseUint64ToStr ParseUint64ToStr
func ParseUint64ToStr(value uint64) string {
	return strconv.FormatUint(value, 10)
}

// ==== CONVERT Number Operations ===== //
// ==== CONVERT Number Operations ===== //

// ParseStrToInt ParseInt
func ParseStrToInt(value string, defaultValue int) int {
	intValue, err := strconv.Atoi(value)

	if err != nil {
		// logrus.Error(err)
		return defaultValue
	}

	return intValue
}

// ParseStrToInt64 ParseInt64
func ParseStrToInt64(value string, defaultValue int64) int64 {
	intValue, err := strconv.ParseInt(value, 10, 64)

	if err != nil {
		// logrus.Error(err)
		return defaultValue
	}

	return intValue
}

// ParseStrToUint ParseUint
func ParseStrToUint(value string, defaultValue uint) uint {
	intValue, err := strconv.ParseUint(value, 10, 0)

	if err != nil {
		// logrus.Error(err)
		return defaultValue
	}

	return uint(intValue)
}

// ParseStrToUint32 ParseUint
func ParseStrToUint32(value string, defaultValue uint32) uint32 {
	intValue, err := strconv.ParseUint(value, 10, 32)

	if err != nil {
		// logrus.Error(err)
		return defaultValue
	}

	return uint32(intValue)
}

// ParseStrToUint64 ParseUint64
func ParseStrToUint64(value string, defaultValue uint64) uint64 {
	intValue, err := strconv.ParseUint(value, 10, 64)

	if err != nil {
		// logrus.Error(err)
		return defaultValue
	}

	return intValue
}

// ==== CONVERT BYTE Operations ===== //
// ==== CONVERT BYTE Operations ===== //

// ParseIntToBytes 整形转换成字节
func ParseIntToBytes(num int) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, num)
	return bytesBuffer.Bytes()
}

// ParseUint32ToBytes 整形转换成字节
func ParseUint32ToBytes(num uint32) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, num)

	return bytesBuffer.Bytes()
}

// ParseUint64ToBytes 整形转换成字节
func ParseUint64ToBytes(num uint64) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, num)

	return bytesBuffer.Bytes()
}

// ParseBytesToInt 字节转换成整形
func ParseBytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}
