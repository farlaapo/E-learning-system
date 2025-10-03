package service

import (
	"e-learning-system/internal/domain/model"
	"e-learning-system/internal/domain/repository"
	"fmt"
	"log"
	"time"

	"github.com/gofrs/uuid"
)

// EnrollmentService interface
type EnrollmentService interface {
	CreateEnrollment(courseId, userId uuid.UUID, CertificateTemplate string, completed bool) (*model.Enrollment, error)
	UpdateEnrollment(enrollment *model.Enrollment) error
	DeletEnrollment(enrollmentID uuid.UUID) error
	GetEnrollmentById(enrollmentID uuid.UUID) (*model.Enrollment, error)
	GetAllEnrollment() ([]*model.Enrollment, error)
}

// enrollmentService struct
type enrollmentService struct {
	repo  repository.EnrolledRepository
	token repository.TokenRepository
}

// CreateEnrollment implements EnrollmentService.
func (s *enrollmentService) CreateEnrollment(courseId uuid.UUID, userId uuid.UUID, certificateTemplate string, completed bool) (*model.Enrollment, error) {

    // generate UUID
    neoEnrollment, err := uuid.NewV4()
    if err != nil {
        return nil, err
    }

    now := time.Now()

    // set CertificateIssuedAt only if completed
    var certIssuedAt *time.Time
    if completed {
        certIssuedAt = &now
    }

    // create enrollment
    amEnrollment := &model.Enrollment{
        ID:                  neoEnrollment,
        CourseID:            courseId,
        UserID:              userId,
        EnrollmentAt:        now,           // enrollment timestamp
        Completed:           completed,
        CertificateIssuedAt: certIssuedAt, // nil if not completed
        CertificateTemplate: &certificateTemplate,
        CreatedAt:           now,
        
    }

    // log the new enrollment
    log.Printf("enrollment created %+v", amEnrollment)

    // save to repository
    err = s.repo.Create(amEnrollment)
    if err != nil {
        return nil, fmt.Errorf("failed to create enrollment: %w", err)
    }

    return amEnrollment, nil
}



// DeletEnrollment implements EnrollmentService.
func (s *enrollmentService) DeletEnrollment(enrollmentID uuid.UUID) error {
	_, err := s.repo.GetByID(enrollmentID)
	if err != nil {
		return  fmt.Errorf("could not find enrollment with ID %s: %v", enrollmentID, err)
	}

	if err := s.repo.Delete(enrollmentID); err != nil {
		return  fmt.Errorf("failed to delete enrollment with ID %s: %v", enrollmentID, err)
	}

	log.Printf("successfully deleted enrollment with id %v", enrollmentID)
  return  nil
}

// GetAllEnrollment implements EnrollmentService.
func (s *enrollmentService) GetAllEnrollment() ([]*model.Enrollment, error) {
	enrollment, err := s.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get  all enrollment  ID  %v", enrollment)
	}

	return  enrollment, nil

}

// GetEnrollmentById implements EnrollmentService.
func (s *enrollmentService) GetEnrollmentById(enrollmentID uuid.UUID) (*model.Enrollment, error) {
	enrollment, err := s.repo.GetByID(enrollmentID)
	if err != nil {
		return nil,  fmt.Errorf("could not find enrollment with ID %s: %v", enrollmentID, err)
	}

	return  enrollment, nil

}

// UpdateEnrollment implements EnrollmentService.
func (s *enrollmentService) UpdateEnrollment(enrollment *model.Enrollment) error {
	_, err := s.repo.GetByID(enrollment.ID)
	if err != nil {
		return fmt.Errorf("failed to find enrollment with ID %s", enrollment.ID)
	}

	if err := s.repo.Update(enrollment); err != nil {
		return fmt.Errorf("failed to enrollment with ID %s: %v", enrollment.ID, err)
	}

	return nil
}

func NewEnrollmentService(enrollmentRepo repository.EnrolledRepository, tokenRepo repository.TokenRepository) EnrollmentService {
	return &enrollmentService{
		repo:  enrollmentRepo,
		token: tokenRepo,
	}
}
