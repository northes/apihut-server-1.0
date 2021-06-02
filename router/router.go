package router

import (
	"apihut-server/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Home")
	})
	r.GET("/greet", controller.GreetHandler)

	return r
}
