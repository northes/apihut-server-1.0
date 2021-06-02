package constant

type LanguageCode int

// Language： cn en
const (
	LanguageChina   LanguageCode = 0
	LanguageEnglish LanguageCode = 1
)

var languageMap = map[LanguageCode]string{
	LanguageChina:   "中文",
	LanguageEnglish: "英文",
}

func (l LanguageCode) Text() string {
	return languageMap[l]
}
