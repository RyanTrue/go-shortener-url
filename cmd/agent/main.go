package main

import (
	"flag"
	"github.com/RyanTrue/go-shortener-url.git/internal/common/config"
	"github.com/RyanTrue/go-shortener-url.git/internal/common/server"
	"github.com/RyanTrue/go-shortener-url.git/internal/common/server/handler"
	"github.com/RyanTrue/go-shortener-url.git/internal/common/service"
	"log"
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
		log.Fatal(err)
	}
}
