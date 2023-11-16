package controller

import (
	"TieTie/logic"
	"TieTie/models"
	"TieTie/myError"
	"errors"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Register 用户注册接口
// @Summary 用户注册接口
// @Description 输入用户名和密码进行注册
// @Tags 用户相关接口
// @Accept json
// @Produce json
// @Param object query models.ParamRegister false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} Response
// @Router /register [post]
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

// Login 用户登录接口
// @Summary 用户登录接口
// @Description 用户输入用户名和密码进行登录
// @Tags 用户相关接口
// @Accept json
// @Produce json
// @Param object query models.ParamLogin false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} Response
// @Router /login [get]
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
