-- ENUMs
CREATE TYPE payment_role AS ENUM ('student', 'tutor');
CREATE TYPE payment_method AS ENUM ('Stripe', 'PayPal', 'MobileMoney', 'Card', 'BankTransfer', 'Other');
CREATE TYPE payment_type AS ENUM ('ONE_TIME', 'SUBSCRIPTION');
CREATE TYPE payment_status AS ENUM ('PENDING', 'SUCCESS', 'FAILED', 'REFUNDED');

-- Payments table
CREATE TABLE IF NOT EXISTS payments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    role payment_role,
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
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

-- CREATE PAYMENT
CREATE OR REPLACE PROCEDURE create_payment (
  IN p_id UUID,
  IN p_user_id UUID,
  IN p_role payment_role,
  IN p_amount DECIMAL(10,2),
  IN p_currency VARCHAR(10),
  IN p_method payment_method,
  IN p_type payment_type,
  IN p_status payment_status,
  IN p_transaction_ref VARCHAR(100),
  IN p_description TEXT,
  IN p_plan_name VARCHAR(100),
  IN p_start_date TIMESTAMP,
  IN p_end_date TIMESTAMP,
  IN p_renewal_date TIMESTAMP,
  IN p_cancelled_at TIMESTAMP,
  IN p_provider_ref VARCHAR(100),
  IN p_is_recurring BOOLEAN
)
LANGUAGE plpgsql AS $$
BEGIN
    INSERT INTO payments (
        id, user_id, role, amount, currency, method, type, status,
        transaction_ref, description, plan_name,
        start_date, end_date, renewal_date, cancelled_at, provider_ref, is_recurring
    ) VALUES (
        p_id, p_user_id, p_role, p_amount, p_currency, p_method, p_type, p_status,
        p_transaction_ref, p_description, p_plan_name,
        p_start_date, p_end_date, p_renewal_date, p_cancelled_at, p_provider_ref, p_is_recurring
    );
END;
$$;

-- UPDATE PAYMENT
CREATE OR REPLACE PROCEDURE update_payment (
  IN p_id UUID,
  IN p_role payment_role,
  IN p_amount DECIMAL(10,2),
  IN p_currency VARCHAR(10),
  IN p_method payment_method,
  IN p_type payment_type,
  IN p_status payment_status,
  IN p_transaction_ref VARCHAR(100),
  IN p_description TEXT,
  IN p_plan_name VARCHAR(100),
  IN p_start_date TIMESTAMP,
  IN p_end_date TIMESTAMP,
  IN p_renewal_date TIMESTAMP,
  IN p_cancelled_at TIMESTAMP,
  IN p_provider_ref VARCHAR(100),
  IN p_is_recurring BOOLEAN
 
)
LANGUAGE plpgsql AS $$
BEGIN
    UPDATE payments
    SET role = p_role,
        amount = p_amount,
        currency = p_currency,
        method = p_method,
        type = p_type,
        status = p_status,
        transaction_ref = p_transaction_ref,
        description = p_description,
        plan_name = p_plan_name,
        start_date = p_start_date,
        end_date = p_end_date,
        renewal_date = p_renewal_date,
        cancelled_at = p_cancelled_at,
        provider_ref = p_provider_ref,
        is_recurring = p_is_recurring,
        updated_at = CURRENT_TIMESTAMP
    WHERE id = p_id;
END;
$$;

-- GET PAYMENT BY ID
CREATE OR REPLACE FUNCTION get_payment_by_id(P_id UUID)
RETURNS TABLE (
   id UUID,
   user_id UUID,
   role payment_role,
   amount DECIMAL(10,2),
   currency VARCHAR(10),
   method payment_method,
   type payment_type,
   status payment_status,
   transaction_ref VARCHAR(100),
   description TEXT,
   plan_name VARCHAR(100),
   start_date TIMESTAMP,
   end_date TIMESTAMP,
   renewal_date TIMESTAMP,
   cancelled_at TIMESTAMP,
   provider_ref VARCHAR(100),
   is_recurring BOOLEAN,
   created_at TIMESTAMP,
   updated_at TIMESTAMP,
   deleted_at TIMESTAMP
)
LANGUAGE plpgsql AS $$
BEGIN
   RETURN QUERY
   SELECT 
       y.id ,
       y.user_id ,
       y.role ,
       y.amount,
       y.currency ,
       y.method ,
       y.type ,
       y.status ,
       y.transaction_ref ,
       y.description ,
       y.plan_name ,
       y.start_date ,
       y.end_date ,
       y.renewal_date ,
       y.cancelled_at ,
       y.provider_ref ,
       y.is_recurring ,
       y.created_at ,
       y.updated_at ,
       y.deleted_at 
         
   FROM payments  y
   WHERE y.id = p_id;
END;
$$;

-- GET ALL PAYMENTS
CREATE OR REPLACE FUNCTION get_all_payments()
RETURNS TABLE (
   id UUID,
   user_id UUID,
   role payment_role,
   amount DECIMAL(10,2),
   currency VARCHAR(10),
   method payment_method,
   type payment_type,
   status payment_status,
   transaction_ref VARCHAR(100),
   description TEXT,
   plan_name VARCHAR(100),
   start_date TIMESTAMP,
   end_date TIMESTAMP,
   renewal_date TIMESTAMP,
   cancelled_at TIMESTAMP,
   provider_ref VARCHAR(100),
   is_recurring BOOLEAN,
   created_at TIMESTAMP,
   updated_at TIMESTAMP,
   deleted_at TIMESTAMP
)
LANGUAGE plpgsql AS $$
BEGIN
   RETURN QUERY
   SELECT 
       y.id ,
       y.user_id ,
       y.role ,
       y.amount,
       y.currency ,
       y.method ,
       y.type ,
       y.status ,
       y.transaction_ref ,
       y.description ,
       y.plan_name ,
       y.start_date ,
       y.end_date ,
       y.renewal_date ,
       y.cancelled_at ,
       y.provider_ref ,
       y.is_recurring ,
       y.created_at ,
       y.updated_at ,
       y.deleted_at 
   FROM payments  y
   WHERE y.deleted_at IS NULL;
END;
$$;

-- DELETE PAYMENT
CREATE OR REPLACE PROCEDURE delete_payment(IN p_id UUID)
LANGUAGE plpgsql AS $$
BEGIN
    DELETE FROM payments WHERE id = p_id;
END;
$$;
