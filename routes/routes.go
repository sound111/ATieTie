package routes

import (
	"TieTie/controller"
	"TieTie/logger"
	"TieTie/middlewares"
	"net/http"

	swaggerFiles "github.com/swaggo/files"

	_ "TieTie/docs"

	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// gin-swagger middleware
// swagger embed files

func Setup() (r *gin.Engine) {
	r = gin.New()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, "success")
	})

	r.POST("/register", controller.Register)

	r.POST("/login", controller.Login)

	v1 := r.Group("/api/v1")
	v1.Use(middlewares.JWTAuthMiddleware())
	{
		v1.GET("/community", controller.GetCommunityList)
		v1.GET("/community/:id", controller.GetCommunityInfo)

		v1.POST("/post", controller.CreatePost)
		v1.GET("/post/:id", controller.GetPostInfo)
		v1.GET("/post/", controller.GetPostList)
		v1.GET("/post2/", controller.GetPostList2)

		v1.POST("/vote", controller.PostVote)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	return
}
