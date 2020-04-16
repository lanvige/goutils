package utils

import (
	"strconv"
	"time"
)

// 以指定时间开始，设计一个全新的 c_timestamp

// CustomTimestampGen CustomTimestampGen
func CustomTimestampGen() string {
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

// ParseTimeNano ParseTimeNano
func ParseTimeNano(value string, defaultValue time.Time) time.Time {
	// layout := "2019-07-19T02:53:32.441909288Z"
	t, err := time.Parse(time.RFC3339Nano, value)

	if err != nil {
		// logrus.Error(err)
		return defaultValue
	}

	return t
}
