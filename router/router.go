package router

import (
	"apihut-server/controller"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()

	r.GET("/greet", controller.GreetHandler)

	return r
}
