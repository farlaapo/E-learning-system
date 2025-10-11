--- ENUM ----
CREATE TYPE module_status AS ENUM ('video', 'quiz', 'assignment');

---- MODULE TABLE ---
CREATE TABLE IF NOT EXISTS modules(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  course_id UUID NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
  title  VARCHAR(255) NOT NULL,
  order_num INT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP NULL
);

----- CREATE MODULE ----
CREATE OR REPLACE PROCEDURE create_module (
  IN p_id UUID,
  IN p_course_id UUID,
  IN p_title VARCHAR(255),
  IN p_order_num INT
)
LANGUAGE plpgsql AS $$
BEGIN
   INSERT INTO modules (
    id, course_id, title, order_num
   ) VALUES (
     p_id,  p_course_id,  p_title, p_order_num
   );
END;
$$;

----- UPDATE MODULE ------
CREATE OR REPLACE PROCEDURE update_module(
  IN p_id UUID,
  IN p_title VARCHAR(255),
  IN p_order_num INT
)
LANGUAGE plpgsql AS $$
BEGIN
   UPDATE modules
   SET  title = p_title,
        order_num = p_order_num,
        updated_at = CURRENT_TIMESTAMP
   WHERE id = p_id AND deleted_at IS NULL;
END;
$$;

---- GET MODULE BY ID -------
CREATE OR REPLACE FUNCTION get_module_by_id(p_id UUID)
RETURNS TABLE (
    id UUID,
    course_id UUID,
    title VARCHAR(255),
    order_num INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
)
LANGUAGE plpgsql AS $$
BEGIN
    RETURN QUERY
    SELECT
        m.id, 
        m.course_id,
        m.title,
        m.order_num,
        m.created_at,
        m.updated_at,
        m.deleted_at
    FROM modules m
    WHERE m.id = p_id;
END;
$$;

------ GET ALL MODULE ------
CREATE OR REPLACE FUNCTION get_all_module()
RETURNS TABLE (
    id UUID,
    course_id UUID,
    title VARCHAR(255),
    order_num INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
)
LANGUAGE plpgsql AS $$
BEGIN
    RETURN QUERY
    SELECT
        m.id, 
        m.course_id,
        m.title,
        m.order_num,
        m.created_at,
        m.updated_at,
        m.deleted_at
    FROM modules m
    WHERE m.deleted_at IS NULL;
END;
$$;

------ DELETE MODULE -------
CREATE OR REPLACE PROCEDURE delete_module (
  IN p_id UUID
)
LANGUAGE plpgsql AS $$
BEGIN
    DELETE FROM modules WHERE id = p_id;
END;
$$;
