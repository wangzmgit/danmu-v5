package common

import (
	"time"

	"github.com/spf13/viper"
	"kuukaa.fun/danmu-v5/model"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserId uint
	Role   int
	jwt.StandardClaims
}

/*********************************************************
** 函数功能: 发放用户Token
** 日    期: 2021/7/10
**********************************************************/
func ReleaseToken(user model.User) (string, error) {
	jwtKey := []byte(viper.GetString("server.jwt_secret"))
	//token过期时间
	expirationTime := time.Now().Add(14 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			//发放时间等
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "danmu",
			Subject:   "token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

/*********************************************************
** 函数功能: 解析Token
** 日    期:2021/7/10
**********************************************************/
func ParseUserToken(tokenString string) (*jwt.Token, *Claims, error) {
	jwtKey := []byte(viper.GetString("server.jwt_secret"))
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, e error) {
		return jwtKey, nil
	})
	return token, claims, err
}
