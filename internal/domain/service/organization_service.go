package service

import (
	"e-learning-system/internal/domain/model"
	"e-learning-system/internal/domain/repository"
	"fmt"
	"log"
	"time"

	"github.com/gofrs/uuid"
)

// OrgsService interface
type OrganizationService interface {
	CreateOrganization(name string, description string, ownerID uuid.UUID, tutors []uuid.UUID) (*model.Organization, error)
	UpdateOrganization(orgs *model.Organization) error
	DeleteOrganization(orgsID uuid.UUID) error
	GetOrganizationById(orgsID uuid.UUID) (*model.Organization, error)
	GetAllOrganization() ([]*model.Organization, error)
}

// organizationService  struct
type organizationService struct {
	repo  repository.OrganizationgRepository
	token repository.TokenRepository
}

// CreateOrganization implements OrganizationService.
func (s *organizationService) CreateOrganization(name string, description string, ownerID uuid.UUID, tutors []uuid.UUID) (*model.Organization, error) {
	// generete uuid
	neoOrgs , err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	now := time.Now()
 /// create new orgs
	newOrgs := &model.Organization{
		ID: neoOrgs,
		Name: name,
		Description: description,
		OwnerID: ownerID,
		Tutors: tutors,
		CreatedAt: now,
	}
	// log the new enrollment
    log.Printf("organization created %+v", newOrgs)

    // save to repository
    err = s.repo.Create(newOrgs)
    if err != nil {
        return nil, fmt.Errorf("failed to create organization: %w", err)
    }

    return newOrgs, nil


}

// DeleteOrganization implements OrganizationService.
func (s *organizationService) DeleteOrganization(orgsID uuid.UUID) error {
	_, err := s.repo.GetByID(orgsID)
	if err != nil {
		return  fmt.Errorf("could not find organization with ID %s: %v", orgsID, err)
	}

	if err := s.repo.Delete(orgsID); err != nil {
		return  fmt.Errorf("failed to delete organization with ID %s: %v", orgsID, err)
	}

	log.Printf("successfully deleted organization with id %v", orgsID)
  return  nil
}

// GetAllOrganization implements OrganizationService.
func (s *organizationService) GetAllOrganization() ([]*model.Organization, error) {
	orgs, err := s.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get  all organization  ID  %v", orgs)
	}

	return  orgs, nil
}

// GetOrganizationById implements OrganizationService.
func (s *organizationService) GetOrganizationById(orgsID uuid.UUID) (*model.Organization, error) {
	orgs, err := s.repo.GetByID(orgsID)
	if err != nil {
		return nil,  fmt.Errorf("could not find organiztion with ID %s: %v", orgsID, err)
	}

	return  orgs, nil
}

// UpdateOrganization implements OrganizationService.
func (s *organizationService) UpdateOrganization(orgs *model.Organization) error {
	_, err := s.repo.GetByID(orgs.ID)
	if err != nil {
		return fmt.Errorf("failed to find organization with ID %s", orgs.ID)
	}

	if err := s.repo.Update(orgs); err != nil {
		return fmt.Errorf("failed to organization with ID %s: %v", orgs.ID, err)
	}

	return nil
}

func NewOrganizationService(orgsRepo repository.OrganizationgRepository, tokenRepo repository.TokenRepository) OrganizationService {
	return &organizationService{
		repo:  orgsRepo,
		token: tokenRepo,
	}
}
