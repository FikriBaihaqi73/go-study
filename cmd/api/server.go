package main

import "net/http"

type Server struct {
	httpServer *http.Server
}

func NewServer() *Server {
	router := NewRouter()

	return &Server{
		httpServer: &http.Server{
			Addr:    ":5000",
			Handler: router,
		},
	}
}

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}
