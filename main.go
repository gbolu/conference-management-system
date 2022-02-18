package main

import (
	"os"

	"github.com/gbolu/conference-management-system/pkg/routes"
	"github.com/gbolu/conference-management-system/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.AttendeeRoutes(app)
	routes.ConferenceRoutes(app)
	routes.TalkRoutes(app)
	routes.EditRoutes(app)
	routes.NotFoundRoute(app)

	utils.StartServerWithGracefulShutdown(app)

	os.Exit(0)
}