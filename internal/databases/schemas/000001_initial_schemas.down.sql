DROP TABLE IF EXISTS matches;
DROP TABLE IF EXISTS swipes;
DROP TABLE IF EXISTS users;

DO $$
DECLARE
    t text;
BEGIN
    FOR t IN
        SELECT routine_name FROM information_schema.routines
        WHERE routine_name LIKE '%timestamp_func'
        AND routine_type = 'FUNCTION' AND routine_schema = 'public'
   LOOP
     EXECUTE 'DROP FUNCTION IF EXISTS ' || t;
   END LOOP;
END;
$$ language 'plpgsql';
