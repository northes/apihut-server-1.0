package controller

import (
	"apihut-server/server"

	"github.com/gin-gonic/gin"
)

// IPHandler IP控制器
func IPHandler(c *gin.Context) {
	var ip string
	ip, has := c.GetQuery("ip")
	if len(ip) == 0 || !has {
		ip = c.ClientIP()
	}
	if len(ip) == 0 {
		ResponseError(c, CodeParameterFailure)
		return
	}

	ipInfo, err := server.GetIPInfo(ip)
	if err != nil {
		if err == server.ErrIPFormat {
			ResponseErrorWithMsg(c, CodeParameterFailure, server.ErrIPFormat.Error())
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, ipInfo)
}
