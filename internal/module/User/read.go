package User

import (
	"gin/internal/global/database"
	"gin/internal/global/errs"
	"gin/internal/global/jwt"
	"gin/internal/global/log"
	"gin/internal/model"
	"github.com/gin-gonic/gin"
)

// 查看用户信息
func Read(c *gin.Context) {
	var user model.User
	user = jwt.Getcurrentuser(c)
	tx := database.DB.Where("id = ?", user.ID).First(&user)
	if tx.Error != nil {
		log.SugarLogger.Error(tx.Error)
		errs.Fail(c, errs.DB_CRUD_ERROR.WithTips("该用户不存在"))
	}
	errs.Success(c, user)
}
