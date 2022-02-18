package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Conference struct {
	ID 	uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	Title string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
	Start_Date time.Time `db:"start_date" json:"start_date"`
	End_Date time.Time `db:"end_date" json:"end_date"`
}

func (b Conference) Value() (driver.Value, error) {
    return json.Marshal(b)
}

func (b *Conference) Scan(value interface{}) error {
    j, ok := value.([]byte)
    if !ok {
        return errors.New("type assertion to []byte failed")
    }

    return json.Unmarshal(j, &b)
}
