package authorization

import (
	"crypto/x509"
	"encoding/base64"
	"errors"

	applog "github.com/lanvige/goutils/logger"

	jwtgo "github.com/dgrijalva/jwt-go"
)

// ==== Verify Operations ===== //
// ==== Verify Operations ===== //

// verifyClaimsHS parseClaims
func verifyClaimsHS(token, secret string, payload jwtgo.Claims) error {
	tokenClaims, err := jwtgo.ParseWithClaims(token, payload, func(token *jwtgo.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if nil != err {
		return err
	}

	if tokenClaims != nil {
		// if _, ok := tokenClaims.Claims.(*AppJWTPayload); ok && tokenClaims.Valid {
		if tokenClaims.Valid {
			return nil
		}

		return errors.New("TokenClaims Not Valid")
	}

	return errors.New("TokenClaims Not Valid")
}

// verifyClaimsRS parseClaims
func verifyClaimsRS(tokenStr, pubkey string, payload jwtgo.Claims) error {
	//
	// 方案一，使用 jwt ParseFromPEM：
	// 	pem := `
	// -----BEGIN PUBLIC KEY-----
	// MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBANTv1tX9wU6TCECPc+NmCVQP/tiIKuysCFSLolc2MNqqQpWljJygX8zm7KNAmz/ed5cYkR8kFQyObXUrPPL5PFcCAwEAAQ==
	// -----END PUBLIC KEY-----
	// `
	// 	parsedKey, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(pem))

	//
	// 方案二，使用 X509：
	// pem := `MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBANTv1tX9wU6TCECPc+NmCVQP/tiIKuysCFSLolc2MNqqQpWljJygX8zm7KNAmz/ed5cYkR8kFQyObXUrPPL5PFcCAwEAAQ==`
	// decodestr, _ := base64.StdEncoding.DecodeString(pem)

	decodestr, _ := base64.StdEncoding.DecodeString(pubkey)

	var errParsePubkey error
	var parsedKey interface{}
	if parsedKey, errParsePubkey = x509.ParsePKIXPublicKey([]byte(decodestr)); errParsePubkey != nil {
		applog.Errorf("Parse public key error: %v", errParsePubkey)
		return errParsePubkey
	}

	tokenClaims, errParseClaims := jwtgo.ParseWithClaims(tokenStr, payload, func(token *jwtgo.Token) (interface{}, error) {
		return parsedKey, nil
	})

	if nil != errParseClaims {
		applog.Errorf("Parse claims error: %v", errParseClaims)
		return errParseClaims
	}

	if tokenClaims != nil {
		// if _, ok := tokenClaims.Claims.(*AppJWTPayload); ok && tokenClaims.Valid {
		if tokenClaims.Valid {
			return nil
		}

		return errors.New("TokenClaims Not Valid")
	}

	return errors.New("TokenClaims Not Valid")
}

// ==== Parse Operations ===== //
// ==== Parse Operations ===== //

// parseClaimsUnverified parseClaimsUnverified
func parseTokenUnverified(tokenStr string, payload jwtgo.Claims) (*jwtgo.Token, error) {
	token, _, err := new(jwtgo.Parser).ParseUnverified(tokenStr, payload)
	if nil != err {
		return nil, err
	}

	if token != nil {
		// if _, ok := token.Claims.(*AppJWTPayload); ok {
		// 	return token, nil
		// }
		return token, nil
		// return nil, errors.New("TokenClaims Not Valid")
	}

	return nil, errors.New("TokenClaims Not Valid")
}

// parseClaimsUnverified parseClaimsUnverified
func parseClaimsUnverified(tokenStr string, payload jwtgo.Claims) (*jwtgo.Token, jwtgo.Claims, error) {
	token, _, err := new(jwtgo.Parser).ParseUnverified(tokenStr, payload)
	if nil != err {
		return nil, nil, err
	}

	if token != nil {
		// if claims, ok := token.Claims.(*AppJWTPayload); ok {
		// 	return token, claims, nil
		// }
		return token, token.Claims, nil
		// return nil, nil, errors.New("TokenClaims Not Valid")
	}

	return nil, nil, errors.New("TokenClaims Not Valid")
}

// ==== Generate Operations ===== //
// ==== Generate Operations ===== //

// GenerateHS256Token GenerateHSToken
func generateHS256Token(secret string, payload jwtgo.Claims) (string, error) {
	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, payload)

	return token.SignedString([]byte(secret))
}

// GenerateRSA256Token GenerateRSAToken
func generateRSA256Token(prikey string, payload jwtgo.Claims) (string, error) {
	token := jwtgo.NewWithClaims(jwtgo.SigningMethodRS256, payload)

	// Decode
	decodestr, errDec := base64.StdEncoding.DecodeString(prikey)
	if nil != errDec {
		applog.Errorf("解析 prikey 出错 %v", errDec)
		return "", errDec
	}

	var err error
	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKCS8PrivateKey(decodestr); err != nil {
		applog.Errorf("Parse cert 出错 %v", err)
		return "", err
	}

	return token.SignedString(parsedKey)
}

// GenerateRSA512Token GenerateRSAToken
func generateRSA512Token(appID, prikey string, payload jwtgo.Claims) (string, error) {
	token := jwtgo.NewWithClaims(jwtgo.SigningMethodRS512, payload)

	// Decode
	decodestr, errDec := base64.StdEncoding.DecodeString(prikey)
	if nil != errDec {
		applog.Errorf("解析 prikey 出错 %v", errDec)
		return "", errDec
	}

	var err error
	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKCS8PrivateKey(decodestr); err != nil {
		return "", err
	}

	return token.SignedString(parsedKey)
}
