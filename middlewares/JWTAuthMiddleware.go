package middlewares

import (
	"TieTie/controller"
	"TieTie/pkg/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			controller.ResponseError(c, controller.CodeNotLogin)

			c.Abort()
			return
		}

		parts := strings.SplitN(token, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controller.ResponseError(c, controller.CodeTokenFormatErr)

			c.Abort()
			return
		}

		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			controller.ResponseError(c, controller.CodeTokenParseErr)

			c.Abort()
			return
		}

		c.Set(controller.CtxUserId, mc.UserId)
		c.Next()
	}
}
