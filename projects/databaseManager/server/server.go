package server

import (
	"fmt"
	"net/http"

	"github.com/HeartLess10/coffee-shop/coffee-shop/utils/customLogger"
)

type Server interface {
	Serve()
}

type server struct {
	port    string
	domName string
	l       customLogger.CustomLogger
}

func NewServer(port string, domName string, l customLogger.CustomLogger) Server {
	return &server{domName: domName, port: port, l: l}
}

func (s *server) Serve() {
	s.l.Message("Starting database manager server address: " + s.domName + ":" + s.port)
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello world")
	})

	if err := http.ListenAndServe(s.domName+":"+s.port, mux); err != nil {
		s.l.Error(err)
	}
}
