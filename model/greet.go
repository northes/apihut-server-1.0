package model

// Language： cn en
// Time： 凌晨 早晨 中午 下午 傍晚 晚上 午夜 深夜
// Weather：晴 阴 云 雨 雾 雪 台风

type Greet struct {
	ID       int    `xorm:"pk index notnull comment(句子id)"`
	Sentence string `xorm:"notnull comment(句子详情)"`
	Author   string `xorm:"comment(句子作者)"`
	Language int    `xorm:"notnull comment(语言)"`
	Time     int    `xorm:"notnull comment(语境时间)"`
	Weather  int    `xorm:"notnull comment(语境天气)"`
	Day      int    `xorm:"notnull comment(星期)"`
	Tags     string `xorm:"comment(标签)"`
}

type GreetRespond struct {
	Words    string `json:"words"`
	Sentence string `json:"sentence"`
	Author   string `json:"author,omitempty"`
}
