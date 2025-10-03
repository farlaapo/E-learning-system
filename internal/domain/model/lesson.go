package model

import (
	"time"

	"github.com/gofrs/uuid"
)


type Lesson struct {
	ID       uuid.UUID    `json:"id"`
	ModuleID uuid.UUID     `json:"module_id"  binding:"required" `
	Title    string        `json:"title"`
	Content  string        `json:"content"`   
	VideoURL []string       `json:"video_url"`
	Order    int           `json:"order"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`            

}
