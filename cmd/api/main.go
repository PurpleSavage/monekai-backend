package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/purplesvage/moneka-ai/cmd/config"
	"github.com/purplesvage/moneka-ai/internal/audio"
	audioHttp "github.com/purplesvage/moneka-ai/internal/audio/in/http"
	audiosse "github.com/purplesvage/moneka-ai/internal/audio/in/sse"
	"github.com/purplesvage/moneka-ai/internal/auth"
	authHttp "github.com/purplesvage/moneka-ai/internal/auth/in/http"
	sharedvalidators "github.com/purplesvage/moneka-ai/internal/shared/in/validators"
	"github.com/purplesvage/moneka-ai/pkg/connection"
)

func main(){
	mainMux := http.NewServeMux()
	config.LoadEnvs()


	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.Envs.Host,
		config.Envs.DbUser,
		config.Envs.DbPassword,
		config.Envs.DbName,
		config.Envs.DbPort,
		config.Envs.SslMode,
	)
	db, err := connection.NewClient(dsn)
	
	if err != nil {
		log.Fatal("Error conectando a la base de datos:", err)
	}



	// ─── 1. INSTANCIAS GLOBALES COMPARTIDAS ───
    dtoValidator := sharedvalidators.NewDTOValidator()


	//--- Módulo AUTH ---
	authHandler := auth.Bootstrap(db)
	authMux := http.NewServeMux()	
	authHttp.MapRoutes(authMux,authHandler)
	mainMux.Handle("/auth/", http.StripPrefix("/auth", authMux))

	
	//--- Módulo AUDIO---
	audioHandler, audioSSEHandler := audio.Bootstrap(dtoValidator)
	audioMux:= http.NewServeMux()
	audioHttp.MapRoutes(audioMux,audioHandler)
	audiosse.MapSSERoutes(audioMux, audioSSEHandler)
	mainMux.Handle("/audio/", http.StripPrefix("/audio", audioMux))
	



	log.Println("Servidor iniciado en Piura: puerto 8080")
	if err := http.ListenAndServe(":8080", mainMux); err != nil {
		log.Fatal(err)
	}
}