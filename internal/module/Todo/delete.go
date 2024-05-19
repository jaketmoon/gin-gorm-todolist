package Todo

import (
	"gin/internal/global/database"
	"gin/internal/global/errs"
	"gin/internal/global/log"
	"gin/internal/model"
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	var Todo model.Todo
	id, ok := c.Params.Get("id")
	if !ok {
		errs.Fail(c, errs.INVALID_REQUEST.WithTips("id不存在"))
		return
	}
	// 先从数据库中查询要删除的Todo
	if err := database.DB.Where("id = ?", id).First(&Todo).Error; err != nil {
		errs.Fail(c, errs.DB_BASE_ERROR.WithOrigin(err))
		log.SugarLogger.Error(err)
		return
	}

	// 进行删除操作
	if err := database.DB.Where("id = ?", id).Delete(&Todo).Error; err != nil {
		errs.Fail(c, errs.DB_CRUD_ERROR.WithOrigin(err))
		log.SugarLogger.Error(err)
		return
	}
	errs.Success(c, "删除代办事项为", Todo)
}
