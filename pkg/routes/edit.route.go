package routes

import (
	"github.com/gbolu/conference-management-system/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func EditRoutes(a *fiber.App) {
	route := a.Group("/api")

	route.Get("/edits", controllers.GetAllEdits)
	route.Post("/conferences", controllers.CreateConference)
}