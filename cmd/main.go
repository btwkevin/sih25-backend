package main

import (
	"github.com/btwkevin/sih25-backend/handlers"
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

	app.Post("/signup", handlers.Signup)

	app.Listen(":8080")
}
