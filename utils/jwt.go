package utils

import (
	"os"
	"time"

	"github.com/HizkiaHalim/Cinema-Ticketing/repository"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Admin  bool   `json:"admin"`
	jwt.RegisteredClaims
}

func GenerateJWT(id uint, email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	user, err := repository.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	claims := &Claims{
		UserID: user.ID,
		Email:  user.Email,
		Admin:  user.Admin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	return claims, nil
}
