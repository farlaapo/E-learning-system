package service

import (
	"E-Learning-System/internal/domain/model"
	"E-Learning-System/internal/domain/repository"
	"fmt"
	"log"
	"time"

	"github.com/gofrs/uuid"
)

// =============================
// AdminService Interface
// =============================
type AdminService interface {
	// 1 Dashboard
	GetDashboardSummary() (*model.AdminDashboard, error)

	// 2 Managed Entities
	CreateManagedEntity(name, entityType, status string) (*model.ManagedEntity, error)
	UpdateManagedEntity(entity *model.ManagedEntity) error
	DeleteManagedEntity(id uuid.UUID) error
	ListAllManagedEntities() ([]model.ManagedEntity, error)

	// 3 Approval Requests
	CreateApprovalRequest(entityType string, entityID uuid.UUID) (*model.ApprovalRequest, error)
	UpdateApprovalStatus(id uuid.UUID, status string, reviewedBy uuid.UUID) error
	ListAllApprovalRequests() ([]model.ApprovalRequest, error)
	ListPendingApprovals() ([]model.ApprovalRequest, error)

	// 4 System Settings
	UpsertSystemSettings(paymentGateway, theme string) (*model.SystemSettings, error)
	GetLatestSystemSettings() (*model.SystemSettings, error)
}

// =============================
// Admin Service Struct
// =============================
type adminService struct {
	repo  repository.AdminRepository
	token repository.TokenRepository
}

// =============================
// Dashboard
// =============================
func (s *adminService) GetDashboardSummary() (*model.AdminDashboard, error) {
	dashboard, err := s.repo.GetDashboardSummary()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch admin dashboard summary: %v", err)
	}
	log.Println("Dashboard summary successfully retrieved.")
	return dashboard, nil
}

// =============================
// Managed Entities
// =============================
func (s *adminService) CreateManagedEntity(name, entityType, status string) (*model.ManagedEntity, error) {
	newID, err := uuid.NewV4()
	if err != nil {
		return nil, fmt.Errorf("failed to generate UUID: %v", err)
	}

	entity := &model.ManagedEntity{
		ID:        newID,
		Name:      name,
		Type:      entityType,
		Status:    status,
		CreatedAt: time.Now(),
	}

	log.Printf("Creating managed entity: %+v", entity)

	if err := s.repo.CreateManagedEntity(entity); err != nil {
		return nil, fmt.Errorf("failed to create managed entity: %v", err)
	}

	return entity, nil
}

func (s *adminService) UpdateManagedEntity(entity *model.ManagedEntity) error {
	if entity == nil {
		return fmt.Errorf("cannot update nil managed entity")
	}

	log.Printf("Updating managed entity: %+v", entity)

	if err := s.repo.UpdateManagedEntity(entity); err != nil {
		return fmt.Errorf("failed to update managed entity with ID %s: %v", entity.ID, err)
	}
	return nil
}

func (s *adminService) DeleteManagedEntity(id uuid.UUID) error {
	if err := s.repo.DeleteManagedEntity(id); err != nil {
		return fmt.Errorf("failed to delete managed entity with ID %s: %v", id, err)
	}
	log.Printf("Managed entity deleted successfully with ID %s", id)
	return nil
}

func (s *adminService) ListAllManagedEntities() ([]model.ManagedEntity, error) {
	entities, err := s.repo.GetAllManagedEntities()
	if err != nil {
		return nil, fmt.Errorf("failed to list managed entities: %v", err)
	}
	return entities, nil
}

// =============================
// Approval Requests
// =============================
func (s *adminService) CreateApprovalRequest(entityType string, entityID uuid.UUID) (*model.ApprovalRequest, error) {
	newID, err := uuid.NewV4()
	if err != nil {
		return nil, fmt.Errorf("failed to generate UUID: %v", err)
	}

	req := &model.ApprovalRequest{
		ID:          newID,
		EntityType:  entityType,
		EntityID:    entityID,
		RequestDate: time.Now(),
		Status:      "pending",
	}

	log.Printf("Creating approval request: %+v", req)

	if err := s.repo.CreateApprovalRequest(req); err != nil {
		return nil, fmt.Errorf("failed to create approval request: %v", err)
	}

	return req, nil
}

func (s *adminService) UpdateApprovalStatus(id uuid.UUID, status string, reviewedBy uuid.UUID) error {
	if err := s.repo.UpdateApprovalStatus(id, status, reviewedBy); err != nil {
		return fmt.Errorf("failed to update approval status for request ID %s: %v", id, err)
	}
	log.Printf("Approval status updated successfully for request ID %s", id)
	return nil
}

func (s *adminService) ListAllApprovalRequests() ([]model.ApprovalRequest, error) {
	requests, err := s.repo.GetAllApprovalRequests()
	if err != nil {
		return nil, fmt.Errorf("failed to list approval requests: %v", err)
	}
	return requests, nil
}

func (s *adminService) ListPendingApprovals() ([]model.ApprovalRequest, error) {
	pending, err := s.repo.GetPendingApprovals()
	if err != nil {
		return nil, fmt.Errorf("failed to list pending approvals: %v", err)
	}
	return pending, nil
}

// =============================
// System Settings
// =============================
func (s *adminService) UpsertSystemSettings(paymentGateway, theme string) (*model.SystemSettings, error) {
	newID, err := uuid.NewV4()
	if err != nil {
		return nil, fmt.Errorf("failed to generate UUID: %v", err)
	}

	setting := &model.SystemSettings{
		ID:             newID,
		PaymentGateway: paymentGateway,
		Theme:          theme,
		//EmailTemplateID: emailTemplateID,
		UpdatedAt: time.Now(),
	}

	log.Printf("Upserting system settings: %+v", setting)

	if err := s.repo.UpsertSystemSettings(setting); err != nil {
		return nil, fmt.Errorf("failed to upsert system settings: %v", err)
	}

	return setting, nil
}

func (s *adminService) GetLatestSystemSettings() (*model.SystemSettings, error) {
	setting, err := s.repo.GetLatestSystemSettings()
	if err != nil {
		return nil, fmt.Errorf("failed to get latest system settings: %v", err)
	}
	return setting, nil
}

// =============================
// Constructor
// =============================
func NewAdminService(adminRepo repository.AdminRepository, tokenRepo repository.TokenRepository) AdminService {
	return &adminService{
		repo:  adminRepo,
		token: tokenRepo,
	}
}
