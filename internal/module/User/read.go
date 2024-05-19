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

// 查看用户信息
func Read(c *gin.Context) {
	payload, exists := c.Get("Payload") // 从token中获取用户信息
	if !exists {
		errs.Fail(c, errs.UNTHORIZATION.WithTips("没有token，没有权限查看"))
		return
	}
	load := payload.(*jwt.Mycustomclaims)                                        // 将用户信息转换为jwt中的载荷,通过断言
	ok := casbin.GetEnforce().CheckUserPolicyForRead(load.User, "users", "read") // 检查用户是否有读取自己的信息的权限
	if !ok {
		errs.Fail(c, errs.UNTHORIZATION.WithTips("你的身份不允许查看该用户信息"))
		return
	}
	var user model.User
	tx := database.DB.Where("name = ?", load.User).First(&user)
	if tx.Error != nil {
		log.SugarLogger.Error(tx.Error)
		errs.Fail(c, errs.DB_CRUD_ERROR.WithTips("该用户不存在"))
	}
	errs.Success(c, user)
}
