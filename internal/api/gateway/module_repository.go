package gateway

import (
	"database/sql"
	"e-learning-system/internal/domain/model"
	"e-learning-system/internal/domain/repository"
	"github.com/gofrs/uuid"
	"log"
	"fmt"
)

// ModuleRepositoryImpl struct
type ModuleRepositoryImpl struct {
	db *sql.DB
}

// Create implements repository.ModuleRepository.
func (r *ModuleRepositoryImpl) Create(module *model.Module) error {
	_, err := r.db.Exec(`CALL create_module($1, $2, $3, $4)`,
   module.ID,
	 module.CourseID,
   module.Title,
   module.Order)

	  if err != nil {
        log.Printf("Error calling create_module: %v", err)
        return err
    }

    log.Printf("module created %+v", module)
    return nil
}

// Delete implements repository.ModuleRepository.
func (r *ModuleRepositoryImpl) Delete(moduleID uuid.UUID) error {
	_, err := r.db.Exec(`CALL delete_module($1)` ,moduleID)
	if err != nil {
		log.Printf("Error calling update_module for ID %v, %v", err, moduleID)
		return err
	}

	log.Printf("module created %+v", moduleID)
	return nil
}

// GetAll implements repository.ModuleRepository.
func (r *ModuleRepositoryImpl) GetAll() ([]*model.Module, error) {
	query := `SELECT * FROM  get_all_module()`
	row, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer row.Close()
	var modules []*model.Module
	for row.Next() {
		var m   model.Module
		err := row.Scan(
			&m.ID,
			&m.CourseID,
			&m.Title,
			&m.Order,
			&m.CreatedAt,
      &m.UpdatedAt,
      &m.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		modules = append(modules, &m)
	} 

	return modules, nil
}

// GetByID implements repository.ModuleRepository.
func (r *ModuleRepositoryImpl) GetByID(moduleID uuid.UUID) (*model.Module, error) {
	var m   model.Module
	query := `SELECT * FROM get_enrollment_by_id($1)`
	row := r.db.QueryRow(query, moduleID)
	err := row.Scan( 
		&m.ID,
			&m.CourseID,
			&m.Title,
			&m.Order,
			&m.CreatedAt,
      &m.UpdatedAt,
      &m.DeletedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("course not found")
			return nil, fmt.Errorf("course not found")
		}
		log.Printf("DB error: %v", err)
		return nil, err
	}

	return   &m, nil

}

// Update implements repository.ModuleRepository.
func (r *ModuleRepositoryImpl) Update(module *model.Module) error {
	_, err := r.db.Exec(`CALL update_module($1, $2, $3 )`,
    module.ID,
	  module.Title,
	  module.Order)

		if err != nil {
		log.Printf("error calling update_module")
		return err
	}

	log.Printf("MODULE updated %+v ", module)
	return nil

}

// NewModuleRepositoryImpl
func NewModuleRepository(db *sql.DB) repository.ModuleRepository {
	return &ModuleRepositoryImpl{
		db: db,
	}
}
