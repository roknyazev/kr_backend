package main

import (
	"fmt"
	"orderManager/internal/config"
	"orderManager/internal/handler"
	"orderManager/internal/server"
	"orderManager/internal/service"
	"os"
)

func main(){

	Config, err := config.Init("configs")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sc		 	:= service.NewService(Config)
	Handlers 	:= handler.NewHandler(sc)
	srv 		:= server.NewServer(Config, Handlers.InitRoutes())
	err = srv.Run()
	if err != nil {
		return
	}
}


