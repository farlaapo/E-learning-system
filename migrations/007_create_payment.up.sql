-- ENUM for user role (optional if you have fixed roles)
CREATE TYPE user_role AS ENUM ('student', 'tutor');

-- ENUM for payment method (local + international)
CREATE TYPE payment_method AS ENUM (
    'Stripe',
    'PayPal',
    'MobileMoney',
    'Card',
    'BankTransfer',
    'Other'
);

-- ENUM for payment type (transaction model)
CREATE TYPE payment_type AS ENUM (
    'ONE_TIME',
    'SUBSCRIPTION'
);
-- ENUM for payment status
CREATE TYPE payment_status AS ENUM (
    'PENDING',
    'SUCCESS',
    'FAILED',
    'REFUNDED'
);

-- Payments table ----
CREATE TABLE IF NOT EXISTS payments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    role user_role,
    amount DECIMAL(10,2) NOT NULL,
    currency VARCHAR(10) NOT NULL,
    method payment_method,
    type payment_type,
    status payment_status,
    transaction_ref VARCHAR(100) UNIQUE,
    description TEXT,
    plan_name VARCHAR(100),
    start_date TIMESTAMP,
    end_date TIMESTAMP,
    renewal_date TIMESTAMP,
    cancelled_at TIMESTAMP,
    provider_ref VARCHAR(100),
    is_recurring BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

----- CREATE PAYMENT ------
CREATE OR REPLACE PROCEDURE  create_payment (
  IN p_id UUID ,
  IN p_ user_id UUID ,
  IN p_role user_role,
  IN p_amount DECIMAL(10,2),
  IN p_currency VARCHAR(10) ,
  IN p_method payment_method,
  IN p_type payment_type,
  IN p_status payment_status,
  IN p_transaction_ref VARCHAR(100) UNIQUE,
  IN p_description TEXT,
  IN p_plan_name VARCHAR(100),
  IN p_start_date TIMESTAMP,
  IN p_end_date TIMESTAMP,
  IN p_renewal_date TIMESTAMP,
  IN p_cancelled_at TIMESTAMP,
  IN p_provider_ref VARCHAR(100),
  IN p_is_recurring BOOLEAN DEFAULT FALSE,
)
LANGUAGE plpgsql AS $$;
BEGIN
   INSERT INTO payments (
      id , user_id , role , amount, currency, method , type, status, transaction_ref, description, plan_name,
      start_date, end_date, renewal_date, cancelled_at, provider_ref,is_recurring
   )VALUES (
    p_id , p_ user_id , p_role , p_amount, p_currency, p_method , p_type , p_status, p_transaction_ref, p_description,  p_plan_name,
      p_start_date, p_end_date, p_renewal_date, p_cancelled_at, p_provider_ref, p_is_recurring
   );
END;
$$;

---- UPDATE PAYMENT ----
CREATE OR REPLACE PROCEDURE update_payment (
  IN p_role user_role,
  IN p_amount DECIMAL(10,2),
  IN p_currency VARCHAR(10) ,
  IN p_method payment_method,
  IN p_type payment_type,
  IN p_status payment_status,
  IN p_transaction_ref VARCHAR(100) UNIQUE,
  IN p_description TEXT,
  IN p_plan_name VARCHAR(100),
  IN p_start_date TIMESTAMP,
  IN p_end_date TIMESTAMP,
  IN p_renewal_date TIMESTAMP,
  IN p_cancelled_at TIMESTAMP,
  IN p_provider_ref VARCHAR(100),
  IN p_is_recurring BOOLEAN DEFAULT FALSE,
) LANGUAGE plpgsql 
AS $$
BEGIN
  UPDATE payments
  SET  role = p_role , amount = p_amount, currency = p_currency , method = p_method , type = p_type , status = p_status, transaction_ref = p_transaction_ref, description = p_description,  plan_name = p_plan_name ,
      start_date = p_start_date, end_date = p_end_date, renewal_date = p_renewal_date, cancelled_at = p_cancelled_at, provider_ref = p_provider_ref, is_recurring = p_is_recurring,
      updated_at = CURRENT_TIMESTAMP
  WHERE id = p_id AND deleted_at IS NULL;

  ----- GET PAYMENT BY ID ----
CREATE OR REPLACE FUNCTION get_payment_id(p_id UUID)
RETURN TABLE (
   p_id UUID ,
   p_ user_id UUID ,
   p_role user_role,
   p_amount DECIMAL(10,2),
   p_currency VARCHAR(10) ,
   p_method payment_method,
   p_type payment_type,
   p_status payment_status,
   p_transaction_ref VARCHAR(100) UNIQUE,
   p_description TEXT,
   p_plan_name VARCHAR(100),
   p_start_date TIMESTAMP,
   p_end_date TIMESTAMP,
   p_renewal_date TIMESTAMP,
   p_cancelled_at TIMESTAMP,
   p_provider_ref VARCHAR(100),
   p_is_recurring BOOLEAN DEFAULT FALSE,
   p_created_at TIMESTAMP ,
   p_updated_at TIMESTAMP 
)
LANGUAGE plpgsql AS $$
BEGIN
   RETURN QUERY
   SELECT
       t_id ,
       t_user_id ,
       t_role ,
       t_amount ,
       t_currency ,
       t_method ,
       t_type ,
       t_status ,
       t_transaction_ref,
       t_description ,
       t_plan_name ,
       t_start_date ,
       t_end_date ,
       t_renewal_date ,
       t_cancelled_at ,
       t_provider_ref ,
       t_is_recurring ,
       t_created_at ,
       t_updated_at
   FROM  payments t
   WHERE t.id = p_id;
END;
$$;

---- GET ALL PAYMENTS ----
CREATE OR REPLACE FUNCTION get_all_payment()
RETURN TABLE (
   p_id UUID ,
   p_ user_id UUID ,
   p_role user_role,
   p_amount DECIMAL(10,2),
   p_currency VARCHAR(10) ,
   p_method payment_method,
   p_type payment_type,
   p_status payment_status,
   p_transaction_ref VARCHAR(100) UNIQUE,
   p_description TEXT,
   p_plan_name VARCHAR(100),
   p_start_date TIMESTAMP,
   p_end_date TIMESTAMP,
   p_renewal_date TIMESTAMP,
   p_cancelled_at TIMESTAMP,
   p_provider_ref VARCHAR(100),
   p_is_recurring BOOLEAN DEFAULT FALSE,
   p_created_at TIMESTAMP ,
   p_updated_at TIMESTAMP 
)
LANGUAGE plpgsql AS $$
BEGIN
   RETURN QUERY
   SELECT
       t_id ,
       t_user_id ,
       t_role ,
       t_amount ,
       t_currency ,
       t_method ,
       t_type ,
       t_status ,
       t_transaction_ref,
       t_description ,
       t_plan_name ,
       t_start_date ,
       t_end_date ,
       t_renewal_date ,
       t_cancelled_at ,
       t_provider_ref ,
       t_is_recurring ,
       t_created_at ,
       t_updated_at
   FROM  payments t
   WHERE t.deleted_at IS NULL; 
END;
$$

----- DELETE PAYMENTS
CREATE OR REPLACE PROCEDURE Delete_payment(
  IN p_id UUID
)
LANGUAGE plpgsql 
AS $$
BEGIN
    DELETE FROM payments WHERE id = p_id;
END;
$$:
   
      
       
   


