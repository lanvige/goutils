package utils

import (
	"testing"

	applog "github.com/lanvige/goutils/logger"

	libphonenumber "github.com/ttacon/libphonenumber"
)

func TestCNPhoneSplit(t *testing.T) {
	code, number, err := SplitPhone("+8612300001234")

	if nil != err {
		applog.Error(code)
		applog.Error(number)
	} else {
		applog.Error(code)
		applog.Error(number)
		applog.Error("normal")
	}
}

func TestUS2PhoneSplit(t *testing.T) {
	code, number, err := SplitPhone("+14246037372")

	if nil != err {
		applog.Error(code)
		applog.Error(number)
	} else {
		applog.Error(code)
		applog.Error(number)
		applog.Error("normal")
	}
}

func TestUSPhoneSplit(t *testing.T) {
	// num := "+14246037372"
	num := "+8612300001234"
	num2, _ := libphonenumber.Parse(num, "CN")

	natSigNumber := libphonenumber.GetNationalSignificantNumber(num2)
	geoCodeLength := libphonenumber.GetLengthOfGeographicalAreaCode(num2)
	abcd := libphonenumber.GetRegionCodeForNumber(num2)
	numberType := libphonenumber.GetNumberType(num2)
	countryCode := libphonenumber.GetCountryCodeForRegion(abcd)
	cc := num2.GetItalianLeadingZero()

	// Extract the area code.
	areaCode := ""
	if geoCodeLength > 0 {
		areaCode = natSigNumber[0:geoCodeLength]
	}

	applog.Error(natSigNumber)
	applog.Error(geoCodeLength)
	applog.Error(abcd)
	applog.Error(numberType)
	applog.Error(areaCode)
	applog.Error(countryCode)
	applog.Error(cc)

	code, number, err2 := SplitPhone(num)

	if nil != err2 {
		applog.Error(code)
		applog.Error(number)
	} else {
		applog.Error("normal")
	}
}
