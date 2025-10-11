----- ENUM ORGANIZATION -------
CREATE TYPE org_status AS ENUM ('active', 'inactive');

-- organizations table
CREATE TABLE IF NOT EXISTS organizations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    owner_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    tutors UUID[] DEFAULT '{}',         -- array of UUIDs for tutor IDs
    status org_status NOT NULL DEFAULT 'active',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL           -- added for soft delete
);

------ CREATE ORG  ---------
CREATE OR REPLACE PROCEDURE create_orgs(
  IN p_id UUID,
  IN p_name VARCHAR(255),
  IN p_description TEXT,
  IN p_owner_id UUID,
  IN p_tutors UUID[]
)
LANGUAGE plpgsql AS $$
BEGIN
   INSERT INTO organizations (
    id, name, description, owner_id, tutors
   ) VALUES (
    p_id, p_name, p_description, p_owner_id, p_tutors
   );
END;
$$;

----- UPDATE ORG --------
CREATE OR REPLACE PROCEDURE update_org(
  IN p_id UUID,
  IN p_name VARCHAR(255),
  IN p_description TEXT
)
LANGUAGE plpgsql AS $$
BEGIN 
    UPDATE organizations 
    SET name = p_name, 
        description = p_description, 
        updated_at = CURRENT_TIMESTAMP
    WHERE id = p_id AND deleted_at IS NULL;
END;
$$;

---- GET ORGANIZATION BY ID ----
CREATE OR REPLACE FUNCTION get_org_by_id(p_id UUID)
RETURNS TABLE (
  id UUID,
  name VARCHAR(255),
  description TEXT,
  owner_id UUID,
  tutors UUID[],
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
)
LANGUAGE plpgsql AS $$
BEGIN
    RETURN QUERY
    SELECT 
        o.id,
        o.name,
        o.description,
        o.owner_id,
        o.tutors,
        o.created_at,
        o.updated_at,
        o.deleted_at
    FROM organizations o
    WHERE o.id = p_id;
END;
$$;

------ GET ALL ORGS ------
CREATE OR REPLACE FUNCTION get_all_orgs()
RETURNS TABLE (
   id UUID,
   name VARCHAR(255),
   description TEXT,
   owner_id UUID,
   tutors UUID[],
   created_at TIMESTAMP,
   updated_at TIMESTAMP,
   deleted_at TIMESTAMP
)
LANGUAGE plpgsql AS $$
BEGIN
    RETURN QUERY
    SELECT 
        o.id,
        o.name,
        o.description,
        o.owner_id,
        o.tutors,
        o.created_at,
        o.updated_at,
        o.deleted_at
    FROM organizations o
    WHERE o.deleted_at IS NULL;  -- exclude soft-deleted
END;
$$;

----------- DELETE ORGS ----------
CREATE OR REPLACE PROCEDURE delete_orgs(
  IN p_id UUID
)
LANGUAGE plpgsql AS $$
BEGIN
    DELETE FROM organizations WHERE id = p_id;
END;
$$;
