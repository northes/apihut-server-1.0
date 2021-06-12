package controller

import (
	"apihut-server/model"
	"apihut-server/server"

	"github.com/gin-gonic/gin"
)

func BilibiliAv2Bv(c *gin.Context) {
	var id model.Av2Bv
	err := c.ShouldBindQuery(&id)
	if err != nil {
		ResponseError(c, CodeParameterFailure)
		return
	}
	// 参数验证
	if id.Aid == 0 && len(id.Bid) == 0 {
		ResponseErrorWithMsg(c, CodeParameterFailure, "bv与av必须传入一个")
		return
	}
	if id.Aid != 0 && len(id.Bid) != 0 {
		ResponseErrorWithMsg(c, CodeParameterFailure, "av与bv只能传入一个")
		return
	}

	ids, err := server.Av2Bv(&id)
	if err != nil {
		ResponseErrorWithMsg(c, CodeParameterFailure, err.Error())
		return
	}

	ResponseSuccess(c, ids)
}

func BilibiliVideoHandler(c *gin.Context) {

}

func BilibiliUserHandler(c *gin.Context) {

}
