package controller

import (
	"e-learning-system/internal/domain/model"
	"e-learning-system/internal/domain/service"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

type LessonController struct {
	lessonService service.LessonService
}

// CreateLesson implements controller
func (lC *LessonController) CreateLesson(c * gin.Context) {
	 var lesson model.Lesson 

	if err := c.BindJSON(&lesson); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	
	createdLesson, err := lC.lessonService.CreateLesson(lesson.ModuleID,  lesson.Title, lesson.Content, lesson.VideoURL, lesson.Order)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(201, createdLesson)
}

// DeletLesson implements controller
func (lC *LessonController) DeletLesson(c * gin.Context)  {
	 	// param
	lessonParam := c.Param("id")
	lessonID, err := uuid.FromString(lessonParam)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	// save service
	err = lC.lessonService.DeletLesson(lessonID)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	// return
	c.JSON(200, "succesfully deleted")
}

// GetAllLesson implements controller
func (lC *LessonController) GetAllLesson(c * gin.Context) {
	lesson, err := lC.lessonService.GetAllLesson()
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, lesson)
}

// GetLessonById implements controller
func (lC *LessonController) GetLessonById(c * gin.Context) {
	// param
	lessonParam := c.Param("id")
	lessonID, err := uuid.FromString(lessonParam)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
  // save service 
	lesson, err := lC.lessonService.GetLessonById(lessonID)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	// return
	c.JSON(200, lesson)
}

// UpdateLesson implements controller
func (lC *LessonController) UpdateLesson(c * gin.Context) {
	var lesson model.Lesson
	// param
	lessonParam := c.Param("id")
	lessonID, err := uuid.FromString(lessonParam)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
  // bind with json
	if err := c.BindJSON(&lesson); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	lesson.ID = lessonID

	// save 
	if err := lC.lessonService.UpdateLesson(&lesson); err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	//retun
	c.JSON(200, "succesfully updated!",)
}

func NewLessonController(lessonService service.LessonService) *LessonController {
	return &LessonController{
		lessonService: lessonService,
	}
}
