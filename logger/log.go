package server

import (
	"embed"
	"net/http"
	"time"

	"domain.local/web/handlers"
)

type ServerConfig struct {
	ListenPort   string
	IdleTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Templates    embed.FS
}

// NewServer initialise une nouvelle configuration de serveur.
func NewServer(listenPort string) *ServerConfig {
	return &ServerConfig{
		ListenPort:   listenPort,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}

// Run démarre le serveur avec les configurations spécifiées.
func (config *ServerConfig) Run() {
	mux := http.NewServeMux()

	// Configuration des routes
	mux.HandleFunc("/healthz", handlers.HealthHandler)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.IndexHandler(w, r, config.Templates)
	})

	// Application des middlewares pour les logs et CORS
	muxWrapped := middleLogger(middleCors(mux))

	// Configuration et démarrage du serveur HTTP
	server := &http.Server{
		Addr:         config.ListenPort,
		Handler:      muxWrapped,
		IdleTimeout:  config.IdleTimeout,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
	}
	server.ListenAndServe()
}
