package routes

import (
	"TieTie/controller"
	"TieTie/logger"
	"TieTie/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup() (r *gin.Engine) {
	r = gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, "success")
	})

	r.POST("/register", controller.Register)

	r.POST("/login", controller.Login)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	return
}
