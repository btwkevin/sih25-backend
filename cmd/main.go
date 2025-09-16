package main

import (
	"github.com/gofiber/fiber/v2"
)

// Todo:= add handler dir
func main() {
	app := fiber.New()
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "healthy",
		})
	})
}
