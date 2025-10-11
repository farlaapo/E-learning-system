package routes

import (
	"e-learning-system/internal/api/controller"
	"e-learning-system/internal/api/middleware"
	"e-learning-system/internal/domain/repository"

	"github.com/gin-gonic/gin"
)

// RegistePaymentRoutes registers payment-related routes
func RegisterpaymentRoutes(router *gin.Engine, PaymentController controller.PaymentController, tokenRepo repository.TokenRepository) {
  
	// Auth-middleware
	authMidlleware := middleware.AuthMiddleware(tokenRepo)
 
	paymentGroup := router.Group("/payments")
	{
     paymentGroup.Use(authMidlleware)
		 {
			paymentGroup.POST("", PaymentController.CreatePayment)
			paymentGroup.GET("/", PaymentController.GetAllPayment)
			paymentGroup.GET("/:id", PaymentController.GetPaymentById)
			paymentGroup.PUT("/:id", PaymentController.UpdatePayment)
			paymentGroup.DELETE("/:id", PaymentController.DeletPayment)
		 }
	}

}