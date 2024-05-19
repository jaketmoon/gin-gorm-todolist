package Todo

import (
	"gin/internal/global/database"
	"gin/internal/global/errs"
	"gin/internal/global/log"
	"gin/internal/model"
	"github.com/gin-gonic/gin"
)

func Read(c *gin.Context) {
	var todolist []model.Todo
	if err := database.DB.Find(&todolist).Error; err != nil {
		errs.Fail(c, errs.DB_CRUD_ERROR.WithOrigin(err))
		log.SugarLogger.Error(err)
		return
	}
	errs.Success(c, "获取代办事项列表", todolist)
}
