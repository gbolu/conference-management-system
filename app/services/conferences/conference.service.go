package services

import (
	"github.com/gbolu/conference-management-system/app/models"
	database_utils "github.com/gbolu/conference-management-system/platform"
	"github.com/google/uuid"
)

func GetAllConferences() (conferences []models.Conference) {
	c := []models.Conference{}
	db := database_utils.GetDatabase()

	db.Find(&c)

	return c
}

func CreateConference(conference *models.Conference) (c *models.Conference) {
	db := database_utils.GetDatabase()
	db.Create(&conference)

	return conference
}

func FindConferenceById(uniqueId uuid.UUID) (conference *models.Conference, errors error) {
	db := database_utils.GetDatabase()
	c := models.Conference{}

	err := db.First(&c, uniqueId).Error


	return &c, err
}

func PersistConference(conference *models.Conference) (c *models.Conference, errors error) {
	db := database_utils.GetDatabase()

	err := db.Save(&conference).Error

	return conference, err
}