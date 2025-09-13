package model

import (
	
	"github.com/gofrs/uuid"
 "time"
	
)


type Enrollment struct {
	ID                  uuid.UUID
	CourseID            uuid.UUID
	UserID              uuid.UUID
	EnrolledAt          time.Time
	Completed           bool
	CertificateIssuedAt *time.Time
	CertificateTemplate string
	Created_at  time.Time
	Updated_at   time.Time
	Deleted_at   *time.Time
}

