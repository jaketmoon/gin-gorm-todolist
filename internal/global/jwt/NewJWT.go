package jwt

import (
	"gin/configs"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Payload struct {
	Authorized bool   `json:"authorized"`
	User       string `json:"user"`
	Userid     string `json:"userid"`
}
type Mycustomclaims struct {
	Payload
	jwt.RegisteredClaims
}

func NewToken(name string, userid string) (string, error) {
	claims := &Mycustomclaims{
		Payload: Payload{
			Authorized: true,
			User:       name,
			Userid:     userid,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    configs.JwtSettings.Issuer,
			Subject:   configs.JwtSettings.Subject,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 10000)), // jwt.NewNumericDate 可以创建一个符合JWT标准的时间格式,这里是24小时
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //创建一个新的令牌,这里是HS256算法,claims是一个对象
	tokenString, err := token.SignedString([]byte(configs.JwtSettings.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
