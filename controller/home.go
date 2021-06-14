package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {

	const baseUrl = "https://apihut.net"

	type Home struct {
		DocsUrl    string `json:"docs_url"`
		IPUrl      string `json:"ip_url"`
		WeatherUrl string `json:"weather_url"`
		GreetUrl   string `json:"greet_url"`
		WordsUrl   string `json:"words_url"`
		RankUrl    string `json:"rank_url"`
		Av2BV      string `json:"av2bv_url"`
		Avatar     string `json:"avatar"`
	}

	//c.JSON(http.StatusOK, Home{
	//	DocsUrl:    "https://docs.apihut.net/",
	//	IPUrl:      baseUrl + "/ip",
	//	WeatherUrl: baseUrl + "/weather",
	//	GreetUrl:   baseUrl + "/greet",
	//	WordsUrl:   baseUrl + "/words",
	//	RankUrl:    baseUrl + "/rank/:site_name",
	//	Av2BV:      baseUrl + "/bilibili/av2bv",
	//	Avatar:     baseUrl + "/avatar",
	//})

	c.HTML(http.StatusOK, "index.html", nil)

}

func NotfoundHandler(c *gin.Context) {

	type Notfound struct {
		Message          string `json:"message"`
		DocumentationUrl string `json:"documentation_url"`
	}

	c.JSON(http.StatusOK, Notfound{
		Message:          "Not Found",
		DocumentationUrl: "https://docs.apihut.net",
	})

}
