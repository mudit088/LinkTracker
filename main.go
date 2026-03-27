package main

import (
	"link-tracker/config"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"link-tracker/routes"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    config.ConnectDB()

    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("API is running 🚀")
    })

    routes.AuthRoutes(app)

    app.Listen(":" + os.Getenv("PORT"))
}