package utils

import (
	"math"
	"strconv"
	"time"

	applog "github.com/lanvige/goutils/logger"
)

// Utc8Lastday 当前服务器时间的东八区昨天日期
// 服务器时区不重要，统一换算为东八区
// 再向前取一天
func Utc8Lastday() time.Time {
	var utc8, _ = time.LoadLocation("Asia/Shanghai") //上海
	timeUtc8 := time.Now().UTC().In(utc8)

	// 对日期减去一天
	yesterday := timeUtc8.AddDate(0, 0, -1)

	return yesterday
}

// Utc8LastdayString 当前服务器时间的东八区昨天日期
// 服务器时区不重要，统一换算为东八区
// 再向前取一天
func Utc8LastdayString() string {
	yesterday := Utc8Lastday()
	// 格式化为只有日期
	yesString := yesterday.Format("2006-01-02")

	return yesString
}

// Utc8Day 当前服务器时间的东八区指定天数日期
// 服务器时区不重要，统一换算为东八区
// 再向前取一天
func Utc8Day(days int) time.Time {
	var utc8, _ = time.LoadLocation("Asia/Shanghai") //上海
	timeUtc8 := time.Now().UTC().In(utc8)

	// 对日期减去一天
	day := timeUtc8.AddDate(0, 0, days)

	return day
}

// Utc8DayString 当前服务器时间的东八区昨天日期
// 服务器时区不重要，统一换算为东八区
// 再向前取一天
func Utc8DayString(days int) string {
	day := Utc8Day(days)
	// 格式化为只有日期
	dayString := day.Format("2006-01-02")

	return dayString
}

// Utc8Today 东八区今天日期
func Utc8Today() time.Time {
	var utc8, _ = time.LoadLocation("Asia/Shanghai") //上海
	timeUtc8 := time.Now().UTC().In(utc8)

	return timeUtc8
}

// Utc8TodayString 东八区今天日期
func Utc8TodayString() string {
	timeUtc8 := Utc8Today()

	// 格式化为只有日期
	todString := timeUtc8.Format("2006-01-02")

	return todString
}

// Utc8DuctionDays 一个日期和另一个日期相差多少天
// 要都在东八区算
// 两边都要先算成 0点，因为两边都算整天的
func Utc8DuctionDays(time1, time2 time.Time) int {
	// 先换算成 东八区 0点
	time10Utc8 := Utc8Day0(time1)
	time20Utc8 := Utc8Day0(time2)

	applog.Info(time10Utc8.Format("2006-01-02T15:04:05-0700"))
	applog.Info(time20Utc8.Format("2006-01-02T15:04:05-0700"))
	// 还余多少天
	// Sub 计算两个时间差
	subM := time20Utc8.Sub(time10Utc8)
	remaining := int(math.Floor(subM.Hours() / 24))
	applog.Errorf("remaining: %v", remaining)

	return remaining
}

// Utc8Day0 把任一时间，换算成东八区0点
func Utc8Day0(time1 time.Time) time.Time {
	var utc8, _ = time.LoadLocation("Asia/Shanghai") //上海
	timeUtc8 := time.Date(time1.Year(), time1.Month(), time1.Day(), 0, 0, 0, 0, utc8)

	return timeUtc8
}

// Utc8Current 取当前时间的 UTC8
func Utc8Current() time.Time {
	// 取 Utc0 时间
	utc8, _ := time.LoadLocation("Asia/Shanghai")
	timeUtc8 := time.Now().In(utc8)
	// 按 ISO8601 输出
	applog.Info(timeUtc8.Format("2006-01-02T15:04:05-0700"))

	return timeUtc8
}

// Utc8Time 取时间的 UTC8
func Utc8Time(tm time.Time) time.Time {
	// 取 Utc0 时间
	utc8, _ := time.LoadLocation("Asia/Shanghai")
	timeUtc0 := tm.In(utc8)
	// 按 ISO8601 输出
	applog.Info(timeUtc0.Format("2006-01-02T15:04:05-0700"))

	return timeUtc0
}

