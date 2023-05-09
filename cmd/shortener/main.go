package main

import (
	"github.com/RyanTrue/go-shortener-url.git/cmd/shortener/config"
	"github.com/RyanTrue/go-shortener-url.git/cmd/shortener/handler"
	"github.com/RyanTrue/go-shortener-url.git/cmd/shortener/server"
	"github.com/RyanTrue/go-shortener-url.git/cmd/shortener/service"
)

var repo = make(map[string]string)

func main() {
	config.InitServerConfig()

	services := service.NewService(repo)
	handler := handler.NewHandler(services)
	server := new(server.Server)

	if err := server.Run(config.ServerAddr, handler.InitRoutes()); err != nil {
		panic(err)
	}
}
