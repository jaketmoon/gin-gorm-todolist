package Todo

import (
	"gin/internal/global/database"
	"gin/internal/global/errs"
	"gin/internal/global/log"
	"gin/internal/model"
	"github.com/gin-gonic/gin"
)

func Add(c *gin.Context) {
	var Todo model.Todo
	if err := c.ShouldBindJSON(&Todo); err != nil {
		errs.Fail(c, errs.INVALID_REQUEST.WithOrigin(err))
		log.SugarLogger.Error(err)
		return
	}

	if err := database.DB.Create(&Todo).Error; err != nil {
		errs.Fail(c, errs.DB_CRUD_ERROR.WithOrigin(err))
		log.SugarLogger.Error(err)
		return
	}
	errs.Success(c, "添加代办事项为", Todo)

}
