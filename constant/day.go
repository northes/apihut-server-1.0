package constant

// Day：星期一 星期二 星期四 星期五 星期六 星期天

type DayCode int

const (
	DaySunday    DayCode = 0
	DayMonday    DayCode = 1
	DayTuesday   DayCode = 2
	DayWednesday DayCode = 3
	DayThursday  DayCode = 4
	DayFriday    DayCode = 5
	DaySaturday  DayCode = 6
)

var dayMap = map[DayCode]string{
	DaySunday:    "星期天",
	DayMonday:    "星期一",
	DayTuesday:   "星期二",
	DayWednesday: "星期三",
	DayThursday:  "星期四",
	DayFriday:    "星期五",
	DaySaturday:  "星期六",
}

// 换取星期几文字
func (d DayCode) Text() string {
	return dayMap[d]
}
