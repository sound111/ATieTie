package routes

import (
	"Web/controller"
	"Web/logger"

	"github.com/gin-gonic/gin"
)

func Setup() (r *gin.Engine) {
	r = gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/", controller.Test)

	return
}
