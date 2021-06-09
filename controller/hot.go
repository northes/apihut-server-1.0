package controller

import (
	"apihut-server/model"
	"apihut-server/server"
	"time"

	"github.com/gin-gonic/gin"
)

func HotHandler(c *gin.Context) {
	hotList, err := server.GetHot()
	if err != nil {
		ResponseError(c, CodeServerRequestFailure)
		return
	}

	ResponseSuccess(c, model.HotRespond{
		ReportTime: time.Now().Format("2006-01-02 15:04:05"),
		List:       hotList,
	})
}
