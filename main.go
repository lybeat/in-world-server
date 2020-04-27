package main

import (
	"fmt"
	"in-world-server/model"
	"in-world-server/pkg/setting"
	"in-world-server/router"
	"net/http"
)

func main() {
	setting.Setup()
	model.Setup()

	router := router.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
