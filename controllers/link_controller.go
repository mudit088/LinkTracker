package controllers

import (
	"fmt"
	"github.com/mudit088/LinkTracker/config"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type LinkInput struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

// Protected
func AddLink(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(int)

	var input LinkInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	// Get profile_id from user_id
	var profileID int
	err := config.DB.QueryRow(
		"SELECT id FROM profiles WHERE user_id=$1",
		userID,
	).Scan(&profileID)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Profile not found",
		})
	}

	query := `
        INSERT INTO links (profile_id, title, url)
        VALUES ($1, $2, $3)
    `

	_, err = config.DB.Exec(query, profileID, input.Title, input.URL)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Link added successfully",
	})
}

func GetLinksByUsername(c *fiber.Ctx) error {
	username := c.Params("username")

	query := `
    SELECT l.id, l.title, l.url
    FROM links l
    JOIN profiles p ON l.profile_id = p.id
    WHERE p.username = $1
    ORDER BY l.id DESC
`

	rows, err := config.DB.Query(query, username)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch links",
		})
	}
	defer rows.Close()

	var links []fiber.Map

	for rows.Next() {
		var id int
		var title, url string

		rows.Scan(&id, &title, &url)

		links = append(links, fiber.Map{
			"id":        id,
			"title":     title,
			"url":       url,
			"short_url": "/r/" + fmt.Sprint(id),
		})
	}

	return c.JSON(links)
}

func RedirectLink(c *fiber.Ctx) error {
	linkID := c.Params("id")

	var url string

	// get original URL
	err := config.DB.QueryRow(
		"SELECT url FROM links WHERE id=$1",
		linkID,
	).Scan(&url)

	if err != nil {
		return c.Status(404).SendString("Link not found")
	}

	// log click event
	_, err = config.DB.Exec(
		"INSERT INTO click_events (link_id, user_agent, ip) VALUES ($1, $2, $3)",
		linkID,
		c.Get("User-Agent"),
		c.IP(),
	)

	if err != nil {
		return c.Status(500).SendString("Error logging click")
	}
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}
	// redirect to actual URL
	return c.Redirect(url, 302)
}
