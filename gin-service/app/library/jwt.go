package library

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/putyy/ai-share/config"
	"time"
)

var jwtSecretApi = []byte(config.Jwt.JwtSecret)
var jwtSecretAdmin = []byte(config.Jwt.JwtSecretAdmin)

type ApiClaims struct {
	Uid     int    `json:"u"`
	Version string `json:"vs"`
	Type    int    `json:"t"` // 登录类型 1账户密码 2小程序授权
	Vip     int    `json:"v"`
	jwt.StandardClaims
}

func GenerateApiToken(claims ApiClaims) (string, error) {
	expireTime := time.Now().Add(48 * time.Hour)
	claims.Version = config.App.Version
	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: expireTime.Unix(),
		Issuer:    "ai-share-server",
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtSecretApi)
}

func ParseApiToken(token string) (*ApiClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &ApiClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretApi, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*ApiClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

type AdminClaims struct {
	Uid int `json:"uid"`
	jwt.StandardClaims
}

func GenerateAdminToken(uid int) (string, error) {
	expireTime := time.Now().Add(48 * time.Hour)
	claims := AdminClaims{
		uid,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "ai-share-server",
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtSecretAdmin)
}

func ParseAdminToken(token string) (*AdminClaims, error) {

	tokenClaims, err := jwt.ParseWithClaims(token, &AdminClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretAdmin, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*AdminClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
