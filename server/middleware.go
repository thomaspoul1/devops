package server

import (
	"net/http"
	"time"
	"domain.local/web/logger"
)

type MiddlewareLogger struct {
	handler http.Handler
}

// Appliquer CORS à toutes les réponses
func middleCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, UPDATE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Cache-Control", "no-cache")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Logger Middleware pour suivre les requêtes
func (ml *MiddlewareLogger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	log := logger.NewLogger()
	ml.handler.ServeHTTP(w, r)
	log.InfoServerRequest(r.Method, r.URL.Path, time.Since(start).String())
}

// Crée un nouveau MiddlewareLogger
func middleLogger(handler http.Handler) *MiddlewareLogger {
	return &MiddlewareLogger{handler}
}
