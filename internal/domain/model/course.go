package model

import (
	"time"

	"github.com/gofrs/uuid"
)

type Course struct {
	ID           uuid.UUID
	Title        string
	Description  string
	InstructorID uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}