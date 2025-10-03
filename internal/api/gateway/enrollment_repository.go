package gateway

import (
	"database/sql"
	"e-learning-system/internal/domain/model"
	"e-learning-system/internal/domain/repository"
	
	"log"

	"github.com/gofrs/uuid"
)

// enrollmentRepositoryImpl
type enrollmentRepositoryImpl struct {
	db *sql.DB
}

func (r *enrollmentRepositoryImpl) GetByID(enrollmentID uuid.UUID) (*model.Enrollment, error) {
	  
    query := `SELECT * FROM get_enrollment_by_id($1)`
    row := r.db.QueryRow(query, enrollmentID)

    var enrollment model.Enrollment
    var deletedAt sql.NullTime

    err := row.Scan(
        &enrollment.ID,
        &enrollment.CourseID,
        &enrollment.UserID,
        &enrollment.EnrollmentAt,
        &enrollment.Completed,
        &enrollment.CertificateIssuedAt,
        &enrollment.CertificateTemplate,
        &enrollment.CreatedAt,
        &enrollment.UpdatedAt,
        &deletedAt,
    )
    if err != nil {
        return nil, err
    }

    // Handle nullable deleted_at
    if deletedAt.Valid {
        enrollment.DeletedAt = &deletedAt.Time
    }

    return &enrollment, nil
}



// Create implements repository.EnrolledRepository.
func (r *enrollmentRepositoryImpl) Create(enrollment *model.Enrollment) error {

    _, err := r.db.Exec(`CALL create_enrollment($1, $2, $3, $4, $5, $6)`,
        enrollment.ID,
        enrollment.CourseID,
        enrollment.UserID,
        enrollment.Completed,
        enrollment.CertificateIssuedAt,
        enrollment.CertificateTemplate,
    )
    
    if err != nil {
        log.Printf("Error calling create_enrollment: %v", err)
        return err
    }

    log.Printf("Enrollment created %+v", enrollment)
    return nil
}

// Delete implements repository.EnrolledRepository.
func (r *enrollmentRepositoryImpl) Delete(enrollmentID uuid.UUID) error {
	_, err := r.db.Exec(`CALL delete_enrollment($1)`, enrollmentID)
	if err != nil {
		log.Printf("Error calling update_enrollment for ID %v, %v", err, enrollmentID)
		return err
	}

	log.Printf("Course created %+v", enrollmentID)
	return nil
}

// GetAll implements repository.EnrolledRepository.
func (r *enrollmentRepositoryImpl) GetAll() ([]*model.Enrollment, error) {
    query := `SELECT * FROM get_all_enrollments()`
    rows, err := r.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var enrollments []*model.Enrollment
    for rows.Next() {
        var e model.Enrollment
        err := rows.Scan(
            &e.ID,
            &e.CourseID,
            &e.UserID,
            &e.EnrollmentAt,
            &e.Completed,
            &e.CertificateIssuedAt,
            &e.CertificateTemplate,
            &e.CreatedAt,
            &e.UpdatedAt,
            &e.DeletedAt,
        )
        if err != nil {
            return nil, err
        }
        enrollments = append(enrollments, &e)
    }

    return enrollments, nil
}



// Update implements repository.EnrolledRepository.
func (r *enrollmentRepositoryImpl) Update(enrollment *model.Enrollment) error {
	_, err := r.db.Exec(`CALL update_enrollment($1, $2, $3)`,
		enrollment.ID,
		enrollment.Completed,
		enrollment.CertificateTemplate)
	if err != nil {
		log.Printf("error calling update enrollments")
		return err
	}

	log.Printf("enrollment updated %+v ", enrollment)
	return nil
}

func NewEnrollmentRepository(db *sql.DB) repository.EnrolledRepository {
	return &enrollmentRepositoryImpl{db: db}
}
