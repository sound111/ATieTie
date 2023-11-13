package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
)

const CtxUserId = "UserId"

var ErrUserNotLogin = errors.New("user not login")

func GetUserId(c *gin.Context) (uint64, error) {
	id, ok := c.Get(CtxUserId)
	if !ok {
		return 0, ErrUserNotLogin
	}

	uid := id.(uint64)

	return uid, nil
}
