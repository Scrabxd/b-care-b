package controllers

import (
	helpers "b-care/Helpers"
	"encoding/json"
	"os"

	"math/rand"

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

	recommendations := []string{"Recuerde hacer ejercicio regularmente.", "Debe mantener un peso saludable", "Deje de fumar, ¡Usted puede lograrlo!", "Se detecto una cantidad inusual de Radon en el air en Durango, porfavor mantenga sus precauciones,"}

	randomIndex := rand.Intn(len(recommendations) + 1)

	randomRecommendation := recommendations[randomIndex]

	// USO DEL APi
	// client := api.ConnectApi()

	// respons, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
	// 	Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
	// 		openai.UserMessage("Say hello scrabby."),
	// 	}),
	// 	Model: openai.F(openai.ChatModelGPT3_5Turbo),
	// })
	// if err != nil {
	// 	response = map[string]interface{}{
	// 		"Message": "Error consiguiendo una recomendación.",
	// 		"Status":  500,
	// 		"Error":   err.Error(),
	// 	}
	// 	return c.Status(500).JSON(response)
	// }
	response = map[string]interface{}{
		"Message": randomRecommendation,
		"Status":  200,
	}

	return c.Status(200).JSON(response)
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
