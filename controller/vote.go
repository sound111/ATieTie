package controller

import (
	"TieTie/logic"
	"TieTie/models"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func PostVote(c *gin.Context) {
	//参数
	p := new(models.ParamVoteData)

	err := c.ShouldBindJSON(p)
	if err != nil {
		_, ok := err.(validator.ValidationErrors) //类型断言
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMessage(c, CodeInvalidParam, err.Error())
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}

	err = logic.PostVote(int64(userId), p)
	if err != nil {
		zap.L().Error(err.Error(), zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, nil)
	return
}
