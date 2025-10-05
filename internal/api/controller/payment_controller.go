package controller

import (
	"e-learning-system/internal/domain/model"
	"e-learning-system/internal/domain/service"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

type PaymentController struct {
	PaymentService service.PaymentService
}

// CreatePayment implements service.PaymentService.
func (pC *PaymentController) CreatePayment(c * gin.Context) {
	var Payment model.Payment 

	if err := c.BindJSON(&Payment); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	
	createdPayment, err := pC.PaymentService.CreatePayment(Payment.UserID, Payment.Role, Payment.Amount, Payment.Currency, Payment.Method, Payment.Type)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(201, createdPayment)
}

// DeletPayment implements service.PaymentService.
func (pC *PaymentController) DeletPayment(c * gin.Context) {
		// param
	paymentParam := c.Param("id")
	paymentID, err := uuid.FromString(paymentParam)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	// save service
	err = pC.PaymentService.DeletPayment(paymentID)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	// return
	c.JSON(200, "succesfully deleted")
}

// GetAllPayment implements service.PaymentService.
func (pC *PaymentController) GetAllPayment(c * gin.Context)  {
	payment, err := pC.PaymentService.GetAllPayment()
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, payment)
}

// GetPaymentById implements service.PaymentService.
func (pC *PaymentController) GetPaymentById(c * gin.Context) {
	// param
	paymentParam := c.Param("id")
	paymentID, err := uuid.FromString(paymentParam)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
  // save service 
	payment, err := pC.PaymentService.GetPaymentById(paymentID)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	// return
	c.JSON(200, payment)
}

// UpdatePayment implements service.PaymentService.
func (pC *PaymentController) UpdatePayment(c * gin.Context) {
	var payment model.Payment
	// param
	paymentParam := c.Param("id")
	paymentID, err := uuid.FromString(paymentParam)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
  // bind with json
	if err := c.BindJSON(&paymentID); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	payment.ID = paymentID

	// save 
	if err := pC.PaymentService.UpdatePayment(&payment); err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	//retun
	c.JSON(200, "succesfully updated!",)
}

func NewPaymentController(PaymentService service.PaymentService) *PaymentController {
	return &PaymentController{
		PaymentService: PaymentService,
	}

}
