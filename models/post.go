package models

import "time"

// Post 内存对齐 将类型相同的字段放到一起
type Post struct {
	PostId      uint64    `json:"post_id,string" db:"post_id"`
	AuthorId    uint64    `json:"author_id,string" db:"author_id"`
	CommunityId int64     `json:"community_id" db:"community_id"`
	Status      int32     `json:"status" db:"status"`
	Title       string    `json:"title" db:"title"`
	Content     string    `json:"content" db:"content"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}
