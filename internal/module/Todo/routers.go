package Todo

import "github.com/gin-gonic/gin"

func (u *ModuleTodo) InitRouter(r *gin.RouterGroup) {
	r.POST("/", Add)
	r.PUT("/:id", Update)
	r.GET("/", Read)
	r.GET("/:id", Seek)
	r.DELETE("/:id", Delete)
}
