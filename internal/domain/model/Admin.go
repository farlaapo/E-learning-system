package model

import (
	"time"

	"github.com/gofrs/uuid"
)

type AdminDashboard struct {
	TotalUsers    int64   `json:"total_users"`
	TotalCourses  int64   `json:"total_courses"`
	TotalRevenue  float64 `json:"total_revenue"`
	ActiveCourses int64   `json:"active_courses"`
}

// User/Organization/Course Management overview
type ManagedEntity struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"` // user, organization, course
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

// Tutor/Organization approval requests
type ApprovalRequest struct {
	ID          uuid.UUID  `json:"id"`
	EntityType  string     `json:"entity_type"` // tutor, organization
	EntityID    uuid.UUID  `json:"entity_id"`
	RequestDate time.Time  `json:"request_date"`
	Status      string     `json:"status"` // pending, approved, rejected
	ReviewedBy  *uuid.UUID `json:"reviewed_by,omitempty"`
	ReviewedAt  *time.Time `json:"reviewed_at,omitempty"`
}

// System-wide settings
type SystemSettings struct {
	ID             uuid.UUID `json:"id"`
	PaymentGateway string    `json:"payment_gateway"`
	Theme          string    `json:"theme"`
	//EmailTemplateID string    `json:"email_template_id"`
	UpdatedAt time.Time `json:"updated_at"`
}
