package User

import "github.com/gin-gonic/gin"

func (u *ModuleUser) InitRouter(r *gin.RouterGroup) {
	r.POST("/register", Register)
	r.POST("/login", Login)
	r.PUT("/update", Update)
	r.GET("/read", Read)
	r.DELETE("/delete", Delete)
}
