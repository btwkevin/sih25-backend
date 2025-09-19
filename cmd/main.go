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
		c.Set("Access-Control-Allow-Origin", "sih25-frontend-psi.vercel.app")
		c.Set("Access-Control-Allow-Headers", "Content-Type")
		c.Set("Access-Control-Allow-Credentials", "true")
		c.Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")

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
	app.Get("/home", handlers.JWTMiddleware, handlers.Home)
	app.Get("/logout", handlers.LogOut)

	app.Listen(":8080")
}
