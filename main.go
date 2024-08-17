package main

import (
	"log"
	"time"

	"github.com/Cesta-UESC/cesta-backend/controller"
	"github.com/Cesta-UESC/cesta-backend/model"
	"github.com/Cesta-UESC/cesta-backend/repository"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/helmet"
	"github.com/gofiber/fiber/v3/middleware/limiter"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/requestid"
)

const VERSION = "0.0.1"
const APP_NAME = "cesta-app"

func configure_middleware(app *fiber.App) {
	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] | (${request_id}) ${ip} ${latency} | ${method} ${path} ${status} ${error}\n",
		CustomTags: map[string]logger.LogFunc{
			"request_id": func(output logger.Buffer, c fiber.Ctx, data *logger.Data, extraParam string) (int, error) {
				return output.WriteString(requestid.FromContext(c))
			},
		},
	}))
	app.Use(cors.New())
	app.Use(helmet.New())
	app.Use(limiter.New(limiter.Config{
		Max:        100,
		Expiration: 60 * time.Second,
		LimitReached: func(c fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"message": "Too many requests",
			})
		},
	}))

	app.Get("/api/system/info", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"name":    APP_NAME,
			"version": VERSION,
			"now":     time.Now(),
		})
	})
}

func main() {
	app := fiber.New()

	db, err := model.DbConnection()
	if err != nil {
		panic("Failed to connect to database")
	}

	repository.SetDefault(db)

	// db.AutoMigrate(&models.User{})

	configure_middleware(app)
	controller.SetDefault(app)

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
