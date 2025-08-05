package routes

import (
	"crud/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	
	api.Post("/contacts", controllers.CreateContact)
	api.Get("/contacts", controllers.GetContacts)
	api.Get("/contacts/:id", controllers.GetContactByID)
	api.Put("/contacts/:id", controllers.UpdateContact)
	api.Delete("/contacts/:id", controllers.DeleteContact)
}
