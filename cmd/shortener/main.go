package main

import (
	f "github.com/RyanTrue/go-shortener-url.git/cmd/shortener/config"
	h "github.com/RyanTrue/go-shortener-url.git/cmd/shortener/httphandlers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func run(m *http.ServeMux) error {
	return http.ListenAndServe(`:8080`, m)
}

func main() {
	app := gin.Default()

	app.POST("/", func(c *gin.Context) {
		h.ShortenURL(c.Writer, c.Request)
	})
	app.GET("/:id", func(c *gin.Context) {
		h.GetOriginalURL(c.Writer, c.Request)
	})
	err := app.Run(f.ServerAddr)
	if err != nil {
		panic(err)
	}
}
