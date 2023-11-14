package mysql

import (
	"TieTie/models"
	"TieTie/myError"
	"TieTie/pkg/snowflakes"
	"database/sql"
	"errors"

	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	PostId, err := snowflakes.GetID()
	p.PostId = PostId

	sqlStr := "insert into post(post_id,title,content,author_id,community_id,status) values(?,?,?,?,?,?)"

	_, err = db.Exec(sqlStr, p.PostId, p.Title, p.Content, p.AuthorId, p.CommunityId, p.Status)

	return
}

func GetPostInfo(id int64) (p *models.ParamPostInfo, err error) {
	p = new(models.ParamPostInfo)
	post := new(models.Post)
	sqlStr := "select post_id,title,content,author_id,community_id,status from post where post_id=?"

	err = db.Get(post, sqlStr, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			zap.L().Warn("no community data in db")
			err = nil
		} else {
			return
		}
	}

	s := new(string)
	sqlStr = "select username from user where user_id=?"
	err = db.Get(s, sqlStr, post.AuthorId)
	if err != nil {
		err = myError.ErrorUserNotExist
		return
	}

	community, err := GetCommunityInfo(post.CommunityId)

	p.Community = community
	p.AuthorName = *s
	p.Post = post
	return
}

// GetPostList 分页查询 传两个路径参数 offset和rows就可以实现
// SELECT * FROM tableName LIMIT [offset,] rows | rows OFFSET offset
//
// # tableName：表名
// # offset：可选项，偏移量，指定了结果集的起始位置(从0开始)，为0时可省略
// # rows：行数，指定了返回结果集的行数
func GetPostList() (p []*models.Post, err error) {
	sqlStr := "select post_id,title from post"

	err = db.Select(&p, sqlStr)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			zap.L().Warn("no community data in db")
			err = nil
		} else {
			return
		}
	}

	return
}
