package main

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/joke", func(c *fiber.Ctx) error {
		var apiResponse map[string]string
		res, err := http.Get("https://api.chucknorris.io/jokes/random")
		if err != nil {
			c.Status(fiber.StatusInternalServerError).SendString("error getting API response")
		}

		json.NewDecoder(res.Body).Decode(&apiResponse)
		return c.JSON(apiResponse["value"])
	})

	app.Listen(":3000")
}
