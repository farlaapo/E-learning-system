package gateway

import (
	"e-learning-system/internal/domain/model"
	"e-learning-system/internal/domain/repository"
	"database/sql"
	"fmt"
	"log"

	"github.com/gofrs/uuid"
)

type AdminRepositoryImpl struct {
	db *sql.DB
}

// =============================
// Dashboard
// =============================
func (r *AdminRepositoryImpl) GetDashboardSummary() (*model.AdminDashboard, error) {
	row := r.db.QueryRow(`SELECT * FROM get_admin_dashboard_summary()`)

	var dashboard model.AdminDashboard
	err := row.Scan(
		&dashboard.TotalUsers,
		&dashboard.TotalCourses,
		&dashboard.TotalRevenue,
		&dashboard.ActiveCourses,
	)
	if err != nil {
		log.Printf("Error fetching admin dashboard summary: %v", err)
		return nil, err
	}

	log.Println("Admin dashboard summary retrieved successfully.")
	return &dashboard, nil
}

// =============================
//  Managed Entities
// =============================

// Create Managed Entity
func (r *AdminRepositoryImpl) CreateManagedEntity(entity *model.ManagedEntity) error {
	_, err := r.db.Exec(`CALL create_managed_entity($1, $2, $3, $4)`,
		entity.ID, entity.Name, entity.Type, entity.Status)
	if err != nil {
		log.Printf("Error calling create_managed_entity: %v", err)
		return err
	}
	log.Printf("Managed entity created successfully: %+v", entity)
	return nil
}


// UpdateManagedEntity
func (r *AdminRepositoryImpl) UpdateManagedEntity(entity *model.ManagedEntity) error {
    // Step 1: Check if the record exists
    var exists bool
    err := r.db.QueryRow(`SELECT EXISTS(SELECT 1 FROM managed_entities WHERE id = $1)`, entity.ID).Scan(&exists)
    if err != nil {
        log.Printf("Error checking managed entity existence: %v", err)
        return err
    }

    if !exists {
        return fmt.Errorf("no managed entity found with ID %v", entity.ID)
    }

    // Step 2: Perform the update
    _, err = r.db.Exec(`CALL update_managed_entity($1, $2, $3)`,
        entity.ID, entity.Name, entity.Status)
    if err != nil {
        log.Printf("Error calling update_managed_entity: %v", err)
        return err
    }

    log.Printf("Managed entity updated successfully: %+v", entity)
    return nil
}

// Delete Managed Entity
func (r *AdminRepositoryImpl) DeleteManagedEntity(id uuid.UUID) error {
    // Step 1: Check if the record exists
    var exists bool
    err := r.db.QueryRow(`SELECT EXISTS(SELECT 1 FROM managed_entities WHERE id = $1)`, id).Scan(&exists)
    if err != nil {
        log.Printf("Error checking managed entity existence: %v", err)
        return err
    }

    if !exists {
        return fmt.Errorf("no managed entity found with ID %v", id)
    }

    // Step 2: Perform the delete
    _, err = r.db.Exec(`CALL delete_managed_entity($1)`, id)
    if err != nil {
        log.Printf("Error calling delete_managed_entity: %v", err)
        return err
    }

    log.Printf("Managed entity deleted successfully: %v", id)
    return nil
}


// Get All Managed Entities
func (r *AdminRepositoryImpl) GetAllManagedEntities() ([]*model.ManagedEntity, error) {
	rows, err := r.db.Query(`SELECT * FROM get_all_managed_entities()`)
	if err != nil {
		log.Printf("Error querying get_all_managed_entities: %v", err)
		return nil, err
	}
	defer rows.Close()

	var entities []*model.ManagedEntity
	for rows.Next() {
		var entity model.ManagedEntity
		err := rows.Scan(
			&entity.ID,
			&entity.Name,
			&entity.Type,
			&entity.Status,
			&entity.CreatedAt,
		)
		if err != nil {
			log.Printf("Error scanning managed entity row: %v", err)
			return nil, err
		}
		entities = append(entities, &entity)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Row iteration error in managed entities: %v", err)
		return nil, err
	}

	log.Printf("Managed entities retrieved: %d", len(entities))
	return entities, nil
}

