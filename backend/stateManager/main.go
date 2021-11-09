package main

import (
	"fmt"
	"os"
	"stateManager/internal"
	"stateManager/internal/config"
	"stateManager/internal/handlers"
	"stateManager/internal/server"
)

func main() {
	// Создание структуры конфигурации
	Config, err := config.Init("configs")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	Handlers := handlers.NewHandler()
	go internal.Ff(Handlers)
	srv := server.NewServer(Config, Handlers.InitRoutes())
	err = srv.Run()
	if err != nil {
		return
	}
}

