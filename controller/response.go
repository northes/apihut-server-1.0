package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseCode int

const (
	CodeSuccess          ResponseCode = 200
	CodeInvalidParameter ResponseCode = 0
	CodeServerBusy       ResponseCode = 1
)

var responseCodeMap = map[ResponseCode]string{
	CodeSuccess:          "成功",
	CodeInvalidParameter: "参数错误",
	CodeServerBusy:       "服务繁忙",
}

type response struct {
	Code ResponseCode `json:"code"`
	Msg  string       `json:"msg"`
	Data interface{}  `json:"data,omitempty"`
}

func (c ResponseCode) Msg() string {
	return responseCodeMap[c]
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &response{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	})
}

func ResponseSuccessWithMsg(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, &response{
		Code: CodeSuccess,
		Msg:  msg,
		Data: data,
	})
}

func ResponseError(c *gin.Context, code ResponseCode) {
	c.JSON(http.StatusOK, &response{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}
