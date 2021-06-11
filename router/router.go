package router

import (
	"apihut-server/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default())

	r.GET("/", controller.HomeHandler)

	r.GET("/greet", controller.GreetHandler)
	r.GET("/weather", controller.WeatherHandler)
	r.GET("/ip", controller.IPHandler)
	r.GET("/words", controller.WordsHandler)
	r.GET("/rank/:site", controller.RankHandler)

	r.NoRoute(controller.NotfoundHandler)
	return r
}
