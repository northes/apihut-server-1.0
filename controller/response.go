package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &response{
		Code: 200,
		Msg:  "success",
		Data: data,
	})
}

func ResponseSuccessWithMsg(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, &response{
		Code: 200,
		Msg:  msg,
		Data: data,
	})
}

func ResponseError(c *gin.Context) {
	c.JSON(http.StatusOK, &response{
		Code: 300,
		Msg:  "错误",
		Data: nil,
	})
}
