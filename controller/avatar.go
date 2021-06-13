package controller

import (
	"apihut-server/config"
	"apihut-server/constant"
	"apihut-server/model"
	"apihut-server/server"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AvatarHandler(c *gin.Context) {
	// 初始化
	p := model.IdenticonAvatar{
		Hash:      c.ClientIP(),
		Namespace: config.Conf.Name,
		Size:      5,
		Density:   3,
		Pixels:    144,
		Output:    constant.ImageOutput,
	}
	err := c.ShouldBindQuery(&p)
	if err != nil {
		fmt.Println(err.Error())
		ResponseError(c, CodeParameterFailure)
		return
	}

	// 获取头像
	img, err := server.GetIdenticonAvatar(&p)
	if err != nil {
		ResponseError(c, CodeServerRequestFailure)
		return
	}

	// 分类输出
	if p.Output == constant.ImageOutput {
		c.File(img)
		return
	} else if p.Output == constant.Base64Output {
		c.String(http.StatusOK, img)
		return
	}

	ResponseError(c, CodeServerBusy)
}
