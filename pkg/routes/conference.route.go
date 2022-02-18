package routes

import (
	"github.com/gbolu/conference-management-system/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func ConferenceRoutes(a *fiber.App) {
	route := a.Group("/api")

	route.Get("/conferences", controllers.GetAllConferences)
	route.Post("/conferences", controllers.CreateConference)
	route.Patch("/conferences/:id", controllers.EditConference)
}