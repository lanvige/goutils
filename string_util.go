package utils

import (
	"encoding/binary"
	"fmt"
	"log"
	"math/rand"
	"time"
)

// GenRandomString GenRandomString
func GenRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	bytes := []byte(str)
	result := []byte{}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}

	return string(result)
}

// GenRandomNumString GenRandomNumString
func GenRandomNumString(l int) string {
	str := "0123456789"
	bytes := []byte(str)
	result := []byte{}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}

	return string(result)
}

// GenRandomSixDigital GenRandomSixDigital
func GenRandomSixDigital() string {
	var b [8]byte
	_, err := rand.Read(b[:])
	if err != nil {
		log.Panicln(err)
	}

	c := binary.LittleEndian.Uint64(b[:]) % 1000000
	if c < 100000 {
		c = 100000 + c
	}

	return fmt.Sprintf("%d", c)
}

// DefaultString DefaultString
func DefaultString(value, defaultValue string) string {
	if "" == value {
		return defaultValue
	}

	return value
}
