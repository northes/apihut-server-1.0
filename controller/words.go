package controller

import (
	"apihut-server/model"
	"apihut-server/server"

	"github.com/gin-gonic/gin"
)

// WordsHandler 分词
func WordsHandler(c *gin.Context) {
	var w model.Words
	err := c.ShouldBindQuery(&w)
	if err != nil {
		ResponseError(c, CodeParameterFailure)
		return
	}

	if w.Type != server.CutTypeBase &&
		w.Type != server.CutTypeSearch &&
		w.Type != server.CutTypeAll {
		ResponseError(c, CodeParameterFailure)
		return
	}

	words, err := server.GetWords(&w)
	if err != nil {
		if err == server.ErrCutType {
			ResponseError(c, CodeParameterFailure)
			return
		}
		ResponseError(c, CodeServerRequestFailure)
		return
	}

	ResponseSuccess(c, words)
}
