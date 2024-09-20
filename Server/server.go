package server

import (
	chat "b-care/Chat"
	controllers "b-care/Controllers"
	helpers "b-care/Helpers"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func Server() {
	app := fiber.New()

	port := helpers.GetEnv("PORT")

	// Inicializacion.

	if port == "" {
		port = "5000"
	}

	//CORS

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000","https://www.b-care.live/"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))
	//Recibe topic en form value
	app.Post("/Recommendation", controllers.Recommendation)
	app.Get("/GetUser", controllers.GetUser)
	// recibe prompt en form value
	app.Post("/ChatBot", chat.ChatBot)

	log.Fatal(app.Listen(":" + port))

}
