package middleware

import (
	"gin/internal/global/jwt"
	"gin/internal/global/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取token
		token := c.GetHeader("Authorization")
		if token == "" {
			// 如果token为空，返回错误信息
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Failed to fetch token",
			})
			c.Abort()
			return
		}
		// 如果token不为空，继续执行
		parseToken, err := jwt.ParseToken(token)
		if err != nil {
			// 如果解析token失败，返回错误信息
			log.SugarLogger.Error(err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg":   "Failed to Auth",
				"error": err,
			})
			c.Abort()
			return
		}
		c.Set("Payload", parseToken)
		c.Next() // 继续执行
	}
}
