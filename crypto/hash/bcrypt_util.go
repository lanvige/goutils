package hash

import (
	"golang.org/x/crypto/bcrypt"
)

// ==== bcrypt Operations ===== //
// ==== bcrypt Operations ===== //

// GenPasswordHash 生成 Hash
func GenPasswordHash(password string) (string, error) {
	// Generate "hash" to store from user password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashPassword), nil
}

// CheckPasswordHash CheckPasswordHash
func CheckPasswordHash(hash, password string) error {
	passwordHashByte := []byte(hash)
	passwordByte := []byte(password)

	return bcrypt.CompareHashAndPassword(passwordHashByte, passwordByte)
}
