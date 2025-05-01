-- Drop indexes for soft delete columns
DROP INDEX IF EXISTS idx_users_deleted_at;
DROP INDEX IF EXISTS idx_restaurants_deleted_at;
DROP INDEX IF EXISTS idx_dishes_deleted_at;
DROP INDEX IF EXISTS idx_orders_deleted_at;
DROP INDEX IF EXISTS idx_order_items_deleted_at;

-- Remove deleted_at column from users table
ALTER TABLE users DROP COLUMN IF EXISTS deleted_at;

-- Remove deleted_at column from restaurants table
ALTER TABLE restaurants DROP COLUMN IF EXISTS deleted_at;

-- Remove deleted_at column from dishes table
ALTER TABLE dishes DROP COLUMN IF EXISTS deleted_at;

-- Remove deleted_at column from orders table
ALTER TABLE orders DROP COLUMN IF EXISTS deleted_at;

-- Remove deleted_at column from order_items table
ALTER TABLE order_items DROP COLUMN IF EXISTS deleted_at; 