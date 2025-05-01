-- Add missing columns to users table
ALTER TABLE users 
    ADD COLUMN city VARCHAR(50) NOT NULL DEFAULT '',
    ADD COLUMN state VARCHAR(50) NOT NULL DEFAULT '',
    ADD COLUMN country VARCHAR(50) NOT NULL DEFAULT '',
    ADD COLUMN postal_code VARCHAR(20) NOT NULL DEFAULT ''; 