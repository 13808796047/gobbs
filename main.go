package main

import (
	"fmt"
	"gobbs/pkg/setting"
	"gobbs/routes"
	"net/http"
)

func main() {
	router := routes.InitRouter()
	//r.Run() // 监听并在 0.0.0.0:8080 上启动服务
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
