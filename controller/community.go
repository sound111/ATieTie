package controller

import (
	"TieTie/logic"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetCommunityList(c *gin.Context) {
	communities, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error(err.Error(), zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, communities)
}

func GetCommunityInfo(c *gin.Context) {
	sid, ok := c.Params.Get("id")
	if !ok {
		ResponseError(c, CodeNoID)
		return
	}

	//base 进制
	id, err := strconv.ParseInt(sid, 10, 64)
	if err != nil {
		ResponseError(c, CodeRequestParamsErr)
		return
	}

	data, err := logic.GetCommunityInfo(id)
	if err != nil {
		zap.L().Error(err.Error(), zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)
}
