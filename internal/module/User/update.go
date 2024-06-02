package User

import (
	"gin/internal/global/casbin"
	"gin/internal/global/database"
	"gin/internal/global/errs"
	"gin/internal/global/jwt"
	"gin/internal/global/log"
	"gin/internal/model"
	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	password := c.PostForm("OldPassword")
	NewPassword := c.PostForm("NewPassword")
	var user model.User
	user = jwt.Getcurrentuser(c)
	ok := casbin.GetEnforce().CheckUserPolicyForRead(user.Name, "users", "write")
	if !ok {
		errs.Fail(c, errs.UNTHORIZATION.WithTips("你的身份没有权限修改"))
		return
	}
	tx := database.DB.Where("id = ?", user.ID).First(&user)
	if tx.Error != nil {
		errs.Fail(c, errs.DB_CRUD_ERROR.WithTips("该用户不存在"))
		return
	}

	if password != user.Password {
		errs.Fail(c, errs.INVALID_REQUEST.WithTips("旧密码错误"))
		return
	}

	tx1 := database.DB.Model(&model.User{}).Where("name = ?", user.Name).Update("password", NewPassword)
	if tx1.Error != nil {
		log.SugarLogger.Error(tx1.Error)
		errs.Fail(c, errs.DB_CRUD_ERROR.WithTips("修改失败"))
		return
	}
	errs.Success(c, user, NewPassword, "请妥善保存你的密码")
}
