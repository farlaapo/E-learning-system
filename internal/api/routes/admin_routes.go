package routes

import (
	"e-learning-system/internal/api/controller"
	"e-learning-system/internal/api/middleware"
	"e-learning-system/internal/domain/repository"

	"github.com/gin-gonic/gin"
)

// RegisterAdminRoutes registers all admin-related routes
func RegisterAdminRoutes(router *gin.Engine, adminController *controller.AdminController, tokenRepo repository.TokenRepository) {
	// Auth middleware
	authMiddleware := middleware.AuthMiddleware(tokenRepo)

	// Group all /admin endpoints
	adminGroup := router.Group("/admin")
	{
		adminGroup.Use(authMiddleware)
		{
	
			// 1 Dashboard
		
			adminGroup.GET("/dashboard", adminController.GetDashboardSummary)

			
			// 2 Managed Entities
			
			adminGroup.POST("/entities", adminController.CreateManagedEntity)
			adminGroup.GET("/entities", adminController.ListAllManagedEntities)
			adminGroup.PUT("/entities/:id", adminController.UpdateManagedEntity)
			adminGroup.DELETE("/entities/:id", adminController.DeleteManagedEntity)

			
			// 3 Approval Requests
			
			adminGroup.POST("/approvals", adminController.CreateApprovalRequest)
			adminGroup.GET("/approvals", adminController.ListAllApprovalRequests)
			adminGroup.GET("/approvals/pending", adminController.ListPendingApprovals)
			adminGroup.PUT("/approvals/:id", adminController.UpdateApprovalStatus)

			
			// 4 System Settings
		
			adminGroup.POST("/settings", adminController.UpsertSystemSettings)
			adminGroup.GET("/settings/latest", adminController.GetLatestSystemSettings)
		}
	}
}
