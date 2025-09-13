package routes

import (
	"e-learning-system/internal/api/controller"
	"e-learning-system/internal/api/middleware"
	"e-learning-system/internal/domain/repository"

	"github.com/gin-gonic/gin"
)

// RegisteCourseRoutes registers enrollment-related routes
func RegisterEnrollmentRoutes(router *gin.Engine, enrollmentController controller.EnrollmentController, tokenRepo repository.TokenRepository) {
  
	// Auth-middleware
	authMidlleware := middleware.AuthMiddleware(tokenRepo)
 
	enrollmentGroup := router.Group("/enrollment")
	{
     enrollmentGroup.Use(authMidlleware)
		 {
			enrollmentGroup.POST("", enrollmentController.CreateEnrollment)
			enrollmentGroup.GET("/", enrollmentController.GetAllEnrollment)
			enrollmentGroup.GET("/:id", enrollmentController.GetEnrollmentById)
			enrollmentGroup.PUT(":/id", enrollmentController.UpdateEnrollment)
			enrollmentGroup.DELETE(":/id", enrollmentController.DeletEnrollment)
		 }
	}
	


}