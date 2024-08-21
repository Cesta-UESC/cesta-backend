package controllers

import (
	"github.com/Cesta-UESC/cesta-backend/middlewares"
	"github.com/Cesta-UESC/cesta-backend/repositories"
	"github.com/gofiber/fiber/v3"
)

func ConfigureUsers(api fiber.Router) {
	users := api.Group("/users")

	users.Get("/", GetUsers, middlewares.Authorize())
}

func GetUsers(c fiber.Ctx) error {
	user, err := repositories.Q.Usuarios.Where(repositories.Usuarios.UsuarioNome.Like("%a%")).Order(repositories.Usuarios.UsuarioNome).Limit(10).Find()
	if err != nil {
		return c.SendStatus(400)
	}

	return c.JSON(fiber.Map{
		"usuarios": user,
	})
}
