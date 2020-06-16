package utils

import (
	"encoding/binary"
	"fmt"
	"log"
	"math/rand"
	"time"
)

// func genSymbol(length int) []string {
// 	charStr := "+=-@#~,.[]()!%^*$"

// 	var genstr := []string

// 	//遍历，生成一个随机index索引,
// 	for i := 0; i < length; i++ {
// 		index := rand.Intn(len(charStr))
// 		genstr[i] = charStr[index]
// 	}

// 	return string(genstr)
// }

// func genDigit(length int) string {
// 	charStr := "0123456789"
// 	var genstr []byte = make([]byte, length, length)

// 	//遍历，生成一个随机index索引,
// 	for i := 0; i < length; i++ {
// 		index := rand.Intn(len(charStr))
// 		genstr[i] = charStr[index]
// 	}

// 	return string(genstr)
// }

// func genChar(length int) string {
// 	charStr := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// 	var passwd []byte = make([]byte, length, length)

// 	//遍历，生成一个随机index索引,
// 	for i := 0; i < length; i++ {
// 		index := rand.Intn(len(charStr))
// 		passwd[i] = charStr[index]
// 	}

// 	return string(passwd)
// }

// // GenPassword GenPassword
// func GenPassword(length, digitLength, symbolLength uint32) string {
// 	digitStr = "0123456789"
// 	charStr := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
// 	symbolStr := "+=-@#~,.[]()!%^*$"

// 	bytes := []byte(str)
// 	result := []byte{}

// 	r := rand.New(rand.NewSource(time.Now().UnixNano()))

// 	for i := 0; i < l; i++ {
// 		result = append(result, bytes[r.Intn(len(bytes))])
// 	}

// 	return string(result)
// }

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
