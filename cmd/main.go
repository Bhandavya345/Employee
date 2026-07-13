package main

import (
	"fmt"
	"log"
	"net/http"

	database "github.com/Bhandavya345/Employee/DB"
	models "github.com/Bhandavya345/Employee/model"
	"github.com/joho/godotenv"

	routes "github.com/Bhandavya345/Employee/Routes"
)

func main() {

	// Load .env

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect Database
	database.ConnectDB()

	database.DB.AutoMigrate(
		&models.Employee{},
	)

	routes.RegisterRoutes()

	fmt.Println("Server Started at :8080")

	http.ListenAndServe(":8080", nil)

	
}
