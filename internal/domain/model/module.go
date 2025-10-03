package model

import (
	"time"

	"github.com/gofrs/uuid"
)

type Module struct {
	ID       uuid.UUID    `json:"id"`
	CourseID uuid.UUID    `json:"course_id"  binding:"required" `
	Title    string       `json:"title"`
	Order    int           `json:"order"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`            

}