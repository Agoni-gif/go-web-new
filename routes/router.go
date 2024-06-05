package routes

import (
	"github.com/gin-gonic/gin"
	"go-web-new/middleware"
	"go-web-new/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	// 初始化总路由
	DefRouter := r.Group("")
	{
		V1RouterInit(DefRouter)
	}
	r.LoadHTMLGlob("static/admin/index.html")
	r.Static("admin", "static/admin")
	r.GET("admin", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	_ = r.Run(utils.HttpPort)
}
