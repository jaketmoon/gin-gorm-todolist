package User

import (
	"gin/internal/global/database"
	"gin/internal/global/errs"
	"gin/internal/global/jwt"
	"gin/internal/model"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	var v2 model.User
	if err := database.DB.Where("name=?", name).First(&v2).Error; err != nil {
		errs.Fail(c, errs.DB_CRUD_ERROR.WithTips("该用户不存在，请注册后登录"))
		return
	}
	if password != v2.Password {
		errs.Fail(c, errs.LOGIN_ERROR.WithTips("密码错误"))
		return
	}
	token, err := jwt.NewToken(name)
	if err != nil {
		errs.Fail(c, errs.UNTHORIZATION.WithOrigin(err))
		return
	}
	errs.Success(c, v2, map[string]string{"token": token})
}
