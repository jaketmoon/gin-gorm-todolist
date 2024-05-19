package jwt

import (
	"gin/configs"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Payload struct {
	Authorized bool   `json:"authorized"`
	User       string `json:"user"`
}
type Mycustomclaims struct {
	Payload
	jwt.RegisteredClaims
}

func NewToken(name string) (string, error) {
	claims := &Mycustomclaims{
		Payload: Payload{
			Authorized: true,
			User:       name,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    configs.JwtSettings.Issuer,
			Subject:   configs.JwtSettings.Subject,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // jwt.NewNumericDate 可以创建一个符合JWT标准的时间格式,这里是24小时
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(configs.JwtSettings.SecretKey))
	return tokenString, err
}
