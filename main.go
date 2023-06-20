package main

import (
	"sandbox-api/database"
	"sandbox-api/models"
	"sandbox-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	// Database
	database.InitDatabase()
	database.DBConn.AutoMigrate(&models.Users{})
	database.DBConn.AutoMigrate(&models.Logs{})

	// Route
	routes.Handle(app)
	app.Listen(":8000")
}
