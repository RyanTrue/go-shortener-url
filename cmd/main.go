package main

import (
	"flag"
	"github.com/RyanTrue/go-shortener-url.git/internal/app/config"
	"github.com/RyanTrue/go-shortener-url.git/internal/app/handler"
	"github.com/RyanTrue/go-shortener-url.git/internal/app/server"
	"github.com/RyanTrue/go-shortener-url.git/internal/app/service"
)

func main() {
	appConfig := config.AppConfig{}
	appConfig.InitAppConfig()
	flag.Parse()

	repo := make(map[string]string)

	services := service.NewServiceContainer(repo, appConfig)
	handler := handler.NewHandler(services)
	server := &server.Server{}

	if err := server.Run(appConfig.Server.ServerAddr, handler.InitRoutes()); err != nil {
		panic(err)
	}
}
