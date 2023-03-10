package main

import (
	"fmt"
	"log"
	"os"

	"gostarter/docs"
	"gostarter/server"
	"gostarter/server/database"

	"github.com/joho/godotenv"
)

// @title go starter
// @version 1.0
// @description This is a starter pack with go

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @BasePath /
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	database := database.InitDB()
	connection, _ := database.DB()
	defer func() {
		if err := connection.Close(); err != nil {
			log.Print(err)
		}
	}()

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))

	app := server.NewServer(database)
	server.ConfigureRoutes(app)

	if err := app.Run(os.Getenv("PORT")); err != nil {
		log.Print(err)
	}
}
