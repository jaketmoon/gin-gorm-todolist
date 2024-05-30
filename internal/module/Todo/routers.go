package Todo

import (
	"gin/internal/global/middleware"
	"github.com/gin-gonic/gin"
)

func (u *ModuleTodo) InitRouter(r *gin.RouterGroup) {
	r.Use(middleware.Auth())
	r.POST("/", Add)
	r.PUT("/:id", Update)
	r.GET("/", Read)
	r.GET("/:id", Seek)
	r.DELETE("/:id", Delete)
}
