package service

import (
	"e-learning-system/internal/domain/model"
	"e-learning-system/internal/domain/repository"
	"fmt"
	"log"
	"time"

	"github.com/gofrs/uuid"
)

// ModuleService interface
type ModuleService interface {
	CreateModule(courseId uuid.UUID, title string, order int) (*model.Module, error)
	UpdateModule(module *model.Module) error
	DeletModule(moduleID uuid.UUID) error
	GetModuleById(moduleID uuid.UUID) (*model.Module, error)
	GetAllModule() ([]*model.Module, error)
}

// moduleService  struct
type moduleService struct {
	repo  repository.ModuleRepository
	token repository.TokenRepository
}

// CreateEnrollment implements ModuleService.
func (s *moduleService) CreateModule(courseId uuid.UUID, title string, order int) (*model.Module, error) {
	neoModule, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	now := time.Now()
// create new  module
	newModule := &model.Module{
		ID: neoModule,
		CourseID: courseId,
	  Title: title,
		Order: order,
		CreatedAt: now,
		UpdatedAt: now,
	}

	
    // log the new enrollment
  log.Printf("module created %+v", newModule)

	// save to repository
    err = s.repo.Create(newModule)
    if err != nil {
        return nil, fmt.Errorf("failed to create enrollment: %w", err)
    }

    return newModule, nil
	

}

// DeletEnrollment implements ModuleService.
func (s *moduleService) DeletModule(moduleID uuid.UUID) error {
	_, err := s.repo.GetByID(moduleID)
	if err != nil {
		return  fmt.Errorf("could not find module with ID %s: %v", moduleID, err)
	}

	if err := s.repo.Delete(moduleID); err != nil {
		return  fmt.Errorf("failed to delete module with ID %s: %v", moduleID, err)
	}

	log.Printf("successfully deleted module with id %v", moduleID)
  return  nil
}

// GetAllEnrollment implements ModuleService.
func (s *moduleService) GetAllModule() ([]*model.Module, error) {
	module, err := s.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get  all module  ID  %v", module)
	}

	return  module, nil
}

// GetEnrollmentById implements ModuleService.
func (s *moduleService) GetModuleById(moduleID uuid.UUID) (*model.Module, error) {
	module, err := s.repo.GetByID(moduleID)
	if err != nil {
		return nil,  fmt.Errorf("could not find module with ID %s: %v", moduleID, err)
	}

	return  module, nil

}

// UpdateEnrollment implements ModuleService.
func (s *moduleService) UpdateModule(module *model.Module) error {
	_, err := s.repo.GetByID(module.ID)
	if err != nil {
		return fmt.Errorf("failed to find module with ID %s", module.ID)
	}

	if err := s.repo.Update(module); err != nil {
		return fmt.Errorf("failed to module with ID %s: %v", module.ID, err)
	}

	return nil
}

func NewModuleService(moduleRepo repository.ModuleRepository, tokenRepo repository.TokenRepository) ModuleService {
	return &moduleService{
		repo:  moduleRepo,
		token: tokenRepo,
	}
}
