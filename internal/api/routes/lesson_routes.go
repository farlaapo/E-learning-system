package routes

import (
	"e-learning-system/internal/api/controller"
	"e-learning-system/internal/api/middleware"
	"e-learning-system/internal/domain/repository"

	"github.com/gin-gonic/gin"
)

// RegisteEnrollmentRoutes registers enrollment-related routes
func RegisterLessonRoutes(router *gin.Engine, lessonController controller.LessonController, tokenRepo repository.TokenRepository) {
  
	// Auth-middleware
	authMidlleware := middleware.AuthMiddleware(tokenRepo)
 
	lessonGroup := router.Group("/lessons")
	{
     lessonGroup.Use(authMidlleware)
		 {
			lessonGroup.POST("", lessonController.CreateLesson)
			lessonGroup.GET("/", lessonController.GetAllLesson)
			lessonGroup.GET("/:id", lessonController.GetLessonById)
			lessonGroup.PUT("/:id", lessonController.UpdateLesson)
			lessonGroup.DELETE("/:id", lessonController.DeletLesson)
		 }
	}
	


}