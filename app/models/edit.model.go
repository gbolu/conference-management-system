package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Edit struct {
	ID 	uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	PreviousState []byte `db:"previousState"`
	CurrentState []byte `db:"currentState"`
	EditType string `db:"edit_type" json:"edit_type"`
	EditTargetID uuid.UUID `db:"edit_target_id" json:"edit_target_id"`
}

func (b Edit) Value() (driver.Value, error) {
    return json.Marshal(b)
}

func (b *Edit) Scan(value interface{}) error {
    j, ok := value.([]byte)
    if !ok {
        return errors.New("type assertion to []byte failed")
    }

    return json.Unmarshal(j, &b)
}
