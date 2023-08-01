package main

import (
	"github.com/gcharalla/url-shortener/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {

	app.Get("/r/:redirect", handlers.Redirect)

	app.Get("/goly", handlers.GetAllGolies)
	app.Get("/goly/:id", handlers.GetGoly)
	app.Post("/goly", handlers.CreateGoly)
	app.Patch("/goly", handlers.UpdateGoly)
	app.Delete("/goly/:id", handlers.DeleteGoly)

}
