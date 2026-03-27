package middleware

import (
    "strings"

    "github.com/gofiber/fiber/v2"
    "github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("secret_key") // same as before

func Protected() fiber.Handler {
    return func(c *fiber.Ctx) error {
        authHeader := c.Get("Authorization")

        if authHeader == "" {
            return c.Status(401).JSON(fiber.Map{
                "error": "Missing Authorization header",
            })
        }

        tokenString := strings.Split(authHeader, " ")
        if len(tokenString) != 2 {
            return c.Status(401).JSON(fiber.Map{
                "error": "Invalid token format",
            })
        }

        token, err := jwt.Parse(tokenString[1], func(token *jwt.Token) (interface{}, error) {
            return jwtSecret, nil
        })

        if err != nil || !token.Valid {
            return c.Status(401).JSON(fiber.Map{
                "error": "Invalid or expired token",
            })
        }

        claims := token.Claims.(jwt.MapClaims)

        c.Locals("user_id", int(claims["user_id"].(float64)))

        return c.Next()
    }
}