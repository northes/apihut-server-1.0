package model

type Words struct {
	Text string `json:"text" form:"text" binding:"required"`
	Type string `json:"type" form:"type" binding:"required"`
	Hmm  bool   `json:"hmm" form:"hmm"`
}

//TODO 返回结果结构体（开始位置，结束位置，词性）
