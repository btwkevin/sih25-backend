package handlers

import (
	"time"

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
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "user signup successful",
	})
}

func Signin(c *fiber.Ctx) error {
	var user User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := database.GetUserByEmail(user.Email, user.Password); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	token, err := GenerateJWT(user.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "jwt"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HTTPOnly = true
	cookie.Secure = true
	cookie.SameSite = fiber.CookieSameSiteNoneMode
	cookie.Path = "/"
	c.Cookie(cookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "user signin successful",
	})
}

func LogOut(c *fiber.Ctx) error {
	cookie := new(fiber.Cookie)
	cookie.Name = "jwt"
	cookie.Value = ""
	cookie.Expires = time.Now().Add(-time.Hour)
	cookie.HTTPOnly = true
	c.Cookie(cookie)

	return c.SendStatus(fiber.StatusNoContent)
}
