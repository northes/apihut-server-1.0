package router

import (
	"apihut-server/controller"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default())

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Home")
	})
	r.GET("/greet", controller.GreetHandler)
	r.GET("/weather", controller.WeatherHandler)
	r.GET("/ip", controller.IPHandler)

	return r
}
