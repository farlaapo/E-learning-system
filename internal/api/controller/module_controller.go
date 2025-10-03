package controller

import (
	"e-learning-system/internal/domain/model"
	"e-learning-system/internal/domain/service"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

type ModuleController struct {
	moduLeService service.ModuleService
}

// CreateEnrollment implements controller
func (mC *ModuleController) CreateModule(c * gin.Context) {
	var module model.Module 

	if err := c.BindJSON(&module); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	
	createdModule, err := mC.moduLeService.CreateModule(module.CourseID, module.Title, module.Order )
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(201, createdModule)
}

// DeletEnrollment implements controller.
func (mC *ModuleController) DeletModule(c * gin.Context)  {
	// param
	moduleParam := c.Param("id")
	moduleID, err := uuid.FromString(moduleParam)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	// save service
	err = mC.moduLeService.DeletModule(moduleID)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	// return
	c.JSON(200, "succesfully deleted")

}

// GetAllEnrollment implements controller
func (mC *ModuleController) GetAllModule(c * gin.Context){
	enrollment, err := mC.moduLeService.GetAllModule()
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, enrollment)
}

// GetEnrollmentById implements controller
func (mC *ModuleController) GetModuleById(c * gin.Context) {
	// param
	moduleParam := c.Param("id")
	moduleID, err := uuid.FromString(moduleParam)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
  // save service 
	module, err := mC.moduLeService.GetModuleById(moduleID)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	// return
	c.JSON(200, module)
}

// UpdateEnrollment implements controller
func (mC *ModuleController) UpdateModule(c * gin.Context)  {
	var module model.Module
	// param
	moduleParam := c.Param("id")
	moduleID, err := uuid.FromString(moduleParam)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
  // bind with json
	if err := c.BindJSON(&moduleID); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	module.ID = moduleID

	// save 
	if err := mC.moduLeService.UpdateModule(&module); err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	//retun
	c.JSON(200, "succesfully updated!",)
}

func NewModuleController(moduLeService service.ModuleService) *ModuleController {
	return &ModuleController{
		moduLeService: moduLeService,
	}

}
