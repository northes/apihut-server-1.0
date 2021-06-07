package model

type Words struct {
	Word string `json:"word" form:"word" binding:"required"`
	Type string `json:"type" form:"type" binding:"required"`
	Hmm  bool   `json:"hmm" form:"hmm"`
}
