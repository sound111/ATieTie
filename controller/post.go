package controller

import (
	"TieTie/logic"
	"TieTie/models"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreatePost(c *gin.Context) {
	//获取参数以及参数校验
	var p *models.Post
	err := c.ShouldBindJSON(&p)
	if err != nil {
		zap.L().Error("Create with invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	//取到当前用户id
	p.AuthorId, err = getUserId(c)

	//业务处理
	err = logic.CreatePost(p)
	if err != nil {
		zap.L().Error(err.Error(), zap.Error(err))
		ResponseErrorWithMessage(c, CodeServerBusy, "register fails")

		return
	}

	//返回响应
	ResponseSuccess(c, "create post success")
}

func GetPostInfo(c *gin.Context) {
	sid, ok := c.Params.Get("id")
	if !ok {
		ResponseError(c, CodeNoID)
		return
	}

	//base 进制
	id, err := strconv.ParseInt(sid, 10, 64)
	if err != nil {
		fmt.Println(err)
		ResponseError(c, CodeRequestParamsErr)
		return
	}

	data, err := logic.GetPostInfo(id)
	if err != nil {
		zap.L().Error(err.Error(), zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)
}

func GetPostList(c *gin.Context) {
	posts, err := logic.GetPostList()
	if err != nil {
		zap.L().Error(err.Error(), zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, posts)
}

// GetPostList2 实现分页展示
// 按时间或者分数排序
func GetPostList2(c *gin.Context) {

	p := models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderTime,
	}

	err := c.ShouldBindQuery(&p)
	if err != nil {
		zap.L().Error(err.Error(), zap.Error(err))
		ResponseError(c, CodeRequestParamsErr)
		return
	}
	fmt.Println(p)

	posts, err := logic.GetPostList2(&p)
	if err != nil {
		zap.L().Error(err.Error(), zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, posts)
}
