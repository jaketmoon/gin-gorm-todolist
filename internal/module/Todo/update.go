package Todo

import (
	"gin/internal/global/database"
	"gin/internal/global/errs"
	"gin/internal/global/log"
	"gin/internal/model"
	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	var Todo model.Todo
	id, ok := c.Params.Get("id")
	if !ok {
		errs.Fail(c, errs.INVALID_REQUEST.WithTips("id不存在"))
		return

	}
	if err := database.DB.Where("id = ?", id).First(&Todo).Error; err != nil {
		errs.Fail(c, errs.DB_BASE_ERROR.WithOrigin(err))
		log.SugarLogger.Error(err)
		return
	}
	// 把status取反
	Todo.Status = !Todo.Status
	errs.Success(c, "更新代办事项为", Todo)
}
