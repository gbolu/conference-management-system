package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Attendee struct {
	ID 	uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	Username string `db:"username" json:"username"`
	Email string `db:"email" json:"email"`
	Attendee_Type string `db:"attendee_type" json:"attendee_type"`
	Talk_id uuid.UUID `db:"talk_id" json:"talk_id"`
}

func (b Attendee) Value() (driver.Value, error) {
    return json.Marshal(b)
}

func (b *Attendee) Scan(value interface{}) error {
    j, ok := value.([]byte)
    if !ok {
        return errors.New("type assertion to []byte failed")
    }

    return json.Unmarshal(j, &b)
}
