package utils

import (
	"strconv"
	"time"

	applog "github.com/lanvige/goutils/logger"
)

// OrderBizIDGen BIZIDGen
// 201907301746-544-0298
func OrderBizIDGen() string {
	now := time.Now()

	dataStr := now.Format("200601021504")
	nanosec := now.Nanosecond()
	nanosecStr := strconv.Itoa(nanosec)
	microsecStr := nanosecStr[0:3]
	rdmStr := GetRandomNum(4)

	// applog.Error(dataStr)
	// applog.Errorf("microsecStr:%v", microsecStr)
	// applog.Errorf("micrrdmStrosecStr:%v", rdmStr)

	str := dataStr + microsecStr + rdmStr

	return str
}

// OrderShortBizIDGen BIZIDGen
// 1907301746-54-029
func OrderShortBizIDGen() string {
	now := time.Now()

	dataStr := now.Format("060102150405")
	nanosec := now.Nanosecond()
	nanosecStr := strconv.Itoa(nanosec)
	centisecStr := nanosecStr[0:2]
	rdmStr := GetRandomNum(3)

	// 不能用除来取，会取到单个数值的情况，像 0，这样订单就不满足长度条件了。
	// centisecond := now.Nanosecond() / 10000000
	// centisecStr := strconv.Itoa(centisecond)

	applog.Debugf("gen short bizid: %v", dataStr)
	// applog.Errorf("centisecStr:%v", centisecStr)
	// applog.Errorf("micrrdmStrosecStr:%v", rdmStr)

	str := dataStr + centisecStr + rdmStr

	return str
}

// OrderShortBizIDGenUint64 OrderShortBizIDGenUint64
func OrderShortBizIDGenUint64() uint64 {
	str := OrderShortBizIDGen()
	num := ParseStrToUint64(str, 0)

	return num
}

// OrderOriIDRestore OrderOriIDRestore
// order-2019080812282674502
func OrderOriIDRestore(bizid string) string {
	oriID := bizid[6:]
	return oriID
}

// // UserBizIDGen UserBizIDGen
// func UserBizIDGen() string {
// 	rdxString := GetRandomString(87)
// 	return cryptoutils.MD5HexShort(rdxString)
// }
