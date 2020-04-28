package utils

import (
	"encoding/base64"
)

// Base64Decode Base64Decode
func Base64Decode(value string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(value)
	if nil != err {
		return "", err
	}

	return string(decoded), nil
}
