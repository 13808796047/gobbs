package routes

import (
	"github.com/gin-gonic/gin"
	"gobbs/controllers"
)

func WebRouter(r *gin.RouterGroup) {
	// 主页
	r.GET("/", controllers.Index)
}
