package routes

import (
	"github.com/gin-gonic/gin"
	"gobbs/pkg/setting"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	//加载模板
	r.LoadHTMLGlob("views/**/*")
	//处理静态资源
	//r.StaticFS("/asset", http.Dir("public"))
	r.Static("/asset", "./public")
	r.GET("/test", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Hello Gin",
		})
	})
	return r
}
