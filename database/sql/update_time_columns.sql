-- First, drop the time check constraint
ALTER TABLE restaurants DROP CONSTRAINT IF EXISTS restaurants_time_check;

-- Alter the column types
ALTER TABLE restaurants 
    ALTER COLUMN opening_time TYPE TIMESTAMP WITH TIME ZONE 
    USING opening_time::text::timestamp with time zone,
    ALTER COLUMN closing_time TYPE TIMESTAMP WITH TIME ZONE 
    USING closing_time::text::timestamp with time zone;

-- Add back the time check constraint
ALTER TABLE restaurants 
    ADD CONSTRAINT restaurants_time_check 
    CHECK (opening_time < closing_time); 