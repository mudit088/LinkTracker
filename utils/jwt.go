package utils

import (
    "time"

    "github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("secret_key") // later move to .env

func GenerateToken(userID int) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 72).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}