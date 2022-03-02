package controller

import (
	"apihut-server/constant"
	"apihut-server/model"
	"apihut-server/server"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GreetHandler 一句问候控制器
func GreetHandler(c *gin.Context) {
	var params model.Greet
	err := c.ShouldBindQuery(&params)
	if err != nil {
		ResponseError(c, CodeParameterFailure)
		return
	}

	g, err := server.Greet(&params)
	if err != nil {
		fmt.Println(err.Error())
		ResponseError(c, CodeServerBusy)
		return
	}

	output, _ := c.GetQuery("output")
	switch output {
	case constant.TextOutput.String():
		break

	}

	if output, _ := c.GetQuery("output"); output == constant.TextOutput.String() {
		c.String(http.StatusOK, g.Sentence)
		return
	}

	ResponseSuccess(c, g)
}
