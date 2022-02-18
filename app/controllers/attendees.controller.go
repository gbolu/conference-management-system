package controllers

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gbolu/conference-management-system/app/models"
	attendeeService "github.com/gbolu/conference-management-system/app/services/attendees"
	talkServices "github.com/gbolu/conference-management-system/app/services/talks"
	response "github.com/gbolu/conference-management-system/pkg/utils/responseHandlers"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func AddAttendeeToTalk(ctx *fiber.Ctx) error {
	uniqueId, uuidError:= uuid.Parse(ctx.Params("talk_id"))

	if uuidError != nil {
		return response.SendErrorResponse(ctx, fiber.StatusBadRequest, "Invalid UUID.", []error{errors.New(uuidError.Error())})
	}

	a:= &models.Attendee{}

	if err := ctx.BodyParser(a); err != nil {
		fmt.Println(err)
		return response.SendErrorResponse(ctx, fiber.StatusBadRequest, "Invalid request body.", []error{errors.New(err.Error())})
	}

	a.ID = uuid.New()
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()
	a.Talk_id = uniqueId

	newAttendee, err := attendeeService.CreateAttendee(a)

	if err != nil {
		if (strings.Contains(err.Error(), "duplicate")) {
			return response.SendErrorResponse(ctx, fiber.StatusBadRequest, "Attendee with that username/email already exists.", []error{errors.New(err.Error())})
		}
		return response.SendErrorResponse(ctx, fiber.StatusInternalServerError, "Sorry. Unable to create attendee.", []error{errors.New(err.Error())})
	}

	return response.SendSuccessResponse(ctx, fiber.StatusCreated, "Attendee added successfully.", fiber.Map{"attendee": newAttendee})
}

func GetAllAttendeesByTalk(ctx *fiber.Ctx) error {
	uniqueId, uuidError:= uuid.Parse(ctx.Params("talk_id"))

	if uuidError != nil {
		return response.SendErrorResponse(ctx, fiber.StatusBadRequest, "Invalid UUID.", []error{errors.New(uuidError.Error())})
	}

	talk, err := talkServices.FindTalkById(uniqueId)

	if (err != nil) {
		if (errors.Is(err, gorm.ErrRecordNotFound)) {
			return response.SendErrorResponse(ctx, fiber.StatusNotFound, "Talk with that uuid not found.", []error{errors.New(err.Error())})
		}

		return response.SendErrorResponse(ctx, fiber.StatusInternalServerError, "Unable to find talk. Please try", []error{errors.New(err.Error())})
	}

	a, err := attendeeService.FindAllAttendeesByTalkId(talk.ID)

	if err != nil {
		return response.SendErrorResponse(ctx, fiber.StatusInternalServerError, "Sorry. Unable to find attendees.", []error{errors.New(err.Error())})
	}

	return response.SendSuccessResponse(ctx, fiber.StatusOK, "Attendees retrieved successfully.", fiber.Map{"attendees": a})
}

func RemoveAttendeeFromTalk(ctx *fiber.Ctx) error {
	uniqueId, uuidError:= uuid.Parse(ctx.Params("talk_id"))
	attendeeId, attendeeuidError:= uuid.Parse(ctx.Params("attendee_id"))

	if uuidError != nil {
		return response.SendErrorResponse(ctx, fiber.StatusBadRequest, "Invalid UUID.", []error{errors.New(uuidError.Error())})
	}

	if attendeeuidError != nil {
		return response.SendErrorResponse(ctx, fiber.StatusBadRequest, "Invalid UUID.", []error{errors.New(uuidError.Error())})
	}

	t, err := attendeeService.DeleteAttendeeByTalkId(uniqueId, attendeeId)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.SendErrorResponse(ctx, fiber.StatusNotFound, "Attendee with that uuid not found.", []error{errors.New(err.Error())})
		}
		return response.SendErrorResponse(ctx, fiber.StatusInternalServerError, "Sorry. Unable to delete attendee.", []error{errors.New(err.Error())})
	}

	if err := ctx.BodyParser(t); err != nil {
		fmt.Println(err)
		return response.SendErrorResponse(ctx, fiber.StatusBadRequest, "Invalid request body.", []error{errors.New(err.Error())})
	}

	return response.SendSuccessResponse(ctx, fiber.StatusOK, "Attendee removed successfully.", fiber.Map{})
}