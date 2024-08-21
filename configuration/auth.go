package configuration

import (
	"strconv"

	"github.com/Cesta-UESC/cesta-backend/utils"
)

const AUTH_PREFIX = "AUTH_"

type AuthConfiguration struct {
	Secret                   []byte
	ExpirationSeconds        int
	AuthorizarionCookieLabel string
}

func SetAuthDefault(auth *AuthConfiguration) {
	auth.Secret = []byte(utils.GetEnvOrDefault(AUTH_PREFIX+"SECRET", "secret"))
	auth.ExpirationSeconds, _ = strconv.Atoi(utils.GetEnvOrDefault(AUTH_PREFIX+"ExpirationSeconds", "1800"))
	auth.AuthorizarionCookieLabel = utils.GetEnvOrDefault(AUTH_PREFIX+"AuthorizarionCookieLabel", "authorization")
}
