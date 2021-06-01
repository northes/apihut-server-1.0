package controller

import (
	"apihut-server/server"

	"github.com/gin-gonic/gin"
)

func GreetHandler(c *gin.Context) {
	g, err := server.Greet()
	if err != nil {
		ResponseError(c)
	}
	ResponseSuccess(c, g)
}