// Utc0Current 取当前时间的 UTC0
func Utc0Current() time.Time {
	// 取 Utc0 时间
	utc0, _ := time.LoadLocation("UTC")
	timeUtc0 := time.Now().In(utc0)
	timeUtc02 := time.Now().UTC()
	// 按 ISO8601 输出
	applog.Info(timeUtc0.Format("2006-01-02T15:04:05-0700"))
	applog.Info(timeUtc02.Format("2006-01-02T15:04:05-0700"))

	return timeUtc02
}

// Utc0Time 取时间的 UTC0
func Utc0Time(tm time.Time) time.Time {
	// 取 Utc0 时间
	utc0, _ := time.LoadLocation("UTC")
	timeUtc0 := tm.In(utc0)
	timeUtc02 := tm.UTC()
	// 按 ISO8601 输出
	applog.Info(timeUtc0.Format("2006-01-02T15:04:05-0700"))
	applog.Info(timeUtc02.Format("2006-01-02T15:04:05-0700"))

	return timeUtc02
}

// LocalCurrent LocalCurrent
func LocalCurrent() time.Time {
	// 取服务器本地时间，默认就是
	utcLocal, _ := time.LoadLocation("Local")
	timeLocal := time.Now().In(utcLocal)
	// 按 ISO8601 输出
	applog.Info(timeLocal.Format("2006-01-02T15:04:05-0700"))

	return timeLocal
}

// Time0 取任一时间对应的 0 点
func Time0(time1 time.Time) time.Time {
	tm1 := time.Date(time1.Year(), time1.Month(), time1.Day(), 0, 0, 0, 0, time1.Location())
	return tm1
}

// Common Common
func Common() {
	// 取服务器本地时间，默认就是
	utcLocal, _ := time.LoadLocation("Local")
	timeUtcLocal := time.Now().In(utcLocal)
	// 按 ISO8601 输出
	applog.Info(timeUtcLocal.Format("2006-01-02T15:04:05-0700"))

	// 取 Utc0 时间
	utc0, _ := time.LoadLocation("UTC")
	timeUtc0 := time.Now().In(utc0)
	timeUtc02 := time.Now().UTC()
	// 按 ISO8601 输出
	applog.Info(timeUtc0.Format("2006-01-02T15:04:05-0700"))
	applog.Info(timeUtc02.Format("2006-01-02T15:04:05-0700"))

	// =====
	// timeA := time.Now()
	// 如何知道这个 time 的时区？

	// currentDay := time.Date(timeUtc8.Year(), timeUtc8.Month(), timeUtc8.Day(), 0, 0, 0, 0, timeUtc8.Location())

	// =====
	// // // 默认UTC
	// // loc, err := time.LoadLocation("")
	// // // 一般为CST
	// // loc, err := time.LoadLocation("Local")
	// // // 美国洛杉矶PDT
	// // loc, err := time.LoadLocation("America/Los_Angeles")
	// // // CST
	// // loc, _ := time.LoadLocation("Asia/Chongqing")

	// // time.FixedZone("CST", 8*3600) // 东八区时间
	// // currentDay := time.Now()

	// =====
	// applog.Info(currentDay.Format("2006-01-02"))

	// applog.Info("start to handle deposit crew!")

	// return currentDay
}

// ISO 8601
// Format("2006-01-02T15:04:05-0700")

// 只要日期
// Format("2006-01-02"))
// Format("20060102"))

// 以指定时间开始，设计一个全新的 c_timestamp

// GenCustomTimestamp GenCustomTimestamp
func GenCustomTimestamp() string {
	now := time.Now()

	dataStr := now.Format("200601021504")
	nanosec := now.Nanosecond()
	nanosecStr := strconv.Itoa(nanosec)
	microsecStr := nanosecStr[0:3]
	rdmStr := GenRandomNumString(4)

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
		// applog.Error(err)
		return defaultValue
	}

	return t
}
