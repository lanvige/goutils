package crypto

import (
	"bytes"
	"encoding/binary"
)

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

// BytesToInt 字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}
