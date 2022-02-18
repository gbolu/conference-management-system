package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/gbolu/conference-management-system/app/models"
	editService "github.com/gbolu/conference-management-system/app/services/edits"
	talkServices "github.com/gbolu/conference-management-system/app/services/talks"
	response "github.com/gbolu/conference-management-system/pkg/utils/responseHandlers"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetAllTalks(ctx *fiber.Ctx) error {
	uniqueId, uuidError:= uuid.Parse(ctx.Params("conference_id"))

	if uuidError != nil {
		panic(fmt.Sprint(uuidError.Error()))
	}
	talks, err := talkServices.GetAllTalksByConferenceID(uniqueId)

	if(err != nil && errors.Is(err, gorm.ErrRecordNotFound)) {
		return response.SendErrorResponse(ctx, fiber.StatusNotFound, "Talks with that conference ID not found.", []error{errors.New(err.Error())})
	}

	return response.SendSuccessResponse(ctx, fiber.StatusOK, "Talks retrieved successfully.", fiber.Map{"talks": talks})
}

func CreateTalks(ctx *fiber.Ctx) error {
	uniqueId, uuidError:= uuid.Parse(ctx.Params("conference_id"))

	if uuidError != nil {
		panic(fmt.Sprint(uuidError.Error()))
	}

	t:= &models.Talk{}

	if err := ctx.BodyParser(t); err != nil {
		fmt.Println(err)
		return response.SendErrorResponse(ctx, fiber.StatusBadRequest, "Invalid request body.", []error{errors.New(err.Error())})
	}

	t.ID = uuid.New()
	t.Conference_ID = uniqueId
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()

	newTalk:= talkServices.CreateTalk(t)

	return response.SendSuccessResponse(ctx, fiber.StatusCreated, "Talk created successfully.", fiber.Map{"talk": newTalk})
}

func EditTalk(ctx *fiber.Ctx) error {
	uniqueId, uuidError:= uuid.Parse(ctx.Params("talk_id"))

	if uuidError != nil {
		panic(fmt.Sprint(uuidError.Error()))
	}

	t, err := talkServices.FindTalkById(uniqueId)

	if(err != nil) {
		if (errors.Is(err, gorm.ErrRecordNotFound)) {
			return response.SendErrorResponse(ctx, fiber.StatusNotFound, "Talk not found.", []error{errors.New(err.Error())})
		}

		return response.SendErrorResponse(ctx, fiber.StatusInternalServerError, "Unable to find talk. Please try again.", []error{errors.New(err.Error())})
	}

	e := models.Edit{}
	json_previous_state, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err)
		return response.SendErrorResponse(ctx, fiber.StatusBadRequest, "JSON issue.", []error{errors.New(err.Error())})
	}

	e.PreviousState = json_previous_state

	if err := ctx.BodyParser(t); err != nil {
		fmt.Println(err)
		return response.SendErrorResponse(ctx, fiber.StatusBadRequest, "Invalid request body.", []error{errors.New(err.Error())})
	}

	e.ID = uuid.New()
	e.CreatedAt = time.Now()
	e.UpdatedAt = time.Now()

	json_current_state, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err)
		return response.SendErrorResponse(ctx, fiber.StatusBadRequest, "JSON issue.", []error{errors.New(err.Error())})
	}
	e.CurrentState = json_current_state
	e.EditTargetID = t.ID
	e.EditType = "talk"

	editService.CreateEdit(&e)

	talkServices.PersistTalk(t)

	return response.SendSuccessResponse(ctx, fiber.StatusOK, "Conference updated successfully.", fiber.Map{"conference": t})
}