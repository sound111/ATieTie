package mysql

import (
	"TieTie/models"
	"TieTie/myError"
	"TieTie/pkg/snowflakes"
	"crypto/md5"
	"encoding/hex"
)

func Register(user *models.User) (err error) {
	sql := "select count(user_id) from user where username=?"

	var count int
	err = db.Get(&count, sql, user.Username)
	//err指数据库查询是否出错

	if count > 0 {
		err = myError.ErrorUserExist
		return
	}

	//生成UID
	userId, err := snowflakes.GetID()
	user.UserId = userId

	//密码加密
	epassword := encrypt(user.Password)
	user.Password = epassword

	sql = "insert into user(user_id,username,password) values(?,?,?)"

	_, err = db.Exec(sql, user.UserId, user.Username, user.Password)

	return
}

func Login(user *models.User) (err error) {
	opassword := user.Password

	sql := "select user_id,username,password from user where username=?"
	err = db.Get(user, sql, user.Username)
	if err != nil {
		err = myError.ErrorUserNotExist
		return
	}

	opassword = encrypt(opassword)

	if opassword != user.Password {
		err = myError.ErrorPwdInvalid
	}

	return
}

func encrypt(password string) (epassword string) {
	var salt = "yaye" //应采用随机盐
	h := md5.New()
	h.Write([]byte(salt))
	epassword = hex.EncodeToString(h.Sum([]byte(password)))

	return
}
