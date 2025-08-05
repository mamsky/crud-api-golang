package test

import (
	"crud/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateContactMock(c *fiber.Ctx) error {
	type Contact struct {
		Name   string `json:"name"`
		Email  string `json:"email"`
		Phone  string `json:"phone"`
		Gender string `json:"gender"`
	}

	var contact Contact
	if err := c.BodyParser(&contact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid body",
		})
	}

	if !utils.IsValidEmail(contact.Email) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid email",
		})
	}

	// Return sukses langsung tanpa database
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Contact created",
		"data":    contact,
	})
}

