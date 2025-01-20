package jwt

import (
	"errors"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

var secretKey []byte

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	secret := os.Getenv("SECRET_KEY")
	if secret == "" {
		panic("SECRET_KEY not set in .env file")
	}
	secretKey = []byte(secret)
}

type Claims struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

func ParseToken(tokenString string) (int64, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return 0, errors.New("invalid token format")
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return 0, errors.New("token expired or not active yet")
			} else {
				return 0, errors.New("could not handle token. did you tried to hack us?")
			}
		}
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.UserID, nil
	}
	return 0, errors.New("invalid token")
}
