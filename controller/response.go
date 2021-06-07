package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseCode int

const (
	CodeSuccess              ResponseCode = 200
	CodeParameterFailure     ResponseCode = 300
	CodeServerRequestFailure ResponseCode = 520
	CodeServerBusy           ResponseCode = 530
)

var responseCodeMap = map[ResponseCode]string{
	CodeSuccess:              "成功",
	CodeParameterFailure:     "参数错误",
	CodeServerRequestFailure: "后端请求失败",
	CodeServerBusy:           "服务繁忙",
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

func ResponseErrorWithMsg(c *gin.Context, code ResponseCode, msg string) {
	c.JSON(http.StatusOK, &response{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}
