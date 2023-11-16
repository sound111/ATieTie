package logic

import (
	"TieTie/dao/myRedis"
	"TieTie/dao/mysql"
	"TieTie/models"
	"fmt"

	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	err = mysql.CreatePost(p)
	if err != nil {
		return
	}

	err = myRedis.CreatePost(p)
	return
}

func GetPostInfo(id int64) (p *models.ParamPostInfo, err error) {
	return mysql.GetPostInfo(id)
}

func GetPostList() (params []*models.ParamPostInfo, err error) {
	posts, err := mysql.GetPostList()

	for _, post := range posts {
		user, err := mysql.GetUserNameById(int64(post.AuthorId))
		if err != nil {
			zap.L().Error("mysql.GetUserNameById(int64(post.AuthorId)) fails",
				zap.Int64("author_id", int64(post.AuthorId)),
				zap.Error(err))
		}

		community, err := mysql.GetCommunityInfo(post.CommunityId)
		if err != nil {
			zap.L().Error("mysql.GetCommunityInfo(post.CommunityId) fails",
				zap.Int64("community_id", int64(post.CommunityId)),
				zap.Error(err))
		}

		param := new(models.ParamPostInfo)

		param.Post = post
		param.Community = community
		param.AuthorName = user

		params = append(params, param)
	}

	return
}

func GetPostList2(p *models.ParamPostList) (params []*models.ParamPostInfo, err error) {
	postIds, err := myRedis.GetPostsInOrder(p)
	if len(postIds) == 0 {
		zap.L().Warn("myRedis.GetPostsInOrder(p) returns no data")
		return
	}

	fmt.Printf("%s", postIds[0])
	fmt.Println("after redis")
	posts, err := mysql.GetPostInfoByList(postIds)
	fmt.Printf("posts:%#v", posts)

	fmt.Println("after mysql")

	data, err := myRedis.GetPostVoteData(postIds)
	if err != nil {
		return nil, err
	}

	for i, post := range posts {
		user, err := mysql.GetUserNameById(int64(post.AuthorId))
		if err != nil {
			zap.L().Error("mysql.GetUserNameById(int64(post.AuthorId)) fails",
				zap.Int64("author_id", int64(post.AuthorId)),
				zap.Error(err))
		}

		fmt.Println(user)
		community, err := mysql.GetCommunityInfo(post.CommunityId)
		if err != nil {
			zap.L().Error("mysql.GetCommunityInfo(post.CommunityId) fails",
				zap.Int64("community_id", int64(post.CommunityId)),
				zap.Error(err))
		}

		param := new(models.ParamPostInfo)

		param.Post = post
		param.Community = community
		param.AuthorName = user
		param.VoteNum = data[i]

		params = append(params, param)
	}

	return
}
