package constant

type SeasonCode int

const (
	SeasonDefault SeasonCode = 0
	SeasonSpring  SeasonCode = 1
	SeasonSummer  SeasonCode = 2
	SeasonAutumn  SeasonCode = 3
	SeasonWinter  SeasonCode = 4
)

var seasonMap = map[SeasonCode]string{
	SeasonDefault: "通用",
	SeasonSpring:  "春天",
	SeasonSummer:  "夏天",
	SeasonAutumn:  "秋天",
	SeasonWinter:  "冬天",
}

func (s SeasonCode) Text() string {
	return seasonMap[s]
}
