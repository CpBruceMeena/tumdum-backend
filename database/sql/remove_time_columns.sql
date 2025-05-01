-- Drop the time check constraint first
ALTER TABLE restaurants DROP CONSTRAINT IF EXISTS restaurants_time_check;

-- Remove the time columns
ALTER TABLE restaurants 
    DROP COLUMN IF EXISTS opening_time,
    DROP COLUMN IF EXISTS closing_time; 