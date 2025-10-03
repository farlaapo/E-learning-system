package model

import (
	"time"

	"github.com/gofrs/uuid"

	
)

// Organization represents a school or company on the platform (MVP version)
type Organization struct {
	ID          uuid.UUID   `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description,omitempty"`
	OwnerID     uuid.UUID   `json:"owner_id"`        // main admin of the org
	Tutors      []uuid.UUID `json:"tutors,omitempty"` // list of tutor IDs under this org
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	DeletedAt   *time.Time    `json:"deleted_at,omitempty"`  
}
