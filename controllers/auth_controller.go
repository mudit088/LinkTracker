package controllers

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/mudit088/LinkTracker/config"
	"github.com/mudit088/LinkTracker/utils"
	"golang.org/x/crypto/bcrypt"
)

type SignupInput struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}
    
func Signup(c *fiber.Ctx) error {
    var input SignupInput

    if err := c.BodyParser(&input); err != nil {
        return c.Status(400).JSON(fiber.Map{
            "error": "Invalid input",
        })
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 14)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "error": "Could not hash password",
        })
    }


    var userID int
    query := `
        INSERT INTO users (email, password)
        VALUES ($1, $2)
        RETURNING id
    `

    err = config.DB.QueryRow(query, input.Email, string(hashedPassword)).Scan(&userID)

    if err != nil {
        if err == sql.ErrNoRows {
            return c.Status(400).JSON(fiber.Map{
                "error": "User not created",
            })
        }

        return c.Status(500).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    return c.JSON(fiber.Map{
        "message": "User created successfully",
        "user_id": userID,
    })
}



type LoginInput struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {
    var input LoginInput

    if err := c.BodyParser(&input); err != nil {
        return c.Status(400).JSON(fiber.Map{
            "error": "Invalid input",
        })
    }

    var userID int
    var hashedPassword string

    query := `SELECT id, password FROM users WHERE email=$1`

    err := config.DB.QueryRow(query, input.Email).Scan(&userID, &hashedPassword)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{
            "error": "Invalid email or password",
        })
    }

    err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(input.Password))
    if err != nil {
        return c.Status(400).JSON(fiber.Map{
            "error": "Invalid email or password",
        })
    }

    token, err := utils.GenerateToken(userID)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "error": "Could not generate token",
        })
    }

    return c.JSON(fiber.Map{
        "message": "Login successful",
        "token":   token,
    })
}