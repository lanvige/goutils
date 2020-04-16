package utils

import (
	apperr "github/lanvige/goutils/error"

	"github.com/ttacon/libphonenumber"
)

// SplitPhone SplitPhone
// +8615618297195",
func SplitPhone(phone string) (string, string, apperr.Error) {
	num2, err := libphonenumber.Parse(phone, "CN")

	if nil != err {
		return "", "", apperr.UserPhoneIncorrect
	}

	countryCode := num2.GetCountryCode()
	nationNumber := num2.GetNationalNumber()

	countryCodeStr := ParseInt32ToStr(countryCode)
	nationNumberStr := ParseUint64ToStr(nationNumber)

	return countryCodeStr, nationNumberStr, nil
}
