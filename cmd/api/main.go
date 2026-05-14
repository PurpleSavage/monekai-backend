package main
import (
	"fmt"
	"log"
	"net/http"
	"github.com/purplesvage/moneka-ai/cmd/config"
	"github.com/purplesvage/moneka-ai/pkg/connection"
	"github.com/purplesvage/moneka-ai/internal/auth"
	authHttp "github.com/purplesvage/moneka-ai/internal/auth/in/http"
)

func main(){
	mainMux := http.NewServeMux()
	config.LoadEnvs()


	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.Envs.Host,
		config.Envs.User, 
		config.Envs.DbPassword,
		config.Envs.DbName,
		config.Envs.DbPort,
		config.Envs.SslMode,
	)
	db, err := connection.NewClient(dsn)
	
	if err != nil {
		log.Fatal("Error conectando a la base de datos:", err)
	}

	//--- Módulo AUTH ---
	authHandler := auth.Bootstrap(db)
	authMux := http.NewServeMux()	
	authHttp.MapRoutes(authMux,authHandler)
	mainMux.Handle("/auth/", http.StripPrefix("/auth", authMux))
	//
	



	log.Println("Servidor iniciado en Piura: puerto 8080")
	if err := http.ListenAndServe(":8080", mainMux); err != nil {
		log.Fatal(err)
	}
}