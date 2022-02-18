package routes

import (
	"errors"

	response "github.com/gbolu/conference-management-system/pkg/utils/responseHandlers"
	"github.com/gofiber/fiber/v2"
)

// 404 Error route
func NotFoundRoute(a *fiber.App) {
    a.Use(
        func(ctx *fiber.Ctx) error {
            // Return HTTP 404 status and JSON response.
            return response.SendErrorResponse(ctx, fiber.StatusNotFound, "Sorry, that route does not exist.", []error{errors.New("invalid route")})
        },
    )
}