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

// ParseIntToStr ParseIntToStr
func ParseIntToStr(value int) string {
	return strconv.Itoa(value)
}

// ParseUintToStr ParseUintToStr
func ParseUintToStr(value uint) string {
	ui64value := uint64(value)
	return strconv.FormatUint(ui64value, 10)
}

// ParseInt32ToStr ParseInt32ToStr
func ParseInt32ToStr(value int32) string {
	i64value := int64(value)
	return strconv.FormatInt(i64value, 10)
}

// ParseUint32ToStr ParseUint32ToStr
func ParseUint32ToStr(value uint32) string {
	ui64value := uint64(value)
	return strconv.FormatUint(ui64value, 10)
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

// ParseStrToInt32 ParseStrToInt32
func ParseStrToInt32(value string, defaultValue int32) int32 {
	intValue, err := strconv.Atoi(value)

	if err != nil {
		// logrus.Error(err)
		return defaultValue
	}

	return int32(intValue)
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

// ==== CONVERT INT/BYTE Operations ===== //
// ==== CONVERT INT/BYTE Operations ===== //

// ParseIntToBytes 整形转换成字节
func ParseIntToBytes(num int) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, num)
	return bytesBuffer.Bytes()
}

// ParseInt32ToBytes 整形转换成字节
func ParseInt32ToBytes(num int32) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, num)
	return bytesBuffer.Bytes()
}

// ParseInt64ToBytes 整形转换成字节
func ParseInt64ToBytes(num int64) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, num)
	return bytesBuffer.Bytes()
}

// ParseUintToBytes 整形转换成字节
func ParseUintToBytes(num uint) []byte {
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

// ==== CONVERT BYTE/INT Operations ===== //
// ==== CONVERT BYTE/INT Operations ===== //

// ParseBytesToInt 字节转换成整形
func ParseBytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}

// ParseBytesToInt32 字节转换成整形
func ParseBytesToInt32(b []byte) int32 {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int32(x)
}

// ParseBytesToInt64 字节转换成整形
func ParseBytesToInt64(b []byte) int64 {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int64(x)
}

// ParseBytesToUint 字节转换成整形
func ParseBytesToUint(b []byte) uint {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return uint(x)
}

// ParseBytesToUint32 字节转换成整形
func ParseBytesToUint32(b []byte) uint32 {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return uint32(x)
}

// ParseBytesToUint64 字节转换成整形
func ParseBytesToUint64(b []byte) uint64 {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return uint64(x)
}
