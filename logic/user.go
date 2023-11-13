package logic

import (
	"TieTie/dao/mysql"
	"TieTie/models"
	"TieTie/pkg/jwt"
)

func Register(username, password string) (err error) {
	user := models.User{
		Username: username,
		Password: password,
	}

	err = mysql.Register(&user)
	return
}

func Login(username, password string) (token string, err error) {
	user := models.User{
		Username: username,
		Password: password,
	}
	//username password是否正确
	err = mysql.Login(&user)
	if err != nil {
		return
	}

	//生成token
	token, err = jwt.GetToken(user.UserId)

	return
}
