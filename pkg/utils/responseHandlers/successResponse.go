package response

import "github.com/gofiber/fiber/v2"

func SendSuccessResponse(ctx *fiber.Ctx, responseCode int, message string, data fiber.Map) error {
	return ctx.JSON(fiber.Map{
		"code": responseCode,
		"data": data,
		"message": message,
		"status": true,
	})
}