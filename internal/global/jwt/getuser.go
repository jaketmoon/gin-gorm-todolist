package jwt

import (
	"gin/internal/global/database"
	"gin/internal/global/errs"
	"gin/internal/global/log"
	"gin/internal/model"
	"github.com/gin-gonic/gin"
)

func Getcurrentuser(c *gin.Context) model.User {
	payload, exists := c.Get("Payload")
	if !exists {
		errs.Fail(c, errs.UNTHORIZATION.WithTips("没有token，无法从中获取用户信息"))
		c.Abort()
		return model.User{}
	}
	load := payload.(*Mycustomclaims)
	var user model.User
	tx := database.DB.Where("id = ?", load.Userid).First(&user)
	if tx.Error != nil {
		log.SugarLogger.Error(tx.Error)
		errs.Fail(c, errs.DB_CRUD_ERROR.WithTips("该用户id不存在"))
		c.Abort()
		return model.User{}
	}
	return user
}
