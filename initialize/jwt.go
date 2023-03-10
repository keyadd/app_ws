package initialize

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const TokenExpireDuration = time.Hour * 2 //过期时间

var myKey = []byte("adb1234556") //加盐值

type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenToken(userID int64, username string) (string, error) {
	c := MyClaims{
		UserID:   userID,
		Username: username, //自定义字段
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), //过期时间
			Issuer:    "my-projet",                                //签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	signingString, err := token.SignedString(myKey)
	return signingString, err
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return myKey, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid { // 校验token
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
