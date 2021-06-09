package controller

import (
	"apihut-server/model"
	"apihut-server/server"
	"time"

	"github.com/gin-gonic/gin"
)

func HotHandler(c *gin.Context) {

	site := c.Param("site")

	hotList, err := server.GetHot(site)
	if err != nil {
		ResponseError(c, CodeServerRequestFailure)
		return
	}

	ResponseSuccess(c, model.HotRespond{
		ReportTime: time.Now().Format("2006-01-02 15:04:05"),
		List:       hotList,
	})
}
