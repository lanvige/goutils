package utils

import (
	"testing"
	"time"

	applog "github.com/lanvige/goutils/logger"
)

func TestGenTimestamp(t *testing.T) {
	code := GenCustomTimestamp()

	// 2020041914465450853
	applog.Error(code)
}

func TestUtc8Lastday(t *testing.T) {
	abc := Utc8Lastday()
	applog.Error(abc)
}

func TestUtc8Duction(t *testing.T) {
	timeA := time.Now()

	timeB := time.Date(2019, 4, 17, 20, 34, 58, 0, time.UTC)
	// timeB := time.Now()
	abc := Utc8DuctionDays(timeA, timeB)
	applog.Error(abc)
}

func TestUtc0Current(t *testing.T) {
	abc := Utc0Current()
	applog.Error(abc)
}

func TestUtc8Current(t *testing.T) {
	abc := Utc8Current()
	applog.Error(abc)
}

func TestLocalCurrent(t *testing.T) {
	abc := LocalCurrent()
	applog.Error(abc)
}

func TestUtc8DayString(t *testing.T) {
	abc := Utc8Day(-30)
	// abcd := Utc8Day0(abc)
	abcde := Utc0Time(abc)
	applog.Error(abcde)
}

func TestABc(t *testing.T) {
	layout := "2006-01-02 15:04:05"
	str := "2019-04-19 17:07:08"
	tm, _ := time.Parse(layout, str)

	abcde := Utc8Time(tm)
	applog.Error(abcde)

}

func TestABc2(t *testing.T) {
	abc, _ := NewDateFromYearMonthDay("2019", "03", "04")

	abccd := TimeFormatDate(abc)
	applog.Error(abccd)

	abccde := TimeFormatDateTime(abc)
	applog.Error(abccde)
}
