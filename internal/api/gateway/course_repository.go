package gateway

import (
	"database/sql"
	"e-learning-system/internal/domain/model"
	"e-learning-system/internal/domain/repository"
	"fmt"
	"log"

	"github.com/gofrs/uuid"
)

type CourseRepositoryImpl struct {
	db *sql.DB
}

// FindInstructor implements repository.CourseRepository.
func (r *CourseRepositoryImpl) FindInstructor(InstructorID string) (*model.Course, error) {
	var course model.Course

	query := `SELECT * FROM et_courses_by_instructor($1)`
	err := r.db.QueryRow(InstructorID, query).Scan(
		&course.ID,
		&course.Title,
		&course.Description,
		&course.InstructorID,
	)
	if err != nil {
		 log.Printf("failed to find instructor ID %v ", err)
	}

	return  &course, nil
}

// Create implements repository.CourseRepository.
func (r *CourseRepositoryImpl) Create(Course *model.Course) error {
	_, err := r.db.Exec(`Call create_course($1, $2, $3, $4)`,
		Course.ID,
		Course.Title,
		Course.Description,
		Course.InstructorID)
	if err != nil {
		log.Printf("Error calling create_course: %v", err)
		return err
	}

	log.Printf("Course created: %+v", Course)
	return nil

}

// Delete implements repository.CourseRepository.
func (r *CourseRepositoryImpl) Delete(CourseID uuid.UUID) error {
	_, err := r.db.Exec(`CALL delete_course($1)`, CourseID)
	if err != nil {
		log.Printf("Error calling delete_course for ID %v: %v", CourseID, err)
		return err
	}

	log.Printf("Course soft-deleted: %v", CourseID)
	return nil
}

// GetAll implements repository.CourseRepository.
func (r *CourseRepositoryImpl) GetAll() ([]*model.Course, error) {
	rows, err := r.db.Query(`SELECT * FROM get_all_courses()`)
	if err != nil {
		log.Printf("Error querying get_all_courses: %v", err)
		return nil, err
	}
	defer rows.Close()

	var courses []*model.Course

	for rows.Next() {
		var course model.Course
		err := rows.Scan(
			&course.ID,
			&course.Title,
			&course.Description,
			&course.InstructorID,
		)
		if err != nil {
			log.Printf("Error scanning course row: %v", err)
			return nil, err
		}
		courses = append(courses, &course)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Row iteration error: %v", err)
		return nil, err
	}

	log.Printf("Courses retrieved: %d", len(courses))
	return courses, nil
}

// GetByID implements repository.CourseRepository.
func (r *CourseRepositoryImpl) GetByID(CourseID uuid.UUID) (*model.Course, error) {
	var course model.Course

	query := `SELECT * FROM get_course_by_id($1)`
	row := r.db.QueryRow(query, CourseID)

	err := row.Scan(
		&course.ID,
		&course.Title,
		&course.Description,
		&course.InstructorID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("course not found")
			return nil, fmt.Errorf("course not found")
		}
		log.Printf("DB error: %v", err)
		return nil, err
	}

	return &course, nil
}

// Update implements repository.CourseRepository.
func (r *CourseRepositoryImpl) Update(Course *model.Course) error {
	_, err := r.db.Exec(`CALL update_course($1, $2, $3, $4)`,
		Course.ID, Course.Title, Course.Description, Course.InstructorID)
	if err != nil {
		log.Printf("Error calling update_course: %v", err)
		return err
	}

	log.Printf("Course updated: %+v", Course)
	return nil
}

func NewCourserRepositry(db *sql.DB) repository.CourseRepository {
	return &CourseRepositoryImpl{db: db}
}
