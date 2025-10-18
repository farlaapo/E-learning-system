package gateway

import (
	"database/sql"
	"e-learning-system/internal/domain/model"
	"e-learning-system/internal/domain/repository"
	"log"

	"github.com/gofrs/uuid"
	"github.com/lib/pq"
)

// / OrganizationRepositoryImpl struct
type OrganizationRepositoryImpl struct {
	db *sql.DB
}

// Create implements repository.OrganizationgRepository.
func (r *OrganizationRepositoryImpl) Create(Organization *model.Organization) error {
	   _, err := r.db.Exec(`CALL  create_orgs($1, $2, $3,$4, $5) `,
		  Organization.ID,
		  Organization.Name,
	    Organization.Description,
      Organization.OwnerID,
      pq.Array(Organization.Tutors))
			 if err != nil {
        log.Printf("Error calling create_orgs: %v", err)
        return err
    }

    log.Printf("orgs created %+v", Organization)
    return nil

}

// Delete implements repository.OrganizationgRepository.
func (r *OrganizationRepositoryImpl) Delete(OrganizationID uuid.UUID) error {
	_, err := r.db.Exec(`CALL  delete_orgs($1)`, OrganizationID)
	if err != nil {
		log.Printf("Error calling update_orgs for ID %v, %v", err, OrganizationID)
		return err
	}

	log.Printf("orgs created %+v", OrganizationID)
	return nil
}

// GetAll implements repository.OrganizationgRepository.
func (r *OrganizationRepositoryImpl) GetAll() ([]*model.Organization, error) {
	query :=  `SELECT * FROM  get_all_org()`
	rows, err := r.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var orgs []*model.Organization
    for rows.Next() {
        var o model.Organization
        err := rows.Scan(
            &o.ID,
			&o.Name,
			&o.Description,
			&o.OwnerID,
			pq.Array(&o.Tutors),
            &o.CreatedAt,
            &o.UpdatedAt,
            &o.DeletedAt,
        )
        if err != nil {
            return nil, err
        }
        orgs = append(orgs, &o)
    }

    return orgs, nil

}

// GetByID implements repository.OrganizationgRepository.
func (r *OrganizationRepositoryImpl) GetByID(OrganizationID uuid.UUID) (*model.Organization, error) {
	
    query := `SELECT * FROM get_org_by_id($1)`
    row := r.db.QueryRow(query, OrganizationID)

    var org model.Organization
    var deletedAt sql.NullTime

    err := row.Scan(
        &org.ID,
				&org.Name,
				&org.Description,
				&org.OwnerID,
				pq.Array(&org.Tutors),
        &org.CreatedAt,
        &org.UpdatedAt,
        &org.DeletedAt,
    )
    if err != nil {
        return nil, err
    }

    // Handle nullable deleted_at
    if deletedAt.Valid {
        org.DeletedAt = &deletedAt.Time
    }

    return &org, nil
}

// Update implements repository.OrganizationgRepository.
func (r *OrganizationRepositoryImpl) Update(Organization *model.Organization) error {
	_, err := r.db.Exec(`CALL update_org($1, $2, $3)`,
		Organization.ID,
		Organization.Name,
		Organization.Description,)
	if err != nil {
		log.Printf("error calling update orgs")
		return err
	}

	log.Printf("orgs updated %+v ", Organization)
	return nil
}

// OrganizationRepository instance
func NewOrganizationRepository(db *sql.DB) repository.OrganizationgRepository {
	return &OrganizationRepositoryImpl{
		db: db,
	}
}
