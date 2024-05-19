package User

import (
	"gin/internal/global/casbin"
	"gin/internal/global/database"
	"gin/internal/global/errs"
	"gin/internal/global/jwt"
	"gin/internal/model"
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	password := c.PostForm("password")
	// 从token中获取用户信息
	payload, exists := c.Get("Payload")
	if !exists {
		errs.Fail(c, errs.UNTHORIZATION.WithTips("没有token，没有权限修改"))
		return
	}
	load := payload.(*jwt.Mycustomclaims)
	// 检查用户权限
	ok := casbin.GetEnforce().CheckUserPolicyForRead(load.User, "users", "write")
	if !ok {
		errs.Fail(c, errs.UNTHORIZATION.WithTips("没有权限修改"))
		return
	}
	// 删除用户
	var user model.User
	tx := database.DB.Where("name = ?", load.User).First(&user)
	if tx.Error != nil {
		errs.Fail(c, errs.DB_CRUD_ERROR.WithTips("该用户不存在"))
	}
	if user.Password != password {
		errs.Fail(c, errs.INVALID_REQUEST.WithTips("密码错误"))
		return
	}
	result := database.DB.Delete(&user)
	if result.Error != nil {
		errs.Fail(c, errs.DB_CRUD_ERROR.WithOrigin(result.Error))
		return
	}
	errs.Success(c, "注销成功")
}
