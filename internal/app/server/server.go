package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:    port,
		Handler: handler,
	}
	err := s.httpServer.ListenAndServe()
	if err != nil {
		fmt.Printf("Server failed: ", err.Error())
	}
	return s.httpServer.ListenAndServe()
}
