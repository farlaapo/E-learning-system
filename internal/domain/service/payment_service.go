package service

import (
	"e-learning-system/internal/domain/model"
	"e-learning-system/internal/domain/repository"
	"fmt"
	"log"
	"time"

	"github.com/gofrs/uuid"
)

type PaymentService interface {
	CreatePayment(UserID uuid.UUID, Role string, Amount float64, Currency string, Method string, Type string) (*model.Payment, error)
	UpdatePayment(payment *model.Payment) error
	DeletPayment(paymentID uuid.UUID) error
	GetPaymentById(paymentID uuid.UUID) (*model.Payment, error)
	GetAllPayment() ([]*model.Payment, error)
}

type paymentService struct {
	repo  repository.PaymnetRepository
	token repository.TokenRepository
}

// CreatePayment implements PaymentService.
func (s *paymentService) CreatePayment(UserID uuid.UUID, Role string, Amount float64, Currency string, Method string, Type string) (*model.Payment, error) {
  // generete UUID
	neoPayment, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	newPayment := &model.Payment{
		ID: neoPayment,
		UserID: UserID,
		Role: Role,
		Amount: Amount,
		Currency: Currency,
		Method: Method,
		Type: Type,
		CreatedAt: time.Now(),
	}
	// log the new enrollment
    log.Printf("payment created %+v", neoPayment)

    // save to repository
    err = s.repo.Create(newPayment)
    if err != nil {
        return nil, fmt.Errorf("failed to create payment: %w", err)
    }

    return newPayment, nil

}

// DeletPayment implements PaymentService.
func (s *paymentService) DeletPayment(paymentID uuid.UUID) error {
	_, err := s.repo.GetByID(paymentID)
	if err != nil {
		return  fmt.Errorf("could not find payment with ID %s: %v", paymentID, err)
	}

	if err := s.repo.Delete(paymentID); err != nil {
		return  fmt.Errorf("failed to delete payment with ID %s: %v", paymentID, err)
	}

	log.Printf("successfully deleted payment with id %v", paymentID)
  return  nil
}

// GetAllPayment implements PaymentService.
func (s *paymentService) GetAllPayment() ([]*model.Payment, error) {
	payment, err := s.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get  all payment  ID  %v", payment)
	}

	return  payment, nil
}

// GetPaymentById implements PaymentService.
func (s *paymentService) GetPaymentById(paymentID uuid.UUID) (*model.Payment, error) {
	payment, err := s.repo.GetByID(paymentID)
	if err != nil {
		return nil,  fmt.Errorf("could not find payment with ID %s: %v", paymentID, err)
	}

	return  payment, nil
}

// UpdatePayment implements PaymentService.
func (s *paymentService) UpdatePayment(payment *model.Payment) error {
   _, err := s.repo.GetByID(payment.ID)
	if err != nil {
		return fmt.Errorf("failed to find payment with ID %s", payment.ID)
	}

	if err := s.repo.Update(payment); err != nil {
		return fmt.Errorf("failed to payment with ID %s: %v", payment.ID, err)
	}

	return nil
}

func NewpaymentService(paymentRepo repository.PaymnetRepository, tokenRepo repository.TokenRepository) PaymentService {
	return &paymentService{
		repo:  paymentRepo,
		token: tokenRepo,
	}
}
