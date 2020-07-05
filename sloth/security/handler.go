package security

import (
	"net/http"
	"sloth/infra"
	"time"
)

func NewHandler(logger *infra.LoggerConfig, jwt *JwtManager) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		token := jwt.CreateToken()

		refreshTokenCookie := http.Cookie{
			HttpOnly: true,
			Name: "X-Refresh-Token",
			Expires: time.Unix(int64(jwt.Config.JwtRefreshTokenLifetime), 0),
		}
	}
}