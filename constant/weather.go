package constant

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
