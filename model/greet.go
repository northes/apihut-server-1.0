package model

import "apihut-server/constant"

type Greet struct {
	ID       int                   `xorm:"pk index notnull comment(句子id)" json:"-" `
	Sentence string                `xorm:"notnull comment(句子详情)" json:"sentence"`
	Author   string                `xorm:"comment(句子作者)" json:"author"`
	Language constant.LanguageCode `xorm:"notnull comment(语言)" json:"language" form:"language"`
	Time     constant.TimeCode     `xorm:"notnull comment(语境时间)" json:"time" form:"time" binding:"gte=-1,lte=7"`
	Weather  constant.WeatherCode  `xorm:"notnull comment(语境天气)" json:"weather" form:"weather"`
	Day      constant.DayCode      `xorm:"notnull comment(星期)" json:"day" form:"day" binding:"gte=-1,lte=7"`
	Season   constant.SeasonCode   `xorm:"notnull comment(季节)" json:"season" form:"season" binding:"gte=-1,lte=4"`
	NoRange  bool                  `xorm:"-" json:"no_range" form:"no_range"`
	Tags     string                `xorm:"comment(标签)" json:"-"`
}

type GreetRespond struct {
	Words    string `json:"word"`
	Sentence string `json:"sentence,omitempty"`
	Author   string `json:"author,omitempty"`
}
