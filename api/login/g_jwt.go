package api_login

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

var JwtKey = []byte(os.Getenv("API_SECRET"))

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func createJWTToken(userinfo string) (string, error) {
	// JWT needs expiration time, claims
	expire := time.Now().Add(time.Hour)
	claims := &Claims{
		Username: userinfo,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expire),
		},
	}
	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tks, err := tk.SignedString(JwtKey)
	if err != nil {
		return "", errors.New("token generation error")
	}
	return tks, nil
}
