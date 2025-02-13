package main

import (
	"embed"

	//choisir le domaine 
	"domain.local/web/server"
)

//go:embed templates
var templates embed.FS

func main() {
	srv := server.ServerCfg{
		ListenPort: ":8080",
		Templates:  templates,
	}

	srv.Run()
}
