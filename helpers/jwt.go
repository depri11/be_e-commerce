package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type claims struct {
	ID       uint
	Username string
	Role     int
	jwt.StandardClaims
}

func NewToken(id uint, username string, role int) *claims {
	return &claims{
		ID:       id,
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
}

func (c *claims) GenerateJWT() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(os.Getenv("JWT_KEY")))
}
