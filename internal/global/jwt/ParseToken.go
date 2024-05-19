package jwt

import (
	"errors"
	"gin/configs"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

// 解析token

func ParseToken(token string) (*Mycustomclaims, error) {
	// 解析方式需要添加 Bearer token模式
	tokenParts := strings.Split(token, " ") //通过空格分隔出两个部分，并且存入数组之中
	if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
		return nil, errors.New("invalid token format,you need add bearer")
	}
	tokenString := tokenParts[1]
	// 解析后续token
	claims := &Mycustomclaims{}
	// 是*token和string之间的转换
	// 这是一个回调函数具体结构就是 jwt.Parse(string,KeyFunc)
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法 HMAC-SHA56签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected Signing Method")
		}
		return []byte(configs.JwtSettings.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return claims, nil
}
