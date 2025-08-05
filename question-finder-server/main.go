package main

import (
	"log"
	AppConfig "question_finder/app"
	LoginHandler "question_finder/login_register/handlers"
	ProfileSettingsHandler "question_finder/profile_settings/handlers"
	"question_finder/utils/postgres"

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
	appConfig := &AppConfig.PostgresConfig{
		DB: conn,
	}

	loginHandler := LoginHandler.ConstructLoginRegisterHandler(appConfig)
	LoginHandler.RegisterRoutes(router, loginHandler)

	profileSettingsHandler := ProfileSettingsHandler.ConstructProfileSettingsHandler()
	ProfileSettingsHandler.ProfileSettingsRoutes(router, profileSettingsHandler)

	router.Run(":8080") // http://localhost:8080
}
