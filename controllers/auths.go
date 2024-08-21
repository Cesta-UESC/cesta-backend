package controllers

import (
	"github.com/Cesta-UESC/cesta-backend/api"
	"github.com/Cesta-UESC/cesta-backend/configuration"
	"github.com/Cesta-UESC/cesta-backend/service"
	"github.com/gofiber/fiber/v3"
)

func ConfigureAuths(api fiber.Router) {
	auths := api.Group("/auths")

	auths.Post("/login", Login)
}

func Login(c fiber.Ctx) error {
	data := new(api.LoginData)

	if err := c.Bind().Body(data); err != nil {
		return err
	}

	response, err := service.Login(data)

	if err != nil {
		return c.Status(400).JSON(api.ApiException{
			Code:    api.ErrorCodeUserCredentialInvalid,
			Details: "Usuário ou senha inválidos",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:  configuration.Auth.AuthorizarionCookieLabel,
		Value: response.Token,
		Path:  "/api",
		// Domain: ,
		Expires: response.Expiration,
		// Secure: true,
		HTTPOnly: true,
	})

	return c.Status(201).JSON(response)
}
