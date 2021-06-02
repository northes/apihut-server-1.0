package controller

import (
	"apihut-server/server"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GreetHandler(c *gin.Context) {
	g, err := server.Greet()
	if err != nil {
		fmt.Println(err.Error())
		ResponseError(c)
		return
	}
	ResponseSuccess(c, g)
}