// =============================
//  Approval Requests
// =============================

// Create Approval Request
func (r *AdminRepositoryImpl) CreateApprovalRequest(req *model.ApprovalRequest) error {
	_, err := r.db.Exec(`CALL create_approval_request($1, $2, $3)`,
		req.ID, req.EntityType, req.EntityID)
	if err != nil {
		log.Printf("Error calling create_approval_request: %v", err)
		return err
	}
	log.Printf("Approval request created: %+v", req)
	return nil
}

// Update Approval Status
func (r *AdminRepositoryImpl) UpdateApprovalStatus(id uuid.UUID, status string, reviewedBy uuid.UUID) error {
	_, err := r.db.Exec(`CALL update_approval_status($1, $2, $3)`,
		id, status, reviewedBy)
	if err != nil {
		log.Printf("Error calling update_approval_status: %v", err)
		return err
	}
	log.Printf("Approval request (ID=%v) updated to status: %s", id, status)
	return nil
}

// Get All Approval Requests
func (r *AdminRepositoryImpl) GetAllApprovalRequests() ([]*model.ApprovalRequest, error) {
	rows, err := r.db.Query(`SELECT * FROM get_all_approval_requests()`)
	if err != nil {
		log.Printf("Error querying get_all_approval_requests: %v", err)
		return nil, err
	}
	defer rows.Close()

	var requests []*model.ApprovalRequest
	for rows.Next() {
		var req model.ApprovalRequest
		err := rows.Scan(
			&req.ID,
			&req.EntityType,
			&req.EntityID,
			&req.RequestDate,
			&req.Status,
			&req.ReviewedBy,
			&req.ReviewedAt,
		)
		if err != nil {
			log.Printf("Error scanning approval request: %v", err)
			return nil, err
		}
		requests = append(requests, &req)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Row iteration error in approval requests: %v", err)
		return nil, err
	}

	log.Printf("Approval requests retrieved: %d", len(requests))
	return requests, nil
}

// Get Pending Approvals
func (r *AdminRepositoryImpl) GetPendingApprovals() ([]*model.ApprovalRequest, error) {
	rows, err := r.db.Query(`SELECT * FROM get_pending_approvals()`)
	if err != nil {
		log.Printf("Error querying get_pending_approvals: %v", err)
		return nil, err
	}
	defer rows.Close()

	var pending []*model.ApprovalRequest
	for rows.Next() {
		var req model.ApprovalRequest
		err := rows.Scan(
			&req.ID,
			&req.EntityType,
			&req.EntityID,
			&req.RequestDate,
			&req.Status,
		)
		if err != nil {
			log.Printf("Error scanning pending approval: %v", err)
			return nil, err
		}
		pending = append(pending, &req)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Row iteration error in pending approvals: %v", err)
		return nil, err
	}

	log.Printf("Pending approvals retrieved: %d", len(pending))
	return pending, nil
}

// =============================
//  System Settings
// =============================

// Upsert System Settings
func (r *AdminRepositoryImpl) UpsertSystemSettings(setting *model.SystemSettings) error {
	_, err := r.db.Exec(`CALL upsert_system_settings($1, $2, $3)`,
		setting.ID, setting.PaymentGateway, setting.Theme)
	if err != nil {
		log.Printf("Error calling upsert_system_settings: %v", err)
		return err
	}
	log.Printf("System settings upserted: %+v", setting)
	return nil
}

// Get Latest System Settings
func (r *AdminRepositoryImpl) GetLatestSystemSettings() (*model.SystemSettings, error) {
	row := r.db.QueryRow(`SELECT * FROM get_latest_system_settings()`)

	var setting model.SystemSettings
	err := row.Scan(
		&setting.ID,
		&setting.PaymentGateway,
		&setting.Theme,
		//&setting.EmailTemplateID,
		&setting.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no system settings found")
		}
		log.Printf("Error scanning system settings: %v", err)
		return nil, err
	}

	log.Println("Latest system settings retrieved successfully.")
	return &setting, nil
}

// =============================
// Constructor
// =============================
func NewAdminRepository(db *sql.DB) repository.AdminRepository {
	return &AdminRepositoryImpl{db: db}
}

