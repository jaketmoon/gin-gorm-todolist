package module

import (
	"gin/internal/module/Todo"
	"gin/internal/module/User"
	"github.com/gin-gonic/gin"
)

type Module interface {
	GetName() string
	Init()
	InitRouter(r *gin.RouterGroup)
}

var Modules []Module

func RegisterModule(m Module) {
	Modules = append(Modules, m)
}

// init函数会在程序运行钱做必要的初始化工作
func init() {
	RegisterModule(&User.ModuleUser{})
	RegisterModule(&Todo.ModuleTodo{})
}
