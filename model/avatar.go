package model

import (
	"apihut-server/constant"
)

type IdenticonAvatar struct {
	Hash            string              `form:"hash" binding:"max=32"`
	Namespace       string              `form:"namespace" binding:"max=32"`
	Size            int                 `form:"size" binding:"gte=4,lte=32"`
	Density         int                 `form:"density" binding:"gte=1,lte=64"`
	Pixels          int                 `form:"pixels" binding:"gte=1,lte=300"`
	Output          constant.OutputCode `form:"output"`
	Random          bool                `form:"random"`
	BackgroundColor string              `form:"bgcolor"`
	FillColor       string              `form:"fcolor"`
}
