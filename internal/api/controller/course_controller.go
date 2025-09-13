package controller

import (
	"e-learning-system/internal/domain/model"
	"e-learning-system/internal/domain/service"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

type CourseController struct {
	courseService service.CourseService
}

func NewCourseController(courseService service.CourseService) *CourseController {
	return &CourseController{
		courseService: courseService,
	}
}

func (Cc *CourseController) CreateCourse(c *gin.Context) {
	var course model.Course

	if err := c.BindJSON(&course); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	createdCourse, err := Cc.courseService.CreateCourse(course.Title, course.Description, course.InstructorID)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(201, createdCourse)
}

func (Cc *CourseController) UpdateCourse(c *gin.Context) {
	var course model.Course

	coursParam := c.Param("id")
	coursID, err := uuid.FromString(coursParam)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	if err := c.BindJSON(&course); err != nil {
		c.JSON(400, gin.H{"messsage": err.Error()})
		return
	}

	course.ID = coursID

	if err := Cc.courseService.UpdateCourse(&course); err != nil {
		c.JSON(500, gin.H{"messsage": err.Error()})
		return
	}

	c.JSON(200, course)

}

func (Cc *CourseController) GetCourseById(c *gin.Context) {
	coursParam := c.Param("id")
	courseID, err := uuid.FromString(coursParam)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	course, err := Cc.courseService.GetCourseById(courseID)

	if err != nil {
		c.JSON(500, gin.H{"messsage": err.Error()})
		return
	}

	c.JSON(200, course)

}

func (Cc *CourseController) DeleteCourse(c *gin.Context) {
	// param id
	coursParam := c.Param("id")
	courseID, err := uuid.FromString(coursParam)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	// service
	err = Cc.courseService.DeleteCourse(courseID)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	// return
	c.JSON(200, "successfully Updated")

}

func (Cc *CourseController) ListAllCourse(c *gin.Context) {
	course, err := Cc.courseService.ListAllCourse()
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, course)

}
