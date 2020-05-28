package main

import (
	"explore/pkg/setting"
	"explore/routers"
	"fmt"
	"net/http"
)

func main() {
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.Config().Server.Http_port),
		Handler:        router,
		ReadTimeout:    setting.Config().Server.Read_timeout,
		WriteTimeout:   setting.Config().Server.Write_timeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}