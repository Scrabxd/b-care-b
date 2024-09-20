package chat

import (
	helpers "b-care/Helpers"

	"github.com/gofiber/fiber/v3"
)

func ChatBot(c fiber.Ctx) error {

	var data map[string]interface{}
	prompt := c.FormValue("prompt")

	if helpers.CheckRegexp(prompt) && c.FormValue("Legit") != "" {
		data = map[string]interface{}{
			"Message": "¡Tu cita ha sido creada!",
			"Status":  200,
		}
		return c.Status(200).JSON(data)
	}

	response := helpers.AnalyzeText(prompt)
	data = map[string]interface{}{
		"Message": response,
		"Status":  200,
	}

	return c.Status(200).JSON(data)
}
