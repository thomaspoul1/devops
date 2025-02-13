package server

import (
	"embed"
	"net/http"
	"time"

	"domain.local/web/handlers"
	"domain.local/web/middleware"
)

type ServerConfig struct {
	ListenPort  string
	IdleTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Templates    embed.FS
}

// Constructeur pour ServerConfig
func NewServer(listenPort string) *ServerConfig {
	return &ServerConfig{
		ListenPort:  listenPort,
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}

// Méthode pour démarrer le serveur
func (config *ServerConfig) Run() {
	mux := http.NewServeMux()

	// Définition des gestionnaires de routes
	mux.HandleFunc("/healthz", handlers.HealthHandler)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.IndexHandler(w, r, config.Templates)
	})

	// Appliquer des middlewares pour le logging et CORS
	muxWithMiddlewares := middleware.MiddleLogger(middleware.MiddleCors(mux))

	// Configuration du serveur HTTP
	server := &http.Server{
		Addr:         config.ListenPort,
		Handler:      muxWithMiddlewares,
		IdleTimeout:  config.IdleTimeout,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
	}

	// Démarrage du serveur
	server.ListenAndServe()
}
