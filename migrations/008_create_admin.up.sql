-- ===========================================
-- ============ ENUM DEFINITIONS =============
-- ===========================================
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'entity_type_enum') THEN
        CREATE TYPE entity_type_enum AS ENUM ('user', 'organization', 'course', 'tutor');
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'approval_status_enum') THEN
        CREATE TYPE approval_status_enum AS ENUM ('pending', 'approved', 'rejected');
    END IF;
END
$$;


-- ===========================================
-- 1 ADMIN DASHBOARD
-- ===========================================
-- Summary view for key platform metrics
CREATE OR REPLACE FUNCTION get_admin_dashboard_summary()
RETURNS TABLE (
    total_users BIGINT,
    total_courses BIGINT,
    total_revenue NUMERIC,
    active_courses BIGINT
)
LANGUAGE plpgsql AS $$
BEGIN
    RETURN QUERY
    SELECT
        (SELECT COUNT(*) FROM users WHERE deleted_at IS NULL) AS total_users,
        (SELECT COUNT(*) FROM courses WHERE deleted_at IS NULL) AS total_courses,
        (SELECT COALESCE(SUM(amount), 0) FROM payments WHERE status = 'completed') AS total_revenue,
        (SELECT COUNT(*) FROM courses WHERE status = 'published') AS active_courses;
END;
$$;


-- ===========================================
-- 2 MANAGED ENTITIES
-- ===========================================
-- Entities managed by admin (users, organizations, courses)
CREATE TABLE IF NOT EXISTS managed_entities (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    type entity_type_enum NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Create
CREATE OR REPLACE PROCEDURE create_managed_entity(
    IN p_id UUID,
    IN p_name VARCHAR,
    IN p_type entity_type_enum,
    IN p_status VARCHAR
)
LANGUAGE plpgsql AS $$
BEGIN
    INSERT INTO managed_entities (id, name, type, status)
    VALUES (p_id, p_name, p_type, p_status);
END;
$$;

-- Update
CREATE OR REPLACE PROCEDURE update_managed_entity(
    IN p_id UUID,
    IN p_name VARCHAR,
    IN p_status VARCHAR
)
LANGUAGE plpgsql AS $$
BEGIN
    UPDATE managed_entities
    SET name = p_name,
        status = p_status
    WHERE id = p_id;
END;
$$;

-- Delete
CREATE OR REPLACE PROCEDURE delete_managed_entity(IN p_id UUID)
LANGUAGE plpgsql AS $$
BEGIN
    DELETE FROM managed_entities WHERE id = p_id;
END;
$$;

-- Get All
CREATE OR REPLACE FUNCTION get_all_managed_entities()
RETURNS TABLE (
    id UUID,
    name VARCHAR,
    type entity_type_enum,
    status VARCHAR,
    created_at TIMESTAMP
)
LANGUAGE plpgsql AS $$
BEGIN
    RETURN QUERY
    SELECT id, name, type, status, created_at
    FROM managed_entities
    ORDER BY created_at DESC;
END;
$$;


-- ===========================================
-- 3 APPROVAL REQUESTS
-- ===========================================
-- Tutor or Organization approval management
CREATE TABLE IF NOT EXISTS approval_requests (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    entity_type entity_type_enum NOT NULL,
    entity_id UUID NOT NULL,
    request_date TIMESTAMP NOT NULL DEFAULT NOW(),
    status approval_status_enum NOT NULL DEFAULT 'pending',
    reviewed_by UUID,
    reviewed_at TIMESTAMP
);

-- Create
CREATE OR REPLACE PROCEDURE create_approval_request(
    IN p_id UUID,
    IN p_entity_type entity_type_enum,
    IN p_entity_id UUID
)
LANGUAGE plpgsql AS $$
BEGIN
    INSERT INTO approval_requests (id, entity_type, entity_id, status)
    VALUES (p_id, p_entity_type, p_entity_id, 'pending');
END;
$$;

-- Update (Approve/Reject)
CREATE OR REPLACE PROCEDURE update_approval_status(
    IN p_id UUID,
    IN p_status approval_status_enum,
    IN p_reviewed_by UUID
)
LANGUAGE plpgsql AS $$
BEGIN
    UPDATE approval_requests
    SET status = p_status,
        reviewed_by = p_reviewed_by,
        reviewed_at = NOW()
    WHERE id = p_id;
END;
$$;

-- Get All Pending
CREATE OR REPLACE FUNCTION get_pending_approvals()
RETURNS TABLE (
    id UUID,
    entity_type entity_type_enum,
    entity_id UUID,
    request_date TIMESTAMP,
    status approval_status_enum
)
LANGUAGE plpgsql AS $$
BEGIN
    RETURN QUERY
    SELECT id, entity_type, entity_id, request_date, status
    FROM approval_requests
    WHERE status = 'pending'
    ORDER BY request_date DESC;
END;
$$;

-- Get All Requests
CREATE OR REPLACE FUNCTION get_all_approval_requests()
RETURNS TABLE (
    id UUID,
    entity_type entity_type_enum,
    entity_id UUID,
    request_date TIMESTAMP,
    status approval_status_enum,
    reviewed_by UUID,
    reviewed_at TIMESTAMP
)
LANGUAGE plpgsql AS $$
BEGIN
    RETURN QUERY
    SELECT id, entity_type, entity_id, request_date, status, reviewed_by, reviewed_at
    FROM approval_requests
    ORDER BY request_date DESC;
END;
$$;


-- ===========================================
--  4 SYSTEM SETTINGS
-- ===========================================
-- Global configuration for system-wide settings
CREATE TABLE IF NOT EXISTS system_settings (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    payment_gateway VARCHAR(100),
    theme VARCHAR(100),
   -- email_template_id VARCHAR(100),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Insert or Update (Upsert)
CREATE OR REPLACE PROCEDURE upsert_system_settings(
    IN p_id UUID,
    IN p_payment_gateway VARCHAR,
    IN p_theme VARCHAR,
  --  IN p_email_template_id VARCHAR
)
LANGUAGE plpgsql AS $$
BEGIN
    IF EXISTS (SELECT 1 FROM system_settings WHERE id = p_id) THEN
        UPDATE system_settings
        SET payment_gateway = p_payment_gateway,
            theme = p_theme,
           -- email_template_id = p_email_template_id,
            updated_at = NOW()
        WHERE id = p_id;
    ELSE
        INSERT INTO system_settings (id, payment_gateway, theme, email_template_id)
        VALUES (p_id, p_payment_gateway, p_theme, /*p_email_template_id*/ NULL);
    END IF;
END;
$$;

-- Get Latest Settings
CREATE OR REPLACE FUNCTION get_latest_system_settings()
RETURNS TABLE (
    id UUID,
    payment_gateway VARCHAR,
    theme VARCHAR,
   -- email_template_id VARCHAR,
    updated_at TIMESTAMP
)
LANGUAGE plpgsql AS $$
BEGIN
    RETURN QUERY
    SELECT id, payment_gateway, theme, updated_at
    FROM system_settings
    ORDER BY updated_at DESC
    LIMIT 1;
END;
$$;
