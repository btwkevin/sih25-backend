package handlers

import (
	"github.com/btwkevin/sih25-backend/database"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Signup(c *fiber.Ctx) error {
	var user User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := database.AddUser(user.Email, user.Password); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Database Add user error",
		})
	}

	defer database.Close()

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "signup successful",
	})
}
