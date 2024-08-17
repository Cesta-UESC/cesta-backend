package main

import (
	"fmt"
	"log"

	"github.com/Cesta-UESC/cesta-backend/app/models"
	"github.com/Cesta-UESC/cesta-backend/plataform/database"
	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	db, err := database.DbConnection()
	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&models.User{})

	api := app.Group("/api") // /api

	users := api.Group("/users")

	users.Get("/", func(c fiber.Ctx) error {
		var user models.User
		db.First(&user)
		return c.SendString(fmt.Sprintf("id %d name %s email %s", user.ID, user.Name, user.Email))
	})

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
