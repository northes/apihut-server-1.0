package model

type Av2Bv struct {
	Aid int    `json:"aid" form:"av"`
	Bid string `json:"bvid" form:"bv"`
}
