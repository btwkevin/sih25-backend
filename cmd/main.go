package main

import (
	"log"

	"github.com/btwkevin/sih25-backend/database"
	"github.com/btwkevin/sih25-backend/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	if err := database.ConnectDb(); err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Headers", "*")

		if c.Method() == fiber.MethodOptions {
			return c.SendStatus(fiber.StatusNoContent)
		}
		return c.Next()
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "healthy",
		})
	})

	app.Post("/signup", handlers.Signup)
	app.Post("/signin", handlers.Signin)

	app.Listen(":8080")
}
