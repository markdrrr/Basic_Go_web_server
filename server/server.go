package server

import (
	"fmt"
	"net/http"
	"server_go/handler"
)


type Server struct {
	Host    string
	Port    int
}

func NewServer(host string, port int) *Server {
	return &Server{
		Host:    host,
		Port:    port,
	}
}

func (s *Server) Run() error {
	address := fmt.Sprintf("%s:%d", s.Host, s.Port)
	fmt.Println("Start server", address)
	mux := s.initRouter()
	return http.ListenAndServe(address, mux)
}

func (s *Server) initRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Index)
	mux.HandleFunc("/add/", handler.Add)
	mux.HandleFunc("/delete/", handler.Delete)
	mux.HandleFunc("/students/", handler.GetStudents)
	mux.HandleFunc("/student/", handler.GetStudent)
	return mux
}
