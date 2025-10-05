package repository

import (
	"e-learning-system/internal/domain/model"

	"github.com/gofrs/uuid"
)

// enrollments interface required method
type PaymnetRepository interface{
  Create(payment *model.Payment) error
	Update(payment *model.Payment) error
	Delete(paymentID  uuid.UUID) error
	GetByID(paymentID uuid.UUID) (*model.Payment, error)
	GetAll()([]*model.Payment, error)
} 