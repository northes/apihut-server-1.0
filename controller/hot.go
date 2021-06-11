package controller

import (
	"apihut-server/constant"
	"apihut-server/model"
	"apihut-server/server"

	"github.com/gin-gonic/gin"
)

func RankHandler(c *gin.Context) {
	siteName := constant.SiteName(c.Param("site"))

	// 参数检查
	if !siteName.IsSiteName() {
		ResponseError(c, CodeParameterFailure)
		return
	}

	rank, err := server.GetRank(siteName)
	if err != nil {
		ResponseError(c, CodeServerRequestFailure)
		return
	}

	ResponseSuccess(c, model.Rank{
		ReportTime: rank.CreatedTime.Format("2006-01-02 15:04:05"),
		List:       rank.List,
	})
}
