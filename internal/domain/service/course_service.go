package service

import (
	"e-learning-system/internal/domain/model"
	"e-learning-system/internal/domain/repository"

	"fmt"
	"log"
	"time"

	"github.com/gofrs/uuid"
)

type CourseService interface {
	CreateCourse(title, description string, instructoRID uuid.UUID, tags []string, category string) (*model.Course, error)
	UpdateCourse(Course *model.Course) error
	DeleteCourse(courseID uuid.UUID) error
	GetCourseById(courseID uuid.UUID) (*model.Course, error)
	ListAllCourse() ([]*model.Course, error)
}

type courseService struct {
	repo  repository.CourseRepository
	token repository.TokenRepository
}

// CreateCourse implements CourseService.
func (s *courseService) CreateCourse(title string, description string, instructoRID uuid.UUID, tags []string,  category string) (*model.Course, error) {
	// Generate a new UUID for the course ID
	neoCourse, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	// Create a new course instance
	amCourse := &model.Course{
		ID:           neoCourse,
		Title:        title,
		Description:  description,
		InstructorID: instructoRID,
		Category: category,
		Tags: tags,
		CreatedAt:    time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: &time.Time{},
	}
	// log the new course creation attemptt
	log.Printf("creating course: %+v", amCourse)

	// save the new course to the repository
	err = s.repo.Create(amCourse)
	if err != nil {
		return nil, fmt.Errorf("failed to create course %v", err)
	}
	return amCourse, nil
}

// DeleteCourse implements CourseService.
func (s *courseService) DeleteCourse(courseID uuid.UUID) error {
	_, err := s.repo.GetByID(courseID)
	if err != nil {
		return fmt.Errorf("could not find cours with ID %s: %v", courseID, err)
	}

	if err := s.repo.Delete(courseID); err != nil {
		return fmt.Errorf("failed to delete course With ID %s:", courseID)
	}

	log.Printf("successfully deleted course with ID %s", courseID)

	return nil

}

// GetCourseById implements CourseService.
func (s *courseService) GetCourseById(courseID uuid.UUID) (*model.Course, error) {
	course, err := s.repo.GetByID(courseID)
	if err != nil {
		return nil, fmt.Errorf("failed to find course withID %v", err)
	}

	return course, nil
}

// ListAllCourse implements CourseService.
func (s *courseService) ListAllCourse() ([]*model.Course, error) {
	course, err := s.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get all course with IDs %v", course)
	}

	return course, nil
}

// UpdateCourse implements CourseService.
func (s *courseService) UpdateCourse(Course *model.Course) error {
	_, err := s.repo.GetByID(Course.ID)
	if err != nil {
		return fmt.Errorf("failed to find course with ID %s", Course.ID)
	}

	if err := s.repo.Update(Course); err != nil {
		return fmt.Errorf("failed to course with ID %s: %v", Course.ID, err)
	}

	return nil
}

func NewCourseService(courseRepo repository.CourseRepository, tokenRepo repository.TokenRepository) CourseService {
	return &courseService{
		repo:  courseRepo,
		token: tokenRepo,
	}

}
