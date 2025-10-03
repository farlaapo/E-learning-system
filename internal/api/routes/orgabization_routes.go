package routes

import (
	"e-learning-system/internal/api/controller"
	"e-learning-system/internal/api/middleware"
	"e-learning-system/internal/domain/repository"

	"github.com/gin-gonic/gin"
)

// RegisteOrganizationRoutes registers enrollment-related routes
func RegisterOrganizationRoutes(router *gin.Engine, orgController controller.OrgsController, tokenRepo repository.TokenRepository) {
  
	// Auth-middleware
	authMidlleware := middleware.AuthMiddleware(tokenRepo)
 
	orgGroup := router.Group("/organizations")
	{
     orgGroup.Use(authMidlleware)
		 {
			orgGroup.POST("", orgController.CreateOrganization)
			orgGroup.GET("/", orgController.GetAllOrganization)
			orgGroup.GET("/:id", orgController.GetOrganizationById)
			orgGroup.PUT("/:id", orgController.UpdateOrganization)
			orgGroup.DELETE("/:id", orgController.DeleteOrganization)
		 }
	}
	


}