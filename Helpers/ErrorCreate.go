package helpers

import "github.com/gofiber/fiber/v3"

func ErrorCreate(c fiber.Ctx, message string, status int64, Errror error) error {
	return c.Status(400).JSON(map[string]interface{}{
		"Message": message,
		"Status":  status,
		"Error":   Errror,
	})
}
