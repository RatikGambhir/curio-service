package main

import (
	AppConfig "curio/app"
	Routing "curio/routing"
	"curio/utils/postgres"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()
	err := godotenv.Load() // automatically loads .env from root
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	conn, err := postgres.ConnectToProdPostgres()
	if err != nil {
		log.Fatal("Failed to connect to postgres", err)
	}
	defer conn.Close()
	appConfig := &AppConfig.AppConfig{
		PostgresConfig: &AppConfig.PostgresConfig{
			DB: conn,
		},
	}

	Routing.RegisterRoutes(router, appConfig)

	// loginHandler := LoginHandler.ConstructLoginRegisterHandler(appConfig)
	// LoginHandler.RegisterRoutes(router, loginHandler)

	// profileSettingsHandler := ProfileSettingsHandler.ConstructProfileSettingsHandler()
	// ProfileSettingsHandler.ProfileSettingsRoutes(router, profileSettingsHandler)

	router.Run(":8080") // http://localhost:8080
}
