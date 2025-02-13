package handlers

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"

	"domain.local/web/logger"
	"github.com/spf13/viper"
)

// IndexHandler gère la page d'accueil.
func IndexHandler(w http.ResponseWriter, r *http.Request, templates embed.FS) {
	log := logger.NewLogger()  // Utilisation du logger personnalisé.
	// Chargement et parsing des templates HTML à partir de l'embedded file system.
	tmpl, err := template.ParseFS(templates, "templates/base.html", "templates/index.html")
	if err != nil {
		log.Error(fmt.Sprintf("Failed to parse templates: %v", err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Configuration de Viper pour la gestion de la configuration.
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/opt")  // Spécifier le chemin d'accès au fichier de configuration.
	if err := viper.ReadInConfig(); err != nil {
		log.Error(fmt.Sprintf("Failed to load configuration: %v", err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Récupération de la valeur de configuration pour l'image.
	img := viper.GetString("image")

	// Exécution du template avec les données de l'image.
	if err := tmpl.Execute(w, map[string]interface{}{"img": img}); err != nil {
		log.Error(fmt.Sprintf("Failed to execute template: %v", err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
