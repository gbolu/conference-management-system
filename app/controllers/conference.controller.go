package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/gbolu/conference-management-system/app/models"
	conferenceServices "github.com/gbolu/conference-management-system/app/services/conferences"
	editService "github.com/gbolu/conference-management-system/app/services/edits"
	response "github.com/gbolu/conference-management-system/pkg/utils/responseHandlers"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetAllConferences(ctx *fiber.Ctx) error {
	conferences:= conferenceServices.GetAllConferences()

	return response.SendSuccessResponse(ctx, fiber.StatusOK, "Conferences retrieved successfully.", fiber.Map{"conferences": conferences})
}

func CreateConference(ctx *fiber.Ctx) error {
	c:= &models.Conference{}

	if err := ctx.BodyParser(c); err != nil {
		fmt.Println(err)
		return response.SendErrorResponse(ctx, fiber.StatusBadRequest, "Invalid request body.", []error{errors.New(err.Error())})
	}

	c.ID = uuid.New()
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
	newConference:= conferenceServices.CreateConference(c)

	return response.SendSuccessResponse(ctx, fiber.StatusCreated, "Conference created successfully.", fiber.Map{"conference": newConference})
}

func EditConference(ctx *fiber.Ctx) error {
	uniqueId, uuidError:= uuid.Parse(ctx.Params("id"))

	if uuidError != nil {
		return response.SendErrorResponse(ctx, fiber.StatusBadRequest, "Invalid UUID.", []error{errors.New(uuidError.Error())})
	}

	c, err := conferenceServices.FindConferenceById(uniqueId)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return response.SendErrorResponse(ctx, fiber.StatusNotFound, "Conference with that uuid not found.", []error{errors.New(err.Error())})
	}

	e := models.Edit{}
	json_previous_state, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
		return response.SendErrorResponse(ctx, fiber.StatusBadRequest, "JSON issue.", []error{errors.New(err.Error())})
	}
	e.PreviousState = json_previous_state

	if err := ctx.BodyParser(c); err != nil {
		fmt.Println(err)
		return response.SendErrorResponse(ctx, fiber.StatusBadRequest, "Invalid request body.", []error{errors.New(err.Error())})
	}

	e.ID = uuid.New()
	e.CreatedAt = time.Now()
	e.UpdatedAt = time.Now()

	json_current_state, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
		return response.SendErrorResponse(ctx, fiber.StatusBadRequest, "JSON issue.", []error{errors.New(err.Error())})
	}
	e.CurrentState = json_current_state
	e.EditTargetID = c.ID
	e.EditType = "conference"

	editService.CreateEdit(&e)

	updatedConference, err := conferenceServices.PersistConference(c)

	if(err != nil) {
		return response.SendErrorResponse(ctx, fiber.StatusInternalServerError, "Sorry. Unable to save conference.", []error{errors.New(err.Error())})
	}

	return response.SendSuccessResponse(ctx, fiber.StatusOK, "Conference updated successfully.", fiber.Map{"conference": updatedConference})
}