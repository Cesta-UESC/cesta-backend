package middlewares

import (
	"strings"

	"github.com/Cesta-UESC/cesta-backend/configuration"
	"github.com/Cesta-UESC/cesta-backend/service"
	"github.com/gofiber/fiber/v3"
)

func Authorize() fiber.Handler {
	return func(c fiber.Ctx) (err error) {
		token := getTokenFromCookie(c)

		if len(token) == 0 {
			token = getTokenFromHeader(c)
		}

		if len(token) == 0 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.ErrUnauthorized)
		}

		t, err := service.Validate(token)

		println("aaa", t.Valid)
		if err != nil || !t.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.ErrUnauthorized)
		}

		err = service.UseClaims(c, t.Claims)

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.ErrUnauthorized)
		}

		return c.Next()
	}
}

func getTokenFromHeader(c fiber.Ctx) string {
	authorization, ok := c.GetReqHeaders()["Authorization"]

	if !ok || len(authorization) < 1 || len(authorization[0]) == 0 {
		return ""
	}

	valueParts := strings.Split(authorization[0], "Bearer ")

	if len(valueParts) != 2 {
		return ""
	}

	return valueParts[1]
}

func getTokenFromCookie(c fiber.Ctx) string {
	value := c.Cookies(configuration.Auth.AuthorizarionCookieLabel)

	if len(value) == 0 {
		return value
	}

	return value
}
