package server

import (
	"fmt"
	"gin/configs"
	"gin/internal/global/casbin"
	"gin/internal/global/database"
	"gin/internal/global/log"
	"gin/internal/module"
	"github.com/gin-gonic/gin"
)

func Init() {
	configs.Init()
	database.Init()
	casbin.Init()
	for _, m := range module.Modules {
		fmt.Println("Init Module: " + m.GetName())
		m.Init()
	}
}

func Run() {
	//r := gin.New()
	r := gin.Default()
	r.Use(log.Init())
	//告诉gin框架静态文件去哪里找
	r.Static("/static", "frontdevelop/static")
	//告诉gin框架模板文件引用的静态文件去哪里找
	r.LoadHTMLGlob("frontdevelop/templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)

	})
	for _, m := range module.Modules {
		fmt.Println("InitRouter: " + m.GetName())
		m.InitRouter(r.Group("/" + m.GetName()))
	}
	panic(r.Run(":8080"))
}
