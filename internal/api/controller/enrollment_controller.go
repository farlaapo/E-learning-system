package controller

import (
	"e-learning-system/internal/domain/model"
	"e-learning-system/internal/domain/service"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

// enrollment_controller struct
type EnrollmentController struct {
	enrollmentService service.EnrollmentService
}
// enrollment_controller instance
func NewEnrollmentController(entollmentService service.EnrollmentService) *EnrollmentController {
	return &EnrollmentController{
		enrollmentService: entollmentService,
	}
}
// CreateEnrollment handler
func (Ec *EnrollmentController)  CreateEnrollment(c * gin.Context) {
   var enrollment model.Enrollment 

	if err := c.BindJSON(&enrollment); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	createdCourse, err := Ec.enrollmentService.CreateEnrollment(enrollment.CourseID, enrollment.UserID, enrollment.CertificateTemplate, enrollment.Completed  )
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(201, createdCourse)
}

// UpdateEnrollment handler
func (Ec *EnrollmentController) UpdateEnrollment(c * gin.Context) {
	var enrollmnet model.Enrollment
	// param
	enrollmnetParam := c.Param("id")
	enrollmnetID, err := uuid.FromString(enrollmnetParam)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
  // bind with json
	if err := c.BindJSON(&enrollmnet); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	enrollmnet.ID = enrollmnetID

	// save 
	if err := Ec.enrollmentService.UpdateEnrollment(&enrollmnet); err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	//retun
	c.JSON(200, "succesfully updated!")


}
// get enrollment byID handler
func (Ec *EnrollmentController) GetEnrollmentById(c * gin.Context) {
	// param
	enrollmnetParam := c.Param("id")
	enrollmnetID, err := uuid.FromString(enrollmnetParam)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
  // save service 
	enrollmnet, err := Ec.enrollmentService.GetEnrollmentById(enrollmnetID)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	// return
	c.JSON(200, enrollmnet)
}
// GetAllEnrollment handler
func (Ec *EnrollmentController)  GetAllEnrollment(c * gin.Context) {
	enrollment, err := Ec.enrollmentService.GetAllEnrollment()
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, enrollment)
}

// DeletEnrollment handler

func (Ec *EnrollmentController) DeletEnrollment(c * gin.Context) {
		// param
	enrollmnetParam := c.Param("id")
	enrollmnetID, err := uuid.FromString(enrollmnetParam)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	// save service
	err = Ec.enrollmentService.DeletEnrollment(enrollmnetID)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	// return
	c.JSON(200, "succesfully deleted")

}







