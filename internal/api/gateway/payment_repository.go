package gateway

import (
	"database/sql"
	"e-learning-system/internal/domain/model"
	"e-learning-system/internal/domain/repository"
	"fmt"
	"log"

	"github.com/gofrs/uuid"
)

// PaymentRepository struct
type PaymentRepository struct {
	db *sql.DB
}

// Create implements repository.PaymnetRepository.
func (r *PaymentRepository) Create(payment *model.Payment) error {
	_, err := r.db.Exec(`SELECT create_payment($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11 $12, $13, $14, $15, $16 $17)`,
	payment.ID ,    
	payment.UserID ,        
	payment.Role  ,       
	payment.Amount  ,       
	payment.Currency ,     
	payment.Method ,      
	payment.Type  ,    
	payment.Status  ,    
	payment.TransactionRef, 
	payment.Description ,
	payment.PlanName , 
	payment.StartDate ,  
	payment.EndDate ,
	payment.RenewalDate  , 
	payment.CancelledAt ,   
	payment.ProviderRef  , 
	payment.IsRecurring ,
	)  
	if err != nil {
        log.Printf("Error calling create_payment: %v", err)
        return err
    }

    log.Printf("payment created %+v", payment)
    return nil

}

// Delete implements repository.PaymnetRepository.
func (r *PaymentRepository) Delete(paymentID uuid.UUID) error {
	 _, err := r.db.Exec(`SELECT Delete_payment($1)`,  paymentID)
	 if err != nil {
		log.Printf("Error calling Delete_payment for ID %v, %v", err, paymentID)
		return err
	}

	log.Printf("Payment created %+v", paymentID)
	return nil
}

// GetAll implements repository.PaymnetRepository.
func (r *PaymentRepository) GetAll() ([]*model.Payment, error) {
	query  := `SELECT * FROM get_all_payment()`
	row, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer row.Close()
	var payments []*model.Payment
	for row.Next() {
		var payment model.Payment
		err := row.Scan(
			&payment.ID ,    
	   	&payment.UserID ,        
	    &payment.Role  ,       
	    &payment.Amount  ,       
	    &payment.Currency ,     
	    &payment.Method ,      
	    &payment.Type  ,    
	    &payment.Status  ,    
	    &payment.TransactionRef, 
	    &payment.Description ,
	    &payment.PlanName , 
	    &payment.StartDate ,  
	    &payment.EndDate ,
	    &payment.RenewalDate  , 
	    &payment.CancelledAt ,   
	    &payment.ProviderRef  , 
	    &payment.IsRecurring ,
		  &payment.CreatedAt,
		  &payment.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		payments = append(payments, &payment)
	}

	return  payments, nil
}

// GetByID implements repository.PaymnetRepository.
func (r *PaymentRepository) GetByID(paymentID uuid.UUID) (*model.Payment, error) {
	query := `SELECT * FROM  get_payment_id($1)`
	row  := r.db.QueryRow(query, paymentID)

	var payment  model.Payment
	err := row.Scan(
		  &payment.ID ,    
	   	&payment.UserID ,        
	    &payment.Role  ,       
	    &payment.Amount  ,       
	    &payment.Currency ,     
	    &payment.Method ,      
	    &payment.Type  ,    
	    &payment.Status  ,    
	    &payment.TransactionRef, 
	    &payment.Description ,
	    &payment.PlanName , 
	    &payment.StartDate ,  
	    &payment.EndDate ,
	    &payment.RenewalDate  , 
	    &payment.CancelledAt ,   
	    &payment.ProviderRef  , 
	    &payment.IsRecurring ,
		  &payment.CreatedAt,
		  &payment.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("payment not found")
			return nil, fmt.Errorf("payment not found")
		}
		log.Printf("DB error: %v", err)
		return nil, err
	}

	return &payment, nil
}

// Update implements repository.PaymnetRepository.
func (r *PaymentRepository) Update(payment *model.Payment) error {
	_, err := r.db.Exec(`SELECT update_payment(1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11 $12, $13, $14, $15, $16) `,
	   payment.ID ,
	   payment.Role  ,       
	   payment.Amount  ,       
	   payment.Currency ,     
	   payment.Method ,      
	   payment.Type  ,    
	   payment.Status  ,    
	   payment.TransactionRef, 
	   payment.Description ,
	   payment.PlanName , 
  	 payment.StartDate ,  
	   payment.EndDate ,
	   payment.RenewalDate  , 
	   payment.CancelledAt ,   
	   payment.ProviderRef  , 
	   payment.IsRecurring ,
  )
	if err != nil {
        log.Printf("Error calling create_payment: %v", err)
        return err
    }

    log.Printf("payment created %+v", payment)
    return nil

}

func NewPaymentRepository(db *sql.DB) repository.PaymnetRepository {
	return &PaymentRepository{
		db: db,
	}
}
