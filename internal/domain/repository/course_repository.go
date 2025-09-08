package repository

import (
	"e-learning-system/internal/domain/model"

	"github.com/gofrs/uuid"
)

// Course interface with required methods

type CourseRepository interface {
	Create(Course *model.Course) error
	Update(Course  *model.Course) error
	Delete(CourseID uuid.UUID) error
	GetByID(CourseID uuid.UUID) ( *model.Course, error)
	GetAll() ([]*model.Course, error)
	FindInstructor(InstructorID string) (*model.Course, error)
}


