--- CREATE ENUM type for enrollment status if not exists 
DO $$ 
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'course_status') THEN
        CREATE TYPE course_status AS ENUM ('pending', 'active', 'completed', 'canceled')
    END IF;
END
$$;

----- Create enrollments table if not exists
CREATE TABLE IF NOT EXISTS enrollments(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  course_id UUID NOT NULL REFERENCES  courses(id) ON DELETE CASCADE,
  user_id UUD NOT NULL REFERENCES users(id) on DELETE CASCADE,
  enrollment_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
  completed  BOOLEAN  NOT NULL DEFAULT FALSE,
  certificate_issued_at TIMESTAMP NULL, 
  certificate_template TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)

----- Create enrollment ------
CREATE OR REPLACE  PROCEDURE create_enrollment(
  IN p_id
  IN P_course_id UUID,
  IN p_user_id UUID,
  
)
LANGUAGE plpgsql AS $$
BEGIN
    INSERT INTO enrollments (
      id, course_id, user_id, enrollment_at, completed, certificate_issued_at, certificate_template
    )VALUES (
      uuid_generate_v4, course_id, user_id, CURRENT_TIMESTAMP, FALSE, NULL, NULL 
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
    WHERE id = p_id AND deleted_at IS  NULL
END;
$$;

----- get_enrollments_by_ID --------
CREATE OR REPLACE FUNCTION get_enrollment_by_course(P_course_id UUID)
RETURNS TABLE (
  id UUID,
  course_id UUID,
  user_id UUID,
  enrollment_at TIMESTAMP  ,
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
        e.updated_at
    FROM enrollments e
    WHER e_course_id = P_course_id AND e.deleted_at IS NULL;
END;
$$;

----- get_enrollments --------

CREATE OR REPLACE FUNCTION  get_all_enrollments()
RETURNS TABLE (
  id UUID,
  course_id UUID,
  user_id UUID,
  enrollment_at TIMESTAMP  ,
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
        e.updated_at
    FROM enrollments e
    WHERE c.deleted_at IS NULL;
END;
$$


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


        


     
    


