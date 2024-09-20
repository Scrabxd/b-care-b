package server

import (
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
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	app.Post("/recommendation", controllers.Recommendation)
	app.Get("/GetUser", controllers.GetUser)

	log.Fatal(app.Listen(":" + port))

}
