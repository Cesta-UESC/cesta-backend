package controller

import (
	"github.com/gofiber/fiber/v3"
)

func SetDefault(app fiber.Router) {
	api := app.Group("/api")

	SetUsersDefault(api)
}
