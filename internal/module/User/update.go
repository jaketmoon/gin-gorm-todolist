package User

import (
	"gin/internal/global/casbin"
	"gin/internal/global/database"
	"gin/internal/global/errs"
	"gin/internal/global/jwt"
	"gin/internal/model"
	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	password := c.PostForm("OldPassword")
	NewPassword := c.PostForm("NewPassword")
	payload, exists := c.Get("Payload")
	if !exists {
		errs.Fail(c, errs.LOGIN_ERROR.WithTips("没有token，没有权限修改"))
		return
	}
	load := payload.(*jwt.Mycustomclaims)
	ok := casbin.GetEnforce().CheckUserPolicyForRead(load.User, "users", "write")
	if !ok {
		errs.Fail(c, errs.UNTHORIZATION.WithTips("你的身份没有权限修改"))
		return
	}
	var user model.User
	tx := database.DB.Where("name = ?", load.User).First(&user)
	if tx.Error != nil {
		errs.Fail(c, errs.DB_CRUD_ERROR.WithTips("该用户不存在"))
		return
	}

	if password != user.Password {
		errs.Fail(c, errs.INVALID_REQUEST.WithTips("旧密码错误"))
		return
	}

}
