package handlers

import (
	"fmt"
	"net/http"
	"time"
)

// HealthHandler fournit une réponse basique de vérification de santé.
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	// Définition des en-têtes avant toute opération.
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	
	// Tu pourrais inclure ici d'autres vérifications, par exemple, la connectivité à la base de données, le statut des API externes, etc.
	// Pour le moment, il renvoie simplement un statut OK.
	heureReponse := time.Now().Format(time.RFC3339)
	reponse := fmt.Sprintf("Statut : OK\nHeure : %s", heureReponse)
	
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(reponse))
}
