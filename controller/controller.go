package controller

import (
	"TieTie/logic"
	"TieTie/models"
	"TieTie/myError"
	"TieTie/settings"
	"errors"
	"net/http"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, settings.Conf.AppConfig.Name)
}

func Register(c *gin.Context) {
	//获取参数以及参数校验
	var p *models.ParamRegister
	err := c.ShouldBindJSON(&p)
	if err != nil {
		zap.L().Error("Register with invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	//业务处理
	err = logic.Register(p.Username, p.Password)
	if err != nil {
		zap.L().Error(err.Error(), zap.Error(err))
		if errors.Is(err, myError.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
		} else {
			ResponseErrorWithMessage(c, CodeServerBusy, "register fails")
		}

		return
	}

	//返回响应
	ResponseSuccess(c, "register success")
}

func Login(c *gin.Context) {
	//获取参数以及参数校验
	var p *models.ParamLogin
	err := c.ShouldBindJSON(&p)
	if err != nil {
		zap.L().Error("Register with invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	//业务处理
	token, err := logic.Login(p.Username, p.Password)

	if err != nil {
		zap.L().Error(err.Error(), zap.Error(err))
		if errors.Is(err, myError.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
		} else if errors.Is(err, myError.ErrorPwdInvalid) {
			ResponseError(c, CodeInvalidPassword)
		} else {
			ResponseErrorWithMessage(c, CodeServerBusy, "login fails")
		}

		return
	}

	//返回响应
	ResponseSuccess(c, token)
}
