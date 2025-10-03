package controller

import (
	"e-learning-system/internal/domain/model"
	"e-learning-system/internal/domain/service"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

type OrgsController struct {
	orgService service.OrganizationService
}

// CreateOrganization implements controller
func (oC *OrgsController) CreateOrganization(c * gin.Context) {
	   var org model.Organization 

	if err := c.BindJSON(&org); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	
	createdOrg, err := oC.orgService.CreateOrganization(org.Name, org.Description, org.OwnerID, org.Tutors )
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(201, createdOrg)
}

// DeleteOrganization implements controller.
func (oC *OrgsController) DeleteOrganization(c * gin.Context) {
		// param
	orgParam := c.Param("id")
	orgID, err := uuid.FromString(orgParam)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	// save service
	err = oC.orgService.DeleteOrganization(orgID)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	// return
	c.JSON(200, "succesfully deleted")
}

// GetAllOrganization implements controller
func (oC *OrgsController) GetAllOrganization(c * gin.Context) {
	org, err := oC.orgService.GetAllOrganization()
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, org)
}

// GetOrganizationById implements controller.
func (oC *OrgsController) GetOrganizationById(c * gin.Context){
	// param
	orgParam := c.Param("id")
	orgID, err := uuid.FromString(orgParam)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
  // save service 
	org, err := oC.orgService.GetOrganizationById(orgID)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	// return
	c.JSON(200, org)
}

// UpdateOrganization implements controller
func (oC *OrgsController) UpdateOrganization(c * gin.Context) {
	 var org model.Organization
	// param
	orgParam := c.Param("id")
	orgID, err := uuid.FromString(orgParam)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
  // bind with json
	if err := c.BindJSON(&org); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	org.ID = orgID

	// save 
	if err := oC.orgService.UpdateOrganization(&org); err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	//retun
	c.JSON(200, "succesfully updated!")
}

func NeworgsController(orgService service.OrganizationService) *OrgsController{
	return &OrgsController{
		orgService: orgService,
	}

}
