package controllers

import (
	helpers "b-care/Helpers"
	"encoding/json"
	"os"

	"github.com/gofiber/fiber/v3"
)

type User struct {
	Name            string `json:"name"`
	Last_name       string `json:"last_name"`
	Cancer_type     string `json:"cancer_type"`
	Cancer_stage    string `json:"cancer_stage"`
	SSN             string `json:"ssn"`
	Country         string `json:"country"`
	State           string `json:"state"`
	Doctor_name     string `json:"doctor_name"`
	Doctor_lastname string `json:"doctor_lastname"`
}

type Response struct {
	User User `json:"user"`
}

func Recommendation(c fiber.Ctx) error {
	var response map[string]interface{}

	topic := c.FormValue("topic")

	if topic == "" {
		response = map[string]interface{}{
			"Message": "No topic provided.",
			"Status":  400,
		}
		return c.Status(400).JSON(response)
	}

	// USO DEL APi
	return nil

}

func GetUser(c fiber.Ctx) error {

	data, err := os.ReadFile("data.json")
	if err != nil {
		return helpers.ErrorCreate(c, "No se puido leer el archivo.", 500, err)
	}

	var response Response

	err = json.Unmarshal(data, &response)
	if err != nil {
		return helpers.ErrorCreate(c, "No se pudo formatear los datos.", 500, err)
	}

	return c.Status(200).JSON(response)
}
