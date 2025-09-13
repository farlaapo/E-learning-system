package gateway

import (
	"database/sql"
	"e-learning-system/internal/domain/model"
	"e-learning-system/internal/domain/repository"
	"fmt"
	"log"

	"github.com/gofrs/uuid"
)

// enrollmentRepositoryImpl
type enrollmentRepositoryImpl struct {
	db *sql.DB
}

// Create implements repository.EnrolledRepository.
func (r *enrollmentRepositoryImpl) Create(enrollment *model.Enrollment) error {
	_, err := r.db.Exec(`Call  create_enrollment ($1, $2 $3 )`,
	   enrollment.ID, 
		 enrollment.CourseID,
		 enrollment.UserID,)
		 if err != nil {
			log.Printf("Error calling  create_enrollment: %v", err)
			return  err
		 }
		 log.Printf("Course created %+v" ,enrollment)
		 return  nil
}

// Delete implements repository.EnrolledRepository.
func (r *enrollmentRepositoryImpl) Delete(enrollmentID uuid.UUID) error {
	_, err := r.db.Exec(`CALL delete_enrollment($1)`,enrollmentID)
	if err != nil {
		log.Printf("Error calling update_enrollment for ID %v, %v", err , enrollmentID)
		return  err
	}

	log.Printf("Course created %+v" ,enrollmentID)
		 return  nil
}

// GetAll implements repository.EnrolledRepository.
func (r *enrollmentRepositoryImpl) GetAll() ([]*model.Enrollment, error) {
	rows, err := r.db.Query(`SELECT * FROM  get_all_enrollments()`)
	if err != nil {
		log.Printf("error calling get_all_enrollments")
		return  nil, err
	}

	defer rows.Close()

	var enrollments []*model.Enrollment

	for rows.Next() {
		var enrollment model.Enrollment
		err := rows.Scan(
			&enrollment.ID,
			&enrollment.CourseID,
			&enrollment.UserID,
			&enrollment.EnrolledAt,
			&enrollment.Completed,
			&enrollment.CertificateIssuedAt, 
			&enrollment.CertificateTemplate,
			&enrollment.Created_at, 
			&enrollment.Updated_at,
		)
		if err != nil {
			log.Printf("error scaning enrollment")
			return  nil, err
		}
		enrollments = append(enrollments, &enrollment)
	}
	if err = rows.Err(); err != nil  {
		log.Printf("Row iterration")
		return   nil, err
	}
	return   enrollments, nil

}

// GetByID implements repository.EnrolledRepository.
func (r *enrollmentRepositoryImpl) GetByID(enrollmentID uuid.UUID) (*model.Enrollment, error) {
	var enrollment model.Enrollment


	query := `SELECT * FROM get_enrollment_by_course($1)`
	rows := r.db.QueryRow(query, enrollmentID)
	err := rows.Scan(
		&enrollment.ID,
			&enrollment.CourseID,
			&enrollment.UserID,
			&enrollment.EnrolledAt,
			&enrollment.Completed,
			&enrollment.CertificateIssuedAt, 
			&enrollment.CertificateTemplate,
			&enrollment.Created_at, 
			&enrollment.Updated_at,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("enrollments  not found")
			return  nil, fmt.Errorf("enrollments  not found")
		}
		log.Printf("DB error: %v", err)
		return  nil, err
	}

	return  &enrollment, nil
}

// Update implements repository.EnrolledRepository.
func (r *enrollmentRepositoryImpl) Update(enrollment *model.Enrollment) error {
	_, err := r.db.Exec(`CALL update_enrollment($1, $2, $3)`,
       enrollment.ID,
       enrollment.Completed,
       enrollment.CertificateTemplate)
			 if err != nil {
				log.Printf("error calling update enrollments")
				return  err
			 }

			 log.Printf("enrollment updated %+v ", enrollment)
			 return  nil
}

func NewEnrollmentRepository(db *sql.DB) repository.EnrolledRepository {
	return &enrollmentRepositoryImpl{db: db}
}
