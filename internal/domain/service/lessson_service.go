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
type LessonService interface {
	CreateLesson(moduleID uuid.UUID, title string, content string, videoUrl []string,  order int) (*model.Lesson, error)
	UpdateLesson(lesson *model.Lesson) error
	DeletLesson(lessonID uuid.UUID) error
	GetLessonById(lessonID uuid.UUID) (*model.Lesson, error)
	GetAllLesson() ([]*model.Lesson, error)
}

// lessonService struct
type lessonService struct {
	repo  repository.LessonRepository
	token repository.TokenRepository
}

// CreateLesson implements LessonService.
func (s *lessonService) CreateLesson(moduleID uuid.UUID, title string, content string, videoUrl []string, order int) (*model.Lesson, error) {
	neoLesson , err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	now := time.Now()

	newLesson := &model.Lesson{
		ID: neoLesson,
		ModuleID: moduleID,
		Title: title,
		Content: content,
		VideoURL: videoUrl,
		Order: order,
		CreatedAt: now,
	}
	// log the new enrollment
    log.Printf("lesson created %+v", neoLesson)

    // save to repository
    err = s.repo.Create(newLesson)
    if err != nil {
        return nil, fmt.Errorf("failed to create lesson: %w", err)
    }

    return newLesson, nil


}

// DeletLesson implements LessonService.
func (s *lessonService) DeletLesson(lessonID uuid.UUID) error {
	_, err := s.repo.GetByID(lessonID)
	if err != nil {
		return  fmt.Errorf("could not find lesson with ID %s: %v", lessonID, err)
	}

	if err := s.repo.Delete(lessonID); err != nil {
		return  fmt.Errorf("failed to delete lesson with ID %s: %v", lessonID, err)
	}

	log.Printf("successfully deleted lesson with id %v", lessonID)
  return  nil
}

// GetAllLesson implements LessonService.
func (s *lessonService) GetAllLesson() ([]*model.Lesson, error) {
	lesson, err := s.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get  all lesson  ID  %v", lesson)
	}

	return  lesson, nil
}

// GetLessonById implements LessonService.
func (s *lessonService) GetLessonById(lessonID uuid.UUID) (*model.Lesson, error) {
	lesson, err := s.repo.GetByID(lessonID)
	if err != nil {
		return nil,  fmt.Errorf("could not find lesson with ID %s: %v", lessonID, err)
	}

	return  lesson, nil
}

// UpdateLesson implements LessonService.
func (s *lessonService) UpdateLesson(lesson *model.Lesson) error {
	_, err := s.repo.GetByID(lesson.ID)
	if err != nil {
		return fmt.Errorf("failed to find lesson with ID %s", lesson.ID)
	}

	if err := s.repo.Update(lesson); err != nil {
		return fmt.Errorf("failed to lesson with ID %s: %v", lesson.ID, err)
	}

	return nil
}

// lessonService instance
func NewlessonService(lessonRepo repository.LessonRepository, tokenRepo repository.TokenRepository) LessonService {
	return &lessonService{
		repo:  lessonRepo,
		token: tokenRepo,
	}
}
