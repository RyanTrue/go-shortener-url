package main

import (
	"flag"
	"github.com/RyanTrue/go-shortener-url.git/cmd/shortener/config"
	"github.com/RyanTrue/go-shortener-url.git/cmd/shortener/handler"
	"github.com/RyanTrue/go-shortener-url.git/cmd/shortener/server"
	"github.com/RyanTrue/go-shortener-url.git/cmd/shortener/service"
)

func main() {
	appConfig := config.AppConfig{}
	appConfig.InitAppConfig()
	flag.Parse()

	repo := make(map[string]string)

	services := service.NewServiceContainer(repo, appConfig)
	handler := handler.NewHandler(services)
	server := new(server.Server)

	if err := server.Run(appConfig.Server.ServerAddr, handler.InitRoutes()); err != nil {
		panic(err)
	}
}
