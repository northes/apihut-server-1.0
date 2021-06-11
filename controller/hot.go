package controller

import (
	"apihut-server/constant"
	"apihut-server/model"
	"apihut-server/server"

	"github.com/gin-gonic/gin"
)

func HotHandler(c *gin.Context) {
	siteName := c.Param("site")

	// 参数检查
	if siteName != constant.SiteNameBaidu &&
		siteName != constant.SiteNameSina &&
		siteName != constant.SiteNameZhihu &&
		siteName != constant.SiteNameThePaper &&
		siteName != constant.SiteNameBilibili &&
		siteName != constant.SiteNameBilibiliShort {
		ResponseError(c, CodeParameterFailure)
		return
	}
	if siteName == constant.SiteNameBilibiliShort {
		siteName = constant.SiteNameBilibili
	}

	hotList, err := server.GetHot(siteName)
	if err != nil {
		ResponseError(c, CodeServerRequestFailure)
		return
	}

	ResponseSuccess(c, model.Hot{
		ReportTime: hotList.CreatedTime.Format("2006-01-02 15:04:05"),
		HotList:    hotList.HotList,
	})
}
