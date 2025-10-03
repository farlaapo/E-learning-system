
-- Course status: draft or published
CREATE TYPE course_status AS ENUM ('draft', 'published');

-- ===================== TABLE =====================
CREATE TABLE IF NOT EXISTS courses (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title VARCHAR(255) NOT NULL,
    description TEXT,
    instructor_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    category VARCHAR(100),
    tags  TEXT[],
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);



-------- create_course -------------

CREATE OR REPLACE PROCEDURE  create_course(
    IN P_id UUID,
    IN p_title VARCHAR(255),
    IN p_description TEXT,
    IN p_instructor_id UUID,
    IN P_category  VARCHAR(100),
    IN P_tags  TEXT[]

)
LANGUAGE plpgsql AS $$
BEGIN   
    INSERT INTO courses (
        id, title, description, instructor_id, category, tags
    )VALUES (
        P_id, p_title, p_description, p_instructor_id, P_category, P_tags
    );
END;
$$;





-- ===================== UPDATE =====================
CREATE OR REPLACE PROCEDURE update_course(
    IN p_id UUID,
    IN p_title VARCHAR,
    IN p_description TEXT,
    IN p_instructor_id UUID,
    IN P_category  VARCHAR(100),
    IN P_tags  TEXT[]
)
LANGUAGE plpgsql
AS $$
BEGIN
    UPDATE courses
    SET title = p_title,
        description = p_description,
        instructor_id = p_instructor_id, category = P_category, tags = P_tags,
        updated_at = NOW()
    WHERE id = p_id;
    
END;
$$;

-- ===================== DELETE =====================
CREATE OR REPLACE PROCEDURE delete_course(
    IN p_id UUID
)
LANGUAGE plpgsql
AS $$
BEGIN
    DELETE FROM courses WHERE id = p_id;
END;
$$;

-- ===================== GET BY ID =====================
CREATE OR REPLACE FUNCTION get_course_by_id(p_id UUID)
RETURNS TABLE (
    id UUID,
    title VARCHAR(255),
    description TEXT,
    instructor_id UUID,
    category VARCHAR(100),
    tags TEXT[],
    created_at TIMESTAMP,
    updated_at TIMESTAMP
) 
LANGUAGE plpgsql AS $$
BEGIN
    RETURN QUERY
    SELECT
        c.id,
        c.title,
        c.description, 
        c.instructor_id,
        c.category,
        c.tags,
        c.created_at,
        c.updated_at
    FROM courses c
    WHERE c.id = p_id;
    END;
    $$;

-------- GET ALL --------
CREATE OR REPLACE FUNCTION get_all_courses()
RETURNS TABLE (
    id UUID,
    title VARCHAR,
    description TEXT,
    instructor_id UUID,
    category VARCHAR(100),
    tags TEXT[],
    created_at TIMESTAMP,
    updated_at TIMESTAMP
) 
LANGUAGE plpgsql AS $$
BEGIN
    RETURN QUERY
    SELECT c.id, c.title, c.description, c.instructor_id,   c.category, c.tags, c.created_at, c.updated_at
    FROM courses c
    WHERE c.deleted_at IS NULL;
END;
$$;


-- ===================== GET BY INSTRUCTOR =====================
CREATE OR REPLACE FUNCTION get_courses_by_instructor(p_instructor_id UUID)
RETURNS TABLE (
    id UUID,
    title VARCHAR,
    description TEXT,
    instructor_id UUID,
    category VARCHAR(100),
    tags TEXT[],
    created_at TIMESTAMP,
    updated_at TIMESTAMP
) AS $$
BEGIN
    RETURN QUERY
    SELECT c.id, c.title, c.description, c.instructor_id, c.category, c.tags, c.created_at, c.updated_at
    FROM courses c
    WHERE c.id =  p_instructor_id;
END;
$$ LANGUAGE plpgsql;
