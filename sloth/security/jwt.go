package security

import (
	"github.com/dgrijalva/jwt-go"
	"sloth/infra"
	"time"
)

type JwtPair struct {
	Token        string
	RefreshToken string
}

type JwtManager struct {
	Config infra.SecurityConfig
}

func (j *JwtManager) createToken(subject string, lifetime int) (string, error) {
	now := time.Now()
	expiresAt := now.Add(time.Duration(lifetime))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		IssuedAt:  now.Unix(),
		ExpiresAt: expiresAt.Unix(),
		Subject:   subject,
	})

	signed, err := token.SignedString(j.Config.JwtSigningSecret)
	if err != nil {
		return "", err
	}

	return signed, nil
}

func (j *JwtManager) CreateToken(subject string) (string, error) {
	return j.createToken(subject, j.Config.JwtTokenLifetime)
}

func (j *JwtManager) CreateRefreshToken(subject string) (string, error) {
	return j.createToken(subject, j.Config.JwtRefreshTokenLifetime)
}
