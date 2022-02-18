package routes

import (
	"github.com/gbolu/conference-management-system/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func AttendeeRoutes(a *fiber.App) {
	route := a.Group("/api")

	route.Get("/conferences/:conference_id/talks/:talk_id/attendees", controllers.GetAllAttendeesByTalk)
	route.Post("/conferences/:conference_id/talks/:talk_id/attendees", controllers.AddAttendeeToTalk)
	route.Delete("/conferences/:conference_id/talks/:talk_id/attendees/:attendee_id", controllers.RemoveAttendeeFromTalk)
}