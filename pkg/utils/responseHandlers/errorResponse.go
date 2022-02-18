package response

import "github.com/gofiber/fiber/v2"

func SendErrorResponse(ctx *fiber.Ctx, responseCode int, message string, errors []error) error {
	return ctx.JSON(fiber.Map{
		"code": responseCode,
		"error": errors,
		"message": message,
		"status": false,
	})
}