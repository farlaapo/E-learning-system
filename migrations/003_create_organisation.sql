-- Enable extensions
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

--- ENUM ORGANIZATION
CREATE TYPE organization_status AS ENUM ('pending', 'approved', 'suspended');

ALTER TABLE organizations
    ALTER COLUMN status TYPE organization_status USING status::organization_status,
    ALTER COLUMN status SET DEFAULT 'pending';


-- organization table

CREATE TABLE IF NOT EXISTS organizations {
    id UUID   PRIMARY KEY DEFAULT uuid_generate_v4(),
    name    VARCHAR(255) NOT NULL,
    description TEXT,
    owner_id  UUID NOT NULL,
    tutors UUID[] DEFAULT '{}',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    
}


-- Create Organization

CREATE OR REPLACE FUNCTION create_organization(
    p_name VARCHAR,
    p_description TEXT,
    p_owner_id UUID
) RETURNS UUID AS $$
DECLARE
    new_id UUID := uuid_generate_v4();
BEGIN
    INSERT INTO organizations (id, name, description, owner_id, created_at, updated_at)
    VALUES (new_id, p_name, p_description, p_owner_id, NOW(), NOW());

    RETURN new_id;
END;
$$ LANGUAGE plpgsql;


-- Get Organization by ID

CREATE OR REPLACE FUNCTION get_organization_by_id(
    p_id UUID
) RETURNS TABLE (
    id UUID,
    name VARCHAR,
    description TEXT,
    owner_id UUID,
    tutors UUID[],
    created_at TIMESTAMP,
    updated_at TIMESTAMP
) AS $$
BEGIN
    RETURN QUERY
    SELECT o.id, o.name, o.description, o.owner_id, o.tutors, o.created_at, o.updated_at
    FROM organizations o
    WHERE o.id = p_id;
END;
$$ LANGUAGE plpgsql;

-- Update Organization

CREATE OR REPLACE FUNCTION update_organization(
    p_id UUID,
    p_name VARCHAR,
    p_description TEXT
) RETURNS VOID AS $$
BEGIN
    UPDATE organizations
    SET name = COALESCE(p_name, name),
        description = COALESCE(p_description, description),
        updated_at = NOW()
    WHERE id = p_id;
END;
$$ LANGUAGE plpgsql;


-- Delete Organization

CREATE OR REPLACE FUNCTION delete_organization(
    p_id UUID
) RETURNS VOID AS $$
BEGIN
    DELETE FROM organizations WHERE id = p_id;
END;
$$ LANGUAGE plpgsql;


