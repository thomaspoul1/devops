package server

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"domain.local/web/handlers"
)

type Configuration struct {
	Port string
	Templates  embed.FS
}

func (config *Configuration) Start() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.IndexHandler(w, r, config.Templates)
	})
	mux.HandleFunc("/healthz", handlers.HealthHandler)

	handler := applyMiddlewares(mux)

	log.Printf("Server running at http://localhost%s\n", config.Port)
	log.Fatal(http.ListenAndServe(config.Port, handler))
}
