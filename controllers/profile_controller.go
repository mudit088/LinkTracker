package controllers

import (
    "link-tracker/config"
    "github.com/gofiber/fiber/v2"
)

type ProfileInput struct {
    Username string `json:"username"`
    Bio      string `json:"bio"`
}

func CreateProfile(c *fiber.Ctx) error {
    userID := c.Locals("user_id").(int)

    var input ProfileInput

    if err := c.BodyParser(&input); err != nil {
        return c.Status(400).JSON(fiber.Map{
            "error": "Invalid input",
        })
    }

    query := `
        INSERT INTO profiles (user_id, username, bio)
        VALUES ($1, $2, $3)
    `

    _, err := config.DB.Exec(query, userID, input.Username, input.Bio)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    return c.JSON(fiber.Map{
        "message": "Profile created successfully",
    })
}


func GetProfile(c *fiber.Ctx) error {
    username := c.Params("username")

    var id int
    var bio string

    query := `
        SELECT id, bio FROM profiles WHERE username=$1
    `

    err := config.DB.QueryRow(query, username).Scan(&id, &bio)
    if err != nil {
        return c.Status(404).JSON(fiber.Map{
            "error": "Profile not found",
        })
    }

    return c.JSON(fiber.Map{
        "username": username,
        "bio":      bio,
    })
}