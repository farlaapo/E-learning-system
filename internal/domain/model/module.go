package model

import (
	
	"github.com/gofrs/uuid"

	
)

type Module struct {
	ID       uuid.UUID
	CourseID uuid.UUID
	Title    string
	Order    int
}