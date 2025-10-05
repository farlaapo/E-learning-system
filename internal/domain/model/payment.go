package model

import (
	"time"

	"github.com/gofrs/uuid"
)



type Payment struct {
	ID             uuid.UUID  `json:"id"`
	UserID         uuid.UUID  `json:"user_id"`          // Student or tutor
	Role           string     `json:"role"`             // "student" or "tutor"
	Amount         float64    `json:"amount"`
	Currency       string     `json:"currency"`         // e.g. "USD", "KES"
	Method         string     `json:"method"`           // e.g. "Stripe", "PayPal", "MobileMoney"
	Type           string     `json:"type"`             // "ONE_TIME" or "SUBSCRIPTION"
	Status         string     `json:"status"`           // "PENDING", "SUCCESS", "FAILED", "REFUNDED"
	TransactionRef string     `json:"transaction_ref"`  // Provider transaction reference
	Description    string     `json:"description"`      // Optional note or description
	PlanName       string     `json:"plan_name"`        // For subscriptions
	StartDate      *time.Time `json:"start_date"`       // Subscription start date
	EndDate        *time.Time `json:"end_date"`         // Subscription end date
	RenewalDate    *time.Time `json:"renewal_date"`     // Next billing date
	CancelledAt    *time.Time `json:"cancelled_at"`     // If cancelled
	ProviderRef    string     `json:"provider_ref"`     // Provider-side reference or ID
	IsRecurring    bool       `json:"is_recurring"`     // True if subscription
	CreatedAt      time.Time  `json:"created_at"`       // Timestamp of record creation
	UpdatedAt      time.Time  `json:"updated_at"`       // Timestamp of last update
	DeletedAt      *time.Time
}
