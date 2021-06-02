package controller

import (
	"apihut-server/model"
	"apihut-server/server"
	"fmt"

	"github.com/gin-gonic/gin"
)

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

	ResponseSuccess(c, g)
}
