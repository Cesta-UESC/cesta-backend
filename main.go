package main

import (
	"log"

	"github.com/Cesta-UESC/cesta-backend/controller"
	"github.com/Cesta-UESC/cesta-backend/model"
	"github.com/Cesta-UESC/cesta-backend/repository"
	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	db, err := model.DbConnection()
	if err != nil {
		panic("Failed to connect to database")
	}

	repository.SetDefault(db)

	// db.AutoMigrate(&models.User{})

	controller.SetDefault(app)

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
