package repository

import (
	"e-learning-system/internal/domain/model"

	"github.com/gofrs/uuid"
)



type ModuleRepository interface{
  Create(module *model.Module) error
	Update(module *model.Module) error
	Delete(moduleID  uuid.UUID) error
	GetByID(moduIDle uuid.UUID) (*model.Module, error)
	GetAll()([]*model.Module, error)
} 