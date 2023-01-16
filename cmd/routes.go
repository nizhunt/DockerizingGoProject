package main

import (
	"tester-go-docker/handlers"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App){

	app.Get("/", handlers.HomeHandler)
	app.Get("/facts", handlers.GetAllFacts)
	app.Post("/facts", handlers.CreateFact)
	app.Get("/fact/:id", handlers.GetFact)
	app.Delete("/fact/:id", handlers.DeleteFact)
}