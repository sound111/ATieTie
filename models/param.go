package models

type ParamRegister struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ParamPostInfo struct {
	AuthorName string `json:"author_name" db:"author_name"`
	*Community `json:"community" db:"community"`
	*Post      //加不加tag，有区别
}

type ParamVoteData struct {
	//UserId从请求中获取
	PostId    int64 `json:"post_id,string" binding:"required"`
	Direction int   `json:"direction" binding:"oneof=1 0 -1"` //0不投票 1赞成票 -1反对票
	//required 当为默认值时，会认为你没填该字段，例如bool为false时，int为0时等
	//validator
}
