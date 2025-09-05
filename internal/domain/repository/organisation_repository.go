package repository

import (
	"e-learning-system/internal/domain/model"

	"github.com/gofrs/uuid"
)

// OrganizationBillingRepository interface with required methods
type OrganizationBillingRepository interface {
	Create(OrganizationBilling *model.Organization) error
	Update(OrganizationBilling *model.Organization) error
	Delete(OrganizationBrandingID uuid.UUID) error
	GetByID(OrganizationBrandingID uuid.UUID) (*model.Organization, error)
	GetAll() ([]*model.Organization, error)
}
