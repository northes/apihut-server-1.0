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

/**************
	天气
**************/
type WeatherCode int

// Weather：晴 阴 云 雨 雾 雪 台风
const (
	WeatherDefault WeatherCode = 0
	WeatherSunny   WeatherCode = 1
	WeatherCloudy  WeatherCode = 2
	WeatherCloud   WeatherCode = 3
	WeatherWind    WeatherCode = 4
	WeatherRain    WeatherCode = 5
	WeatherFog     WeatherCode = 6
	WeatherTyphoon WeatherCode = 7
)

var weatherMap = map[WeatherCode]string{
	WeatherDefault: "默认",
	WeatherSunny:   "晴",
	WeatherCloudy:  "阴",
	WeatherCloud:   "云",
	WeatherWind:    "风",
	WeatherRain:    "雨",
	WeatherFog:     "雪",
	WeatherTyphoon: "台风",
}

// 换取天气文字
func (w WeatherCode) Text() string {
	return weatherMap[w]
}
