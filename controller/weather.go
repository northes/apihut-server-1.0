package controller

import (
	"apihut-server/model"
	"apihut-server/server"
	"apihut-server/util"

	"github.com/gin-gonic/gin"
)

// WeatherHandler 天气控制器
func WeatherHandler(c *gin.Context) {
	var p model.Weather
	err := c.ShouldBindQuery(&p)
	if err != nil {
		ResponseError(c, CodeParameterFailure)
		return
	}

	weather, err := server.GetWeather(&p)
	if err != nil {
		if err == util.ErrAPIRequest {
			ResponseError(c, CodeServerRequestFailure)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, weather)
}
