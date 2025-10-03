-- CREATE ENUM type for enrollment status if not exists  --

   CREATE TYPE enrollment_status AS ENUM ('pending', 'active', 'completed', 'canceled');
   

-- Create enrollments table if not exists --

CREATE TABLE IF NOT EXISTS enrollments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    course_id UUID NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    enrollment_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    completed BOOLEAN NOT NULL DEFAULT FALSE,
    certificate_issued_at TIMESTAMP NULL,
    certificate_template TEXT DEFAULT 'default-template-v1',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);
-- --------- ffghjkl;
-- CREATE OR REPLACE PROCEDURE  create_course(
--     IN P_id UUID,
--     IN p_title VARCHAR(255),
--     IN p_description TEXT,
--     IN p_instructor_id UUID,
--     IN P_category  VARCHAR(100),
--     IN P_tags  TEXT[]

-- )
-- LANGUAGE plpgsql AS $$
-- BEGIN   
--     INSERT INTO courses (
--         id, title, description, instructor_id, category, tags
--     )VALUES (
--         P_id, p_title, p_description, p_instructor_id, P_category, P_tags
--     );
-- END;
-- $$;


--- Create enrollment ---
CREATE OR REPLACE PROCEDURE create_enrollment(
  IN p_id UUID,
  IN p_course_id UUID,
  IN p_user_id UUID,
  IN p_completed BOOLEAN,
  IN p_certificate_issued_at TIMESTAMP,
  IN p_certificate_template TEXT
)
LANGUAGE plpgsql AS $$
BEGIN
    INSERT INTO enrollments (
      id, course_id, user_id, completed, certificate_issued_at, certificate_template
    ) VALUES (
      p_id, p_course_id, p_user_id, p_completed, p_certificate_issued_at, p_certificate_template
    );
END;
$$;





----- Update_enrollments --------
CREATE OR REPLACE PROCEDURE update_enrollment(
  IN P_id UUID,
  IN P_completed BOOLEAN,
  IN p_certificate_template TEXT
)
LANGUAGE plpgsql
AS $$
BEGIN 
    UPDATE enrollments 
    SET completed = P_completed,
        certificate_template = p_certificate_template,
        updated_at = CURRENT_TIMESTAMP
    WHERE id = p_id AND deleted_at IS  NULL;
END;
$$;

----Get_Enrollment_BY_id
CREATE OR REPLACE FUNCTION get_enrollment_by_id(P_id UUID)
RETURNS TABLE (
  id UUID,
  course_id UUID,
  user_id UUID,
  enrollment_at TIMESTAMP,
  completed BOOLEAN,
  certificate_issued_at TIMESTAMP,
  certificate_template TEXT,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
)

LANGUAGE plpgsql AS $$
BEGIN
    RETURN QUERY
    SELECT 
        e.id,
        e.course_id,
        e.user_id,
        e.enrollment_at,
        e.completed,
        e.certificate_issued_at,
        e.certificate_template,
        e.created_at,
        e.updated_at,
        e.deleted_at
    FROM enrollments e
    WHERE e.id = P_id;
END;
$$;





CREATE OR REPLACE FUNCTION get_all_enrollments()
RETURNS TABLE (
  id UUID,
  course_id UUID,
  user_id UUID,
  enrollment_at TIMESTAMP,
  completed BOOLEAN,
  certificate_issued_at TIMESTAMP,
  certificate_template TEXT,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
)
LANGUAGE plpgsql AS $$
BEGIN
    RETURN QUERY
    SELECT 
        e.id,
        e.course_id,
        e.user_id,
        e.enrollment_at,
        e.completed,
        e.certificate_issued_at,
        e.certificate_template,
        e.created_at,
        e.updated_at,
        e.deleted_at
    FROM enrollments e
    WHERE e.deleted_at IS NULL;  -- exclude soft-deleted
END;
$$;



----------- Delete_enrollments----------
CREATE OR REPLACE PROCEDURE delete_enrollment (
  IN p_id UUID
)
LANGUAGE plpgsql 
AS $$
BEGIN
    DELETE FROM enrollments WHERE id = p_id;
END;
$$;


        


     
    


