package service

import (
	"errors"
	"strconv"
	"time"

	"github.com/Cesta-UESC/cesta-backend/api"
	"github.com/Cesta-UESC/cesta-backend/configuration"
	"github.com/Cesta-UESC/cesta-backend/models"
	"github.com/Cesta-UESC/cesta-backend/repositories"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func Login(data *api.LoginData) (api.LoginResponse, error) {
	user, err := repositories.Q.Usuarios.Where(repositories.Usuarios.UsuarioNome.Eq(data.Username)).First()

	if err != nil || user.UsuarioID == 0 || user.UsuarioSenha != data.Password {
		return api.LoginResponse{}, errors.New("usuário ou senha inválidos")
	}

	// TODO: ENCRIPT PASSWORD

	return CreateToken(user)
}

func UseClaims(c fiber.Ctx, claims jwt.Claims) error {
	sub, err := claims.GetSubject()

	if err != nil {
		return err
	}

	c.Locals("UserId", sub)

	return nil
}

func Validate(token string) (*jwt.Token, error) {
	t, err := jwt.Parse(
		token,
		func(t *jwt.Token) (interface{}, error) { return configuration.Auth.Secret, nil },
		jwt.WithoutClaimsValidation(),
	)

	if err != nil {
		return t, err
	}

	expiration, expirationOk := t.Claims.GetExpirationTime()
	if expirationOk != nil || expiration.Before(time.Now()) {
		t.Valid = false
		return t, expirationOk
	}
	return t, err
}

func CreateToken(user *models.Usuarios) (api.LoginResponse, error) {
	expiration := time.Now().Add(time.Second * time.Duration(configuration.Auth.ExpirationSeconds))

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   strconv.Itoa(int(user.UsuarioID)),
		"exp":   expiration.Unix(),
		"admin": false,
		"roles": "",
	})

	token, err := t.SignedString(configuration.Auth.Secret)

	if err != nil {
		return api.LoginResponse{}, err
	}

	return api.LoginResponse{
		Token:      token,
		Type:       "Bearer",
		Expiration: expiration,
	}, nil
}
