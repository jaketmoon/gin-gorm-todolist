package User

import (
	"gin/internal/global/middleware"
	"github.com/gin-gonic/gin"
)

func (u *ModuleUser) InitRouter(r *gin.RouterGroup) {
	r.POST("/register", Register)
	r.POST("/login", Login)
	r.PUT("/update", middleware.Auth(), Update)
	r.GET("/read", middleware.Auth(), Read)
	r.DELETE("/delete", middleware.Auth(), Delete)
}
