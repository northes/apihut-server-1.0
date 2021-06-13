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
	r.StaticFile("favicon.ico", "./static/favicon.ico")

	r.GET("/greet", controller.GreetHandler)
	r.GET("/weather", controller.WeatherHandler)
	r.GET("/ip", controller.IPHandler)
	r.GET("/words", controller.WordsHandler)
	r.GET("/rank/:site", controller.RankHandler)
	bilibili := r.Group("/bilibili")
	{
		bilibili.GET("/av2bv", controller.BilibiliAv2Bv)
	}
	r.GET("/avatar", controller.AvatarHandler)

	r.NoRoute(controller.NotfoundHandler)
	return r
}
