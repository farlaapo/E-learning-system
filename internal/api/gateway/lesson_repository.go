package gateway

import (
	"database/sql"
	"e-learning-system/internal/domain/model"
	"e-learning-system/internal/domain/repository"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/lib/pq"

	"log"
)

// LessonRepository  struct
type LessonRepository struct {
	db *sql.DB
}

// Create implements repository.LessonRepository.
func (r *LessonRepository) Create(lesson *model.Lesson) error {
	 _, err := r.db.Exec(`CALL create_lesson($1, $2, $3, $4, $5, $6)`,
	     lesson.ID,
			 lesson.ModuleID,
			 lesson.Title,
			 lesson.Content,
			 pq.Array(lesson.VideoURL),
			 lesson.Order,
    )

    if err != nil {
        log.Printf("Error calling create_lesson: %v", err)
        return err
    }

    log.Printf("lesson created %+v", lesson)
    return nil

}

// Delete implements repository.LessonRepository.
func (r *LessonRepository) Delete(lessonID uuid.UUID) error {
	_, err := r.db.Exec(`CALL  delete_lesson($1)`, lessonID)
	if err != nil {
		log.Printf("Error calling update_lesson for ID %v, %v", err, lessonID)
		return err
	}

	log.Printf("lesson created %+v", lessonID)
	return nil
}

// GetAll implements repository.LessonRepository.
func (r *LessonRepository) GetAll() ([]*model.Lesson, error) {
	query := `SELECT * FROM get_all_lesson()`
	row , err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer row.Close()

	var lessons []*model.Lesson
	for row.Next() {
		var l model.Lesson
		err := row.Scan(
			 &l.ID,
			 &l.ModuleID,
			 &l.Title,
			 &l.Content,
			 pq.Array(&l.VideoURL),
			 &l.Order,
			 &l.CreatedAt,
       &l.UpdatedAt,
       &l.DeletedAt,
		)
		if err != nil {
			return nil, err
		}

		lessons = append(lessons, &l)
	}
	return  lessons, nil
}

// GetByID implements repository.LessonRepository.
func (r *LessonRepository) GetByID(lessonID uuid.UUID) (*model.Lesson, error) {
	 var l  model.Lesson
	 query := `SELECT * FROM get_lesson_by_id($1)`
	 row := r.db.QueryRow(query, lessonID)
	 err := row.Scan(
		  &l.ID,
			 &l.ModuleID,
			 &l.Title,
			 &l.Content,
			 pq.Array(&l.VideoURL),
			 &l.Order,
			 &l.CreatedAt,
       &l.UpdatedAt,
       &l.DeletedAt,
	 )
	 if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("lesson not found")
			return nil, fmt.Errorf("lesson not found")
		}
		log.Printf("DB error: %v", err)
		return nil, err
	}

	return   &l, nil

}

// Update implements repository.LessonRepository.
func (r *LessonRepository) Update(lesson *model.Lesson) error {
	_, err  := r.db.Exec(`CALL update_lesson($1, $2, $3, $4, $5)`,
     lesson.ID,
		 lesson.Title,
		 lesson.Content,
		pq.Array(lesson.VideoURL),
		lesson.Order,
    )

		if err != nil {
        log.Printf("Error calling update_lesson: %v", err)
        return err
    }

    log.Printf("lesson updated %+v", lesson)
    return nil

}

func NewLessonRepository(db *sql.DB) repository.LessonRepository {
	return &LessonRepository{
		db: db,
	}
}
