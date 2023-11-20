package helpers

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/vamsikartik01/Authentication_service/api/models"
)

func GenerateJwt(accountId int, username string) (string, error) {
	claims := models.JwtClaims{
		UserId:   accountId,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(Config.Jwt.Secret))
	if err != nil {
		log.Println("Error Signing JWT Token", err)
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (*models.JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Config.Jwt.Secret), nil
	})

	if err != nil {
		log.Println("Error Parsing JWT token", err)
		return nil, err
	}

	if claims, ok := token.Claims.(*models.JwtClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("Invalid token")
}
