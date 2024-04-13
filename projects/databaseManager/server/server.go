package server

import (
	"net/http"

	"github.com/HeartLess10/coffee-shop/coffee-shop/projects/databaseManager/dbManager"
	"github.com/HeartLess10/coffee-shop/coffee-shop/projects/databaseManager/server/handlers"
	"github.com/HeartLess10/coffee-shop/coffee-shop/projects/databaseManager/server/middleware"
	"github.com/HeartLess10/coffee-shop/coffee-shop/utils/customLogger"
)

type Server interface {
	Serve()
}

type server struct {
	m   middleware.Middleware
	s   http.Server
	l   customLogger.CustomLogger
	dbm *dbManager.DbManager
}

func NewServer(port string, domName string, dbm *dbManager.DbManager, l customLogger.CustomLogger) Server {
	m := middleware.NewMiddleware(l)
	s := http.Server{
		Addr: domName + ":" + port,
	}

	return &server{dbm: dbm, m: m, s: s, l: l}
}

func (s *server) Serve() {
	s.l.Message("Starting database manager server address: " + s.s.Addr)
	router := http.NewServeMux()
	h := handlers.NewHandler(s.dbm)
	initMuxHandlers(router, h)

	s.s.Handler = s.m.Logging(router)
	if err := s.s.ListenAndServe(); err != nil {
		panic(err)
	}
}

func initMuxHandlers(router *http.ServeMux, h handlers.Handler) {
	router.HandleFunc("POST /addUser", h.AddUser)
}
