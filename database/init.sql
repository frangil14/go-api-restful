-- Create database if not exists
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_database WHERE datname = 'mydatabase') THEN
        EXECUTE 'CREATE DATABASE mydatabase';
        RAISE NOTICE 'mydatabase created by Init script';
    ELSE
        RAISE NOTICE 'mydatabase already existed';
    END IF;
END
$$;

-- Create table
\c mydatabase

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'users') THEN
        EXECUTE '
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            name VARCHAR(50) NOT NULL,
            age INTEGER NOT NULL
        )';
        RAISE NOTICE 'users table created by Init script';
    ELSE
        RAISE NOTICE 'users table already existed';
    END IF;
END
$$;