package common

import (
	"regexp"
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

const (
	AccessTypeToken  = "access"
	RefreshTypeToken = "refresh"
)

/*********************************************************
** 函数功能: 发放用户Token
** 返 回 值: refreshToken, accessToken, error
** 日    期: 2022年8月25日14:37:20
**********************************************************/
func ReleaseToken(user model.User) (string, string, error) {
	//生成refreshToken
	refreshToken, err := ReleaseRefreshToken(user)
	if err != nil {
		return "", "", err
	}
	//生成accessToken
	accessToken, err := ReleaseAccessToken(user)
	if err != nil {
		return "", "", err
	}

	return refreshToken, accessToken, nil
}

/*********************************************************
** 函数功能: 解析Token
** 日    期: 2021/7/10
**********************************************************/
func ParseUserToken(tokenString, tokenType string) (*jwt.Token, *Claims, error, bool) {
	var jwtKey []byte
	if tokenType == AccessTypeToken {
		jwtKey = []byte(viper.GetString("server.access_jwt_secret"))
	} else {
		jwtKey = []byte(viper.GetString("server.refresh_jwt_secret"))
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, e error) {
		return jwtKey, nil
	})

	isExpired := false
	// 判断err是否为token过期
	if err != nil {
		reg := regexp.MustCompile(`token is expired`)
		if reg.MatchString(err.Error()) {
			isExpired = true
		}
	}
	return token, claims, err, isExpired
}

/*********************************************************
** 函数功能: 发放refreshToken
** 返 回 值: refreshToken, error
** 日    期: 2022年8月25日15:28:29
**********************************************************/
func ReleaseRefreshToken(user model.User) (string, error) {
	refreshJwtKey := []byte(viper.GetString("server.refresh_jwt_secret"))
	// token过期时间
	expirationTime := time.Now().Add(14 * 24 * time.Hour) // 14天有效

	refreshClaims := &Claims{
		UserId: user.ID,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			//发放时间等
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "danmu",
			Subject:   RefreshTypeToken,
		},
	}

	return GenerateToken(refreshJwtKey, refreshClaims)
}

/*********************************************************
** 函数功能: 发放accessToken
** 返 回 值: accessToken, error
** 日    期: 2022年8月25日15:19:26
**********************************************************/
func ReleaseAccessToken(user model.User) (string, error) {
	accessJwtKey := []byte(viper.GetString("server.access_jwt_secret"))
	// token过期时间
	expirationTime := time.Now().Add(5 * time.Minute) // 5分钟有效

	accessClaims := &Claims{
		UserId: user.ID,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			//发放时间等
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "danmu",
			Subject:   AccessTypeToken,
		},
	}

	return GenerateToken(accessJwtKey, accessClaims)
}

/*********************************************************
** 函数功能: 生成token字符串
** 返 回 值: token, error
** 日    期: 2022年8月25日15:26:01
**********************************************************/
func GenerateToken(key []byte, claims *Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
