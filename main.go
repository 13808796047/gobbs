package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//加载模板
	r.LoadHTMLGlob("views/*")
	//处理静态资源
	//r.StaticFS("/asset", http.Dir("public"))
	r.Static("/asset", "./public")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Hello Gin",
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
