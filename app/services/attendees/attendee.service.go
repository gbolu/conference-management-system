package attendeeService

import (
	"github.com/gbolu/conference-management-system/app/models"
	database_utils "github.com/gbolu/conference-management-system/platform"
	"github.com/google/uuid"
)

func CreateAttendee(attendee *models.Attendee) (c *models.Attendee, errors error) {
	db := database_utils.GetDatabase()
	err := db.Create(&attendee).Error;

	return attendee, err
}

func FindAllAttendeesByTalkId(talkId uuid.UUID) (md *[]models.Attendee, errors error) {
	a := []models.Attendee{}
	db := database_utils.GetDatabase()
	err := db.Where("talk_id = ?", talkId).Find(&a).Error;

	return &a, err
}

func DeleteAttendeeByTalkId(uniqueId uuid.UUID, talkId uuid.UUID) (attendee *models.Attendee, errors error) {
	db := database_utils.GetDatabase()
	a := models.Attendee{}

	err := db.Where("talk_id = ? AND id = ?", talkId, uniqueId).Delete(&a).Error;

	return &a, err
}
