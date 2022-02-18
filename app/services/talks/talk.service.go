package talkServices

import (
	"github.com/gbolu/conference-management-system/app/models"
	database_utils "github.com/gbolu/conference-management-system/platform"
	"github.com/google/uuid"
)

func GetAllTalksByConferenceID(uniqueId uuid.UUID) (Talks []models.Talk, errors error) {
	c := []models.Talk{}
	db := database_utils.GetDatabase()

	err := db.Where("conference_id = ?", uniqueId).Find(&c).Error

	return c, err
}

func CreateTalk(Talk *models.Talk) (c *models.Talk) {
	db := database_utils.GetDatabase()
	db.Create(&Talk)

	return Talk
}

func FindTalkById(uniqueId uuid.UUID) (Talk *models.Talk, errors error) {
	db := database_utils.GetDatabase()
	t := models.Talk{}

	err := db.First(&t, uniqueId).Error

	return &t, err
}

func PersistTalk(Talk *models.Talk) (c *models.Talk, errors error) {
	db := database_utils.GetDatabase()

	err := db.Save(&Talk).Error

	return Talk, err
}