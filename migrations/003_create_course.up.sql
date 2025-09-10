
-- Course status: draft or published
CREATE TYPE course_status AS ENUM ('draft', 'published');

-- ===================== TABLE =====================
CREATE TABLE IF NOT EXISTS courses (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(255) NOT NULL,
    description TEXT,
    instructor_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- ===================== CREATE =====================
-- CREATE OR REPLACE PROCEDURE create_course(
--     IN p_title VARCHAR,
--     IN p_description TEXT,
--     IN p_instructor_id UUID,
--     OUT p_course_id UUID
-- )
-- LANGUAGE plpgsql
-- AS $$
-- BEGIN
--     p_course_id := gen_random_uuid();
--     INSERT INTO courses (id, title, description, instructor_id, created_at, updated_at)
--     VALUES (p_course_id, p_title, p_description, p_instructor_id, NOW(), NOW());
-- END;
-- $$;

CREATE OR REPLACE FUNCTION create_course(
    p_title VARCHAR,
    p_description text,
    p_instructor_id UUID
)
RETURNS UUID
LANGUAGE plpgsql
AS $$
DECLARE
    new_course_id UUID;
BEGIN
    INSERT INTO courses (id, title, description,  instructor_id)
    VALUES (uuid_generate_v4(), p_title, p_description, p_instructor_id)
    RETURNING id into new_course_id;

    RETURN new_course_id;
END;
$$;


-- 3. Function: create_user
-- CREATE OR REPLACE FUNCTION public.create_user(
--     p_email VARCHAR,
--     p_password TEXT,
--     p_first_name VARCHAR,
--     p_last_name VARCHAR,
--     p_role user_role DEFAULT 'student'
-- )
-- RETURNS UUID
-- LANGUAGE plpgsql
-- AS $$
-- DECLARE
--     new_user_id UUID;
-- BEGIN
--     INSERT INTO users (id, email, password, first_name, last_name, role)
--     VALUES (uuid_generate_v4(), p_email, p_password, p_first_name, p_last_name, p_role)
--     RETURNING id INTO new_user_id;

--     RETURN new_user_id;
-- END;
-- $$;

-- ===================== UPDATE =====================
CREATE OR REPLACE PROCEDURE update_course(
    IN p_id UUID,
    IN p_title VARCHAR,
    IN p_description TEXT,
    IN p_instructor_id UUID
)
LANGUAGE plpgsql
AS $$
BEGIN
    UPDATE courses
    SET title = p_title,
        description = p_description,
        instructor_id = p_instructor_id,
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
    title VARCHAR,
    description TEXT,
    instructor_id UUID,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
) AS $$
BEGIN
    RETURN QUERY
    SELECT id, title, description, instructor_id, created_at, updated_at
    FROM courses
    WHERE id = p_id;
END;
$$ LANGUAGE plpgsql;

-- ===================== GET ALL =====================
CREATE OR REPLACE FUNCTION get_all_courses()
RETURNS TABLE (
    id UUID,
    title VARCHAR,
    description TEXT,
    instructor_id UUID,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
) AS $$
BEGIN
    RETURN QUERY
    SELECT id, title, description, instructor_id, created_at, updated_at
    FROM courses
    ORDER BY created_at DESC;
END;
$$ LANGUAGE plpgsql;

-- ===================== GET BY INSTRUCTOR =====================
CREATE OR REPLACE FUNCTION get_courses_by_instructor(p_instructor_id UUID)
RETURNS TABLE (
    id UUID,
    title VARCHAR,
    description TEXT,
    instructor_id UUID,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
) AS $$
BEGIN
    RETURN QUERY
    SELECT id, title, description, instructor_id, created_at, updated_at
    FROM courses
    WHERE instructor_id = p_instructor_id
    ORDER BY created_at DESC;
END;
$$ LANGUAGE plpgsql;
