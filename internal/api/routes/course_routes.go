package routes

import (
	"e-learning-system/internal/api/controller"
	"e-learning-system/internal/api/middleware"
	"e-learning-system/internal/domain/repository"

	"github.com/gin-gonic/gin"
)

// RegisteCourseRoutes registers user-related routes
func RegisterCourseRoutes(router *gin.Engine, courseController *controller.CourseController, tokenRepo repository.TokenRepository) {
	// Auth middleware
	authMiddleware := middleware.AuthMiddleware(tokenRepo)

	// Group all /course endpoints
	courseGroup := router.Group("/courses")
	{
		courseGroup.Use(authMiddleware)
		{
			courseGroup.GET("", courseController.CreateCourse)
			courseGroup.GET("/:id", courseController.GetCourseById)
			courseGroup.PUT("/:id", courseController.UpdateCourse)
			courseGroup.DELETE("/:id", courseController.DeleteCourse)
		}
	}
}