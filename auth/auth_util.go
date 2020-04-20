package authorization

import (
	"errors"
	"strings"
)

// ==== public Operations ===== //
// ==== public Operations ===== //

// parseAuthScheme parseAuthScheme
func parseAuthScheme(authStr string) (AuthSchemeEnum, string, error) {
	if len(authStr) < 10 {
		return AuthSchemeEnumOther, "", errors.New("TokenClaims Not Valid")
	}

	s := strings.Split(authStr, " ")
	if len(s) != 2 {
		return AuthSchemeEnumOther, "", errors.New("TokenClaims Not Valid")
	}

	// 是 哪种类型 Basic/Bearer
	schemeStr := s[0]
	tokenStr := s[1]
	schemeType := NewAuthSchemeEnum(schemeStr)

	return schemeType, tokenStr, nil
}

// ==== Type Definition ===== //
// ==== Type Definition ===== //

// AuthSchemeEnum 枚举
type AuthSchemeEnum string

const (
	// AuthSchemeEnumBasic Basic
	AuthSchemeEnumBasic AuthSchemeEnum = "Basic"
	// AuthSchemeEnumDigest Digest
	AuthSchemeEnumDigest AuthSchemeEnum = "Digest"
	// AuthSchemeEnumBearer Bearer
	AuthSchemeEnumBearer AuthSchemeEnum = "Bearer"
	// AuthSchemeEnumOther unkonwn
	AuthSchemeEnumOther AuthSchemeEnum = "Unkonwn"
)

// String String
func (e AuthSchemeEnum) String() string {
	switch e {
	case AuthSchemeEnumBasic:
		return "Basic"
	case AuthSchemeEnumDigest:
		return "Digest"
	case AuthSchemeEnumBearer:
		return "Bearer"
	default:
		return "Unkonwn"
	}
}

// NewAuthSchemeEnum NewAuthSchemeEnum
func NewAuthSchemeEnum(scheme string) AuthSchemeEnum {
	lowerscheme := strings.ToLower(scheme)
	switch lowerscheme {
	case "basic":
		return AuthSchemeEnumBasic
	case "digest":
		return AuthSchemeEnumDigest
	case "bearer":
		return AuthSchemeEnumBearer
	default:
		return AuthSchemeEnumOther
	}
}
