package constant

type TimeCode int

// Time：通用 凌晨 早晨 中午 下午 傍晚 晚上  深夜
const (
	TimeDefault      TimeCode = 0
	TimeEarlyMorning TimeCode = 1
	TimeMorning      TimeCode = 2
	TimeNoon         TimeCode = 3
	TimeAfternoon    TimeCode = 4
	TimeEvening      TimeCode = 5
	TimeNight        TimeCode = 6
	TimeLateNight    TimeCode = 7
)

var timeMap = map[TimeCode]string{
	TimeDefault:      "默认",
	TimeEarlyMorning: "凌晨",
	TimeMorning:      "早上",
	TimeNoon:         "中午",
	TimeAfternoon:    "下午",
	TimeEvening:      "傍晚",
	TimeNight:        "晚上",
	TimeLateNight:    "深夜",
}

// 换取时间段文字
func (c TimeCode) Text() string {
	return timeMap[c]
}
