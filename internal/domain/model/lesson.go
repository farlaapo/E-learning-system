package model



import (

	"github.com/gofrs/uuid"

	
)


type Lesson struct {
	ID       uuid.UUID
	ModuleID uuid.UUID
	Title    string
	Content  string
	VideoURL string
	Order    int
}
