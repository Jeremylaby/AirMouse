package server

import "fmt"

type Server struct{}

func New() *Server {
	return &Server{}
}

func (s *Server) Start() error {
	fmt.Println("Server started")
	return nil
}

func (s *Server) Stop() error {
	fmt.Println("Server stopped")
	return nil
}
