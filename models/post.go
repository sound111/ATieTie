package models

// Post 内存对齐 将类型相同的字段放到一起
type Post struct {
	PostId      int64  `json:"post_id" db:"post_id"`
	AuthorId    int64  `json:"author_id" db:"author_id"`
	CommunityId int64  `json:"community_id" db:"community_id"`
	Status      int32  `json:"status" db:"status"`
	Title       string `json:"title" db:"title"`
	Content     string `json:"content" db:"content"`
}
