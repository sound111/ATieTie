package controller

import (
	"TieTie/myError"

	"github.com/gin-gonic/gin"
)

const (
	CtxUserId = "UserId"
)

func getUserId(c *gin.Context) (uint64, error) {
	id, ok := c.Get(CtxUserId)
	if !ok {
		return 0, myError.ErrUserNotLogin
	}

	uid := id.(uint64)

	return uid, nil
}
