package controller

import (
	"apihut-server/model"
	"apihut-server/server"

	"github.com/gin-gonic/gin"
)

func IPHandler(c *gin.Context) {
	var p model.IP
	err := c.ShouldBindQuery(&p)
	if err != nil {
		ResponseError(c, CodeParameterFailure)
		return
	}

	if len(p.IP) == 0 {
		p.IP = c.ClientIP()
	}

	ipInfo, err := server.GetIPInfo(&p)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, ipInfo)
}
