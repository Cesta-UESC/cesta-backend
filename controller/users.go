package controller

import (
	"github.com/Cesta-UESC/cesta-backend/repository"
	"github.com/gofiber/fiber/v3"
)

func SetUsersDefault(api fiber.Router) {
	users := api.Group("/users")

	users.Get("/", GetUsers)
}

func GetUsers(c fiber.Ctx) error {
	user, err := repository.Q.Usuarios.Where(repository.Usuarios.UsuarioNome.Like("%a%")).Order(repository.Usuarios.UsuarioNome).Limit(10).Find()
	if err != nil {
		return c.SendStatus(400)
	}

	return c.JSON(fiber.Map{
		"usuarios": user,
	})
}
