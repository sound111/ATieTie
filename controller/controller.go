package controller

import (
	"Web/settings"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, settings.Conf.AppConfig.Name)
}
