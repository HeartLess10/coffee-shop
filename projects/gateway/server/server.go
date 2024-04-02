package server

import (
	"fmt"
	"net/http"
	"strconv"
)

type Server struct {
	port    int
	domName string
	l       *MyLogger
}

func (s *Server) Serve() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello world")
	})

	if err := http.ListenAndServe("localhost:"+strconv.Itoa(s.port), mux); err != nil {
		fmt.Println(err.Error())
	}
}
