package model

import (
	"apihut-server/constant"
	"time"
)

type Rank struct {
	ID          int64             `xorm:"pk autoincr notnull comment(id)" json:"-" `
	ReportTime  string            `xorm:"-" json:"report_time"`
	SiteName    constant.SiteName `xorm:"notnull VARCHAR(255) comment(网站名称)" json:"-"`
	List        []RankItem        `xorm:"notnull json comment(热榜)" json:"lists"`
	CreatedTime time.Time         `xorm:"created comment(创建时间)" json:"-"`
}

type RankItem struct {
	Title   string `json:"title"`            // 标题
	Url     string `json:"url"`              // 链接
	Popular string `json:"popular"`          // 热度
	Trend   string `json:"trend,omitempty"`  // 趋势
	Author  string `json:"author,omitempty"` // 作者
	Extra   string `json:"extra"`            // 附加
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
