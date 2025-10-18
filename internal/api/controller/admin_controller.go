package controller

import (
	"e-learning-system/internal/domain/model"
	"e-learning-system/internal/domain/service"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

type AdminController struct {
	adminService service.AdminService
}

// Constructor
func NewAdminController(adminService service.AdminService) *AdminController {
	return &AdminController{
		adminService: adminService,
	}
}

// =============================
// 1 Dashboard
// =============================
func (ac *AdminController) GetDashboardSummary(c *gin.Context) {
	dashboard, err := ac.adminService.GetDashboardSummary()
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, dashboard)
}

// =============================
// 2 Managed Entities
// =============================
func (ac *AdminController) CreateManagedEntity(c *gin.Context) {
	var entity model.ManagedEntity

	if err := c.ShouldBindJSON(&entity); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	createdEntity, err := ac.adminService.CreateManagedEntity(entity.Name, entity.Type, entity.Status)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(201, createdEntity)
}

func (ac *AdminController) UpdateManagedEntity(c *gin.Context) {
	var entity model.ManagedEntity
 // param
	entityParam := c.Param("id")
	entityID, err := uuid.FromString(entityParam)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	if err := c.BindJSON(&entity); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	entity.ID = entityID

	if err := ac.adminService.UpdateManagedEntity(&entity); err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, entity)
}

func (ac *AdminController) DeleteManagedEntity(c *gin.Context) {
	entityParam := c.Param("id")
	entityID, err := uuid.FromString(entityParam)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	err = ac.adminService.DeleteManagedEntity(entityID)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Managed entity deleted successfully"})
}

func (ac *AdminController) ListAllManagedEntities(c *gin.Context) {
	entities, err := ac.adminService.ListAllManagedEntities()
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, entities)
}

// =============================
// 3 Approval Requests
// =============================
func (ac *AdminController) CreateApprovalRequest(c *gin.Context) {
	var req model.ApprovalRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	createdReq, err := ac.adminService.CreateApprovalRequest(req.EntityType, req.EntityID)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(201, createdReq)
}

func (ac *AdminController) UpdateApprovalStatus(c *gin.Context) {
	var req struct {
		Status     string    `json:"status"`
		ReviewedBy uuid.UUID `json:"reviewed_by"`
	}

	idParam := c.Param("id")
	requestID, err := uuid.FromString(idParam)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	if err := ac.adminService.UpdateApprovalStatus(requestID, req.Status, req.ReviewedBy); err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Approval status updated successfully"})
}

func (ac *AdminController) ListAllApprovalRequests(c *gin.Context) {
	requests, err := ac.adminService.ListAllApprovalRequests()
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, requests)
}

func (ac *AdminController) ListPendingApprovals(c *gin.Context) {
	pending, err := ac.adminService.ListPendingApprovals()
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, pending)
}

// =============================
// 4 System Settings
// =============================
func (ac *AdminController) UpsertSystemSettings(c *gin.Context) {
	var setting model.SystemSettings

	if err := c.ShouldBindJSON(&setting); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	updatedSetting, err := ac.adminService.UpsertSystemSettings(setting.PaymentGateway, setting.Theme)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, updatedSetting)
}

func (ac *AdminController) GetLatestSystemSettings(c *gin.Context) {
	setting, err := ac.adminService.GetLatestSystemSettings()
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, setting)
}
