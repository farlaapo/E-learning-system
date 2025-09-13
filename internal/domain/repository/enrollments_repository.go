package repository

import (
	"e-learning-system/internal/domain/model"

	"github.com/gofrs/uuid"
)

// enrollments interface required method
type EnrolledRepository interface{
  Create(enrollment *model.Enrollment) error
	Update(enrollment *model.Enrollment) error
	Delete(enrollmentID  uuid.UUID) error
	GetByID(enrollmentID uuid.UUID) (*model.Enrollment, error)
	GetAll()([]*model.Enrollment, error)
} 