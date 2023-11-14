package models

// User 前后端数据传输失真问题   ,string
type User struct {
	UserId   uint64 `db:"user_id,string"`
	Username string `db:"username"`
	Password string `db:"password"`
}
