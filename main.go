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

	// Load .env only in local (ignore error in production)
	_ = godotenv.Load()

	config.ConnectDB()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("API is running 🚀")
	})

	routes.AuthRoutes(app)
	//routes.LinkRoutes(app) // make sure you added this

	// Railway PORT fix
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal(app.Listen(":" + port))
}