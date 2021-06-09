package model

type HotRespond struct {
	ReportTime string    `json:"report_time"`
	List       []HotItem `json:"lists"`
}

type HotItem struct {
	// 标题
	Title string `json:"title"`
	// 链接
	Url string `json:"url"`
	// 热度
	Popular string `json:"popular"`
	// 趋势
	Trend string `json:"trend,omitempty"`
	// 作者
	Author string `json:"author,omitempty"`
	// 附加
	Extra string `json:"extra"`
}

type ZhihuHot struct {
	InitialState InitialState `json:"initialState"`
}

type InitialState struct {
	Topstory Topstory `json:"topstory"`
}

type Topstory struct {
	HotList []HotList `json:"hotList"`
}

type HotList struct {
	FeedSpecific FeedSpecific `json:"feedSpecific"`
	Target       Target       `json:"target"`
}

type Target struct {
	TitleArea   Area         `json:"titleArea"`
	ExcerptArea Area         `json:"excerptArea"`
	ImageArea   ImageArea    `json:"imageArea"`
	MetricsArea Area         `json:"metricsArea"`
	LabelArea   LabelArea    `json:"labelArea"`
	Link        ImageArea    `json:"link"`
	TextTagArea *TextTagArea `json:"textTagArea,omitempty"`
}
type FeedSpecific struct {
	AnswerCount int64 `json:"answerCount"`
}

type Area struct {
	Text string `json:"text"`
}

type ImageArea struct {
	URL string `json:"url"`
}
type LabelArea struct {
	Type        LabelAreaType `json:"type"`
	Trend       int64         `json:"trend"`
	NightColor  NightColor    `json:"nightColor"`
	NormalColor NormalColor   `json:"normalColor"`
}
type TextTagArea struct {
	Text       string `json:"text"`
	FontColor  string `json:"fontColor"`
	Background string `json:"background"`
}

type LabelAreaType string

const (
	Trend LabelAreaType = "trend"
)

type NightColor string

const (
	B7302D NightColor = "#B7302D"
)

type NormalColor string

const (
	F1403C NormalColor = "#F1403C"
)
