package editService

import (
	"github.com/gbolu/conference-management-system/app/models"
	database_utils "github.com/gbolu/conference-management-system/platform"
)

// func GetAllConferences() (conferences []models.Conference) {
// 	c := []models.Conference{}
// 	db := database_utils.GetDatabase()

// 	db.Find(&c)

// 	return c
// }

func GetAllEdits() (edits *[]models.Edit) {
	e := []models.Edit{}

	db := database_utils.GetDatabase()

	db.Find(&e)

	return &e
}

func CreateEdit(edit *models.Edit) (e *models.Edit) {
	db := database_utils.GetDatabase()

	db.Create(&edit)

	return edit
}