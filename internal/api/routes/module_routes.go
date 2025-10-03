package routes

import (
	"e-learning-system/internal/api/controller"
	"e-learning-system/internal/api/middleware"
	"e-learning-system/internal/domain/repository"

	"github.com/gin-gonic/gin"
)

// RegisteModuleRoutes registers module-related routes
func RegisterModuleRoutes(router *gin.Engine, moduleController controller.ModuleController, tokenRepo repository.TokenRepository) {
  
	// Auth-middleware
	authMidlleware := middleware.AuthMiddleware(tokenRepo)
 
	moduleGroup := router.Group("/modules")
	{
     moduleGroup.Use(authMidlleware)
		 {
			moduleGroup.POST("", moduleController.CreateModule)
			moduleGroup.GET("/", moduleController.GetAllModule)
			moduleGroup.GET("/:id", moduleController.GetModuleById)
			moduleGroup.PUT("/:id", moduleController.UpdateModule)
			moduleGroup.DELETE("/:id", moduleController.DeletModule)
		 }
	}
	


}