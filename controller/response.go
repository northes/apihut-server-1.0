package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

func ResponseSuccessWithMsg(c *gin.Context,msg string,data interface{})  {
	c.JSON(http.StatusOK,&response{
		Code: 200,
		Msg:  msg,
		Data: data,
	})
}