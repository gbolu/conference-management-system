package routes

import (
	"github.com/gbolu/conference-management-system/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func TalkRoutes(a *fiber.App) {
	route := a.Group("/api")

	route.Get("/conferences/:conference_id/talks", controllers.GetAllTalks)
	route.Post("/conferences/:conference_id/talks", controllers.CreateTalks)
	route.Patch("/conferences/:conference_id/talks/:talk_id", controllers.EditTalk)
}