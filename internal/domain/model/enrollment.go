package model

import (
	
	"github.com/gofrs/uuid"
 "time"
	
)

type Enrollment struct {
    ID                  uuid.UUID   `json:"id"`
    CourseID            uuid.UUID   `json:"course_id"  binding:"required" `
    UserID              uuid.UUID   `json:"user_id"  binding:"required"`
    EnrollmentAt        time.Time   `json:"enrollment_at"`
    Completed           bool        `json:"completed"`
    CertificateIssuedAt *time.Time  `json:"certificate_issued_at,omitempty"` // nullable
    CertificateTemplate *string     `json:"certificate_template,omitempty"`  // nullable
    CreatedAt           time.Time   `json:"created_at"`
    UpdatedAt           time.Time   `json:"updated_at"`
    DeletedAt           *time.Time  `json:"deleted_at,omitempty"`            // nullable
}











