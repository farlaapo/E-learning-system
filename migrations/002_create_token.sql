-- ===============================
-- Drop existing objects
-- ===============================
DROP PROCEDURE IF EXISTS create_token(UUID, UUID, TEXT, TIMESTAMP, TIMESTAMP, TIMESTAMP);
DROP FUNCTION IF EXISTS get_token_by_token(TEXT);
DROP TABLE IF EXISTS tokens CASCADE;

-- ===============================
-- Enable UUID extensions
-- ===============================
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- ===============================
-- Create tokens table
-- ===============================
CREATE TABLE IF NOT EXISTS tokens (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    token TEXT NOT NULL UNIQUE,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

-- ===============================
-- Create procedure: create_token
-- ===============================
CREATE OR REPLACE PROCEDURE create_token(
    IN p_id UUID,
    IN p_user_id UUID,
    IN p_token TEXT,
    IN p_expires_at TIMESTAMP,
    IN p_created_at TIMESTAMP,
    IN p_updated_at TIMESTAMP
)
LANGUAGE plpgsql
AS $$
BEGIN
    INSERT INTO tokens (
        id, user_id, token, expires_at, created_at, updated_at
    ) VALUES (
        p_id, p_user_id, p_token, p_expires_at, NOW(), NOW()
    );
END;
$$;

-- ===============================
-- Create function: get_token_by_token
-- Only returns valid, non-expired, non-deleted tokens
-- ===============================
CREATE OR REPLACE FUNCTION get_token_by_token(p_token TEXT)
RETURNS TABLE (
    id UUID,
    user_id UUID,
    token TEXT,
    expires_at TIMESTAMP,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
)
LANGUAGE plpgsql
AS $$
BEGIN
    RETURN QUERY
    SELECT 
        t.id,
        t.user_id,
        t.token,
        t.expires_at,
        t.created_at,
        t.updated_at,
        t.deleted_at
    FROM tokens t
    WHERE 
        t.token = p_token
        AND t.deleted_at IS NULL
        AND t.expires_at > NOW();
END;
$$;
