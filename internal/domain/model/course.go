package model

import (
	"time"

	"github.com/gofrs/uuid"
)

type Course struct {
	ID           uuid.UUID  `json:"id,omitempty"`
	Title        string     `json:"title" binding:"required"`
	Description  string     `json:"description,omitempty"`
	InstructorID uuid.UUID  `json:"instructor_id" binding:"required"`
  Category     string      `json:"category"`
  Tags         []string    `json:"tags"`
	CreatedAt    time.Time  `json:"created_at,omitempty"`
	UpdatedAt    time.Time  `json:"updated_at,omitempty"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
}




 