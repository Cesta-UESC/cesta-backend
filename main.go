package main

import (
	"log"

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

	api := app.Group("/api") // /api

	users := api.Group("/users")

	users.Get("/", func(c fiber.Ctx) error {
		user, err := repository.Q.Usuarios.Where(repository.Usuarios.UsuarioNome.Like("%a%")).Order(repository.Usuarios.UsuarioNome).Limit(10).Find()
		if err != nil {
			return c.SendStatus(400)
		}
		// return c.SendString(fmt.Sprintf("id %d name %s email %s", user.UsuarioID, user.UsuarioNome, user.UsuarioEmail))
		return c.JSON(user)
	})

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
