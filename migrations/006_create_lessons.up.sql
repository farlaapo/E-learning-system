---- ENUM ----
CREATE TYPE lesson_status AS ENUM ('video', 'text', 'quiz');

------ CREATE LESSON TABLE
CREATE TABLE IF NOT EXISTS lessons (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    module_id UUID NOT NULL REFERENCES modules(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    content TEXT,
    video_url TEXT[] DEFAULT '{}',   -- maps to []string in Go
    order_num INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

---- CREATE LESSON -----
CREATE OR REPLACE PROCEDURE create_lesson(
  IN p_id UUID,
  IN p_module_id UUID,
  IN p_title VARCHAR(255),
  IN p_content TEXT,
  IN p_video_url TEXT[],
  IN p_order_num INT
)
LANGUAGE plpgsql AS $$
BEGIN
   INSERT INTO lessons(
    id, module_id, title, content, video_url, order_num)
   VALUES (
    p_id, p_module_id, p_title, p_content, p_video_url, p_order_num
    );
END;
$$;

----- UPDATE LESSON --------
CREATE OR REPLACE PROCEDURE update_lesson(
  IN p_id UUID,
  IN p_title VARCHAR(255),
  IN p_content TEXT,
  IN p_video_url TEXT[],
  IN p_order_num INT
)
LANGUAGE plpgsql AS $$
BEGIN
   UPDATE lessons
   SET title = p_title,
       content = p_content,
       video_url = p_video_url,
       order_num = p_order_num,
       updated_at = CURRENT_TIMESTAMP
   WHERE id = p_id AND deleted_at IS NULL;
END;
$$;

---- GET LESSON BY ID ----
CREATE OR REPLACE FUNCTION get_lesson_by_id(p_id UUID)
RETURNS TABLE ( 
    id UUID,
    module_id UUID,
    title VARCHAR(255),
    content TEXT,
    video_url TEXT[],
    order_num INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
) LANGUAGE plpgsql AS $$
BEGIN
    RETURN QUERY
    SELECT
        l.id,
        l.module_id,
        l.title,
        l.content,
        l.video_url,
        l.order_num,
        l.created_at,
        l.updated_at,
        l.deleted_at
    FROM lessons l
    WHERE l.id = p_id;
END;
$$;

------- GET ALL LESSON -------
CREATE OR REPLACE FUNCTION get_all_lesson()
RETURNS TABLE ( 
    id UUID,
    module_id UUID,
    title VARCHAR(255),
    content TEXT,
    video_url TEXT[],
    order_num INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
) LANGUAGE plpgsql AS $$
BEGIN
    RETURN QUERY
    SELECT
        l.id,
        l.module_id,
        l.title,
        l.content,
        l.video_url,
        l.order_num,
        l.created_at,
        l.updated_at,
        l.deleted_at
    FROM lessons l
    WHERE l.deleted_at IS NULL;
END;
$$;

---- DELETE LESSON ------
CREATE OR REPLACE PROCEDURE delete_lesson(
  IN p_id UUID
)
LANGUAGE plpgsql AS $$
BEGIN
    DELETE FROM lessons WHERE id = p_id;
END;
$$;
