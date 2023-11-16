package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response
type Response struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, &Response{
		code,
		code.Msg(),
		nil,
	})
}

func ResponseErrorWithMessage(c *gin.Context, code ResCode, msg string) {
	c.JSON(http.StatusOK, &Response{
		code,
		msg,
		nil,
	})
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		CodeSuccess,
		CodeSuccess.Msg(),
		data,
	})
}
