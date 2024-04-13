package middleware

import (
	"fmt"
	"net/http"

	"github.com/HeartLess10/coffee-shop/coffee-shop/utils/customLogger"
)

type Middleware interface {
	Logging(next http.Handler) http.Handler
}

type middleware struct {
	l customLogger.CustomLogger
}

func NewMiddleware(l customLogger.CustomLogger) Middleware {
	return &middleware{l: l}
}

func (m *middleware) Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.l.Debug(fmt.Sprintf("%s %s", r.Method, r.URL.Path))
		next.ServeHTTP(w, r)
	})
}
