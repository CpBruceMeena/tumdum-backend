-- Add deleted_at column to users table
ALTER TABLE users ADD COLUMN deleted_at TIMESTAMP WITH TIME ZONE;

-- Add deleted_at column to restaurants table
ALTER TABLE restaurants ADD COLUMN deleted_at TIMESTAMP WITH TIME ZONE;

-- Add deleted_at column to dishes table
ALTER TABLE dishes ADD COLUMN deleted_at TIMESTAMP WITH TIME ZONE;

-- Add deleted_at column to orders table
ALTER TABLE orders ADD COLUMN deleted_at TIMESTAMP WITH TIME ZONE;

-- Add deleted_at column to order_items table
ALTER TABLE order_items ADD COLUMN deleted_at TIMESTAMP WITH TIME ZONE;

-- Create indexes for soft delete columns
CREATE INDEX idx_users_deleted_at ON users(deleted_at);
CREATE INDEX idx_restaurants_deleted_at ON restaurants(deleted_at);
CREATE INDEX idx_dishes_deleted_at ON dishes(deleted_at);
CREATE INDEX idx_orders_deleted_at ON orders(deleted_at);
CREATE INDEX idx_order_items_deleted_at ON order_items(deleted_at); 