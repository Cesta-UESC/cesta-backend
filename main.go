package main

import (
	"log"
	"time"

	"github.com/Cesta-UESC/cesta-backend/controllers"
	"github.com/Cesta-UESC/cesta-backend/middlewares"
	"github.com/Cesta-UESC/cesta-backend/models"
	"github.com/Cesta-UESC/cesta-backend/repositories"
	"github.com/go-playground/validator"

	"github.com/gofiber/fiber/v3"
)

const VERSION = "0.0.1"
const APP_NAME = "cesta-app"

type structValidator struct {
	validate *validator.Validate
}

func (v *structValidator) Validate(out any) error {
	return v.validate.Struct(out)
}

func main() {
	app := fiber.New(fiber.Config{
		StructValidator: &structValidator{validate: validator.New()},
	})

	db, err := models.DbConnection()
	if err != nil {
		panic("Failed to connect to database")
	}

	repositories.Configure(db)

	// db.AutoMigrate(&models.User{})

	middlewares.Configure(app)
	controllers.Configure(app)

	app.Get("/api/system/info", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"name":    APP_NAME,
			"version": VERSION,
			"now":     time.Now(),
		})
	})

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
