package controllers

import (
	editService "github.com/gbolu/conference-management-system/app/services/edits"
	response "github.com/gbolu/conference-management-system/pkg/utils/responseHandlers"
	"github.com/gofiber/fiber/v2"
)

func GetAllEdits(ctx *fiber.Ctx) error {
	edits:= editService.GetAllEdits()

	return response.SendSuccessResponse(ctx, fiber.StatusOK, "Edits retrieved successfully.", fiber.Map{"edits": edits})
}
