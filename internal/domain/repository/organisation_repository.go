package repository

import (
	"e-learning-system/internal/domain/model"

	"github.com/gofrs/uuid"
)

// OrganizationBillingRepository interface with required methods
type OrganizationgRepository interface {
	Create(Organization *model.Organization) error
	Update(Organization *model.Organization) error
	Delete(OrganizationID uuid.UUID) error
	GetByID(OrganizationID uuid.UUID) (*model.Organization, error)
	GetAll() ([]*model.Organization, error)
}
