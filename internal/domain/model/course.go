package model


import (
	
	"github.com/gofrs/uuid"

	
)

type Course struct {
	ID           uuid.UUID
	Title        string
	Description  string
	InstructorID uuid.UUID
}