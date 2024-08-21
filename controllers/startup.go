package controllers

import (
	"github.com/gofiber/fiber/v3"
)

func Configure(app fiber.Router) {
	api := app.Group("/api")

	ConfigureUsers(api)
	ConfigureAuths(api)
}
