package User

import (
	"gin/internal/global/casbin"
	"gin/internal/global/database"
	"gin/internal/global/errs"
	"gin/internal/global/log"
	"gin/internal/model"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.SugarLogger.Error(err)
		errs.Fail(c, errs.INVALID_REQUEST.WithOrigin(err))
		return
	}
	//检查输入是否出现问题
	if len(user.Name) == 0 || len(user.Password) == 0 || len(user.Email) == 0 {
		errs.Fail(c, errs.LOGIN_ERROR.WithTips("please check your name,email,and password"))
		return
	}
	//检查是否同名
	var v1 model.User
	result := database.DB.Where("name=?", user.Name).First(&v1)
	if result.Error == nil {
		errs.Fail(c, errs.LOGIN_ERROR.WithTips("repeated name"))
		return
	}
	//create user
	if err := database.DB.Create(&user).Error; err != nil {
		log.SugarLogger.Error(err)
		errs.Fail(c, errs.SERVE_INTERNAL.WithOrigin(err))
		return
	}
	err := casbin.Enforce.LinkUserWithPolicy(user.Name) //link user with policy
	if err != nil {
		log.SugarLogger.Error(err)
		errs.Fail(c, errs.SERVE_INTERNAL.WithOrigin(err))
		return
	}
	errs.Success(c, "register successfully")

}
