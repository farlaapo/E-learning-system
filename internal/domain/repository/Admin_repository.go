package repository

import (
	"e-learning-system/internal/domain/model"

	"github.com/gofrs/uuid"
)

// OrganizationRepository interface with required methods
// type AdminRepository interface {
// 	Create(Admin *model.Admin) error
// 	Update(Admin *model.Admin) error
// 	Delete(AdminID uuid.UUID) error
// 	GetByID(AdminID uuid.UUID) (*model.Admin, error)
// 	GetAll() ([]*model.Admin, error)
// }

type AdminRepository interface {
	// 1 Dashboard
	GetDashboardSummary() (*model.AdminDashboard, error)

	// 2 Managed Entities
	CreateManagedEntity(entity *model.ManagedEntity) error
	UpdateManagedEntity(entity *model.ManagedEntity) error
	DeleteManagedEntity(entityID uuid.UUID) error
	GetAllManagedEntities() ([]*model.ManagedEntity, error)

	// 3 Approval Requests
	CreateApprovalRequest(req *model.ApprovalRequest) error
	UpdateApprovalStatus(id uuid.UUID, status string, reviewedBy uuid.UUID) error
	GetPendingApprovals() ([]*model.ApprovalRequest, error)
	GetAllApprovalRequests() ([]*model.ApprovalRequest, error)

	// 4 System Settings
	UpsertSystemSettings(setting *model.SystemSettings) error
	GetLatestSystemSettings() (*model.SystemSettings, error)
}
