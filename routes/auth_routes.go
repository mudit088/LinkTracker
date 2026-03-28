package routes

import (
	"github.com/mudit088/LinkTracker/controllers"
	"github.com/mudit088/LinkTracker/middleware"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
    app.Post("/signup", controllers.Signup)
	app.Post("/login", controllers.Login)

// 	app.Get("/protected", middleware.Protected(), func(c *fiber.Ctx) error {
//     userID := c.Locals("user_id")
//     return c.JSON(fiber.Map{
//         "message": "You are authorized",
//         "user_id": userID,
//     })
// })

    app.Post("/profile", middleware.Protected(), controllers.CreateProfile)
    app.Get("/profile/:username", controllers.GetProfile)

	app.Post("/links", middleware.Protected(), controllers.AddLink)
    app.Get("/links/:username", controllers.GetLinksByUsername)

	app.Get("/r/:id", controllers.RedirectLink)
}