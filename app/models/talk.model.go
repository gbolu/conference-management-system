package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Talk struct {
	ID 	uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	Title string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
	Scheduled_Date time.Time `db:"scheduled_date" json:"scheduled_date"`
	Duration int `db:"duration" json:"duration"`
	Conference_ID uuid.UUID `db:"conferenceId" json:"conferenceId"`
}

func (b Talk) Value() (driver.Value, error) {
    return json.Marshal(b)
}

func (b *Talk) Scan(value interface{}) error {
    j, ok := value.([]byte)
    if !ok {
        return errors.New("type assertion to []byte failed")
    }

    return json.Unmarshal(j, &b)
}
