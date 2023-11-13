package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const TokenExpireDuration = time.Hour * 2

var Salt = []byte("嘿哈")

type MyClaims struct {
	UserId uint64
	jwt.RegisteredClaims
}

func GetToken(userid uint64) (string, error) {
	c := MyClaims{
		userid,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)), //过期时间
			Issuer:    "TieTie",                                                //签发人
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return token.SignedString(Salt)
}

func ParseToken(tokenStr string) (*MyClaims, error) {
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenStr, mc, func(token *jwt.Token) (interface{}, error) {
		return Salt, nil
	})

	if err != nil {
		return nil, err
	}

	if token.Valid {
		return mc, nil
	}

	return nil, errors.New("invalid token")
}
