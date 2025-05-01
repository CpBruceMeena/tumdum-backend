-- First, delete existing data in the correct order to respect foreign key constraints
DELETE FROM order_items;
DELETE FROM orders;
DELETE FROM dishes;
DELETE FROM restaurants;
DELETE FROM users;

-- Reset sequences
ALTER SEQUENCE users_id_seq RESTART WITH 1;
ALTER SEQUENCE restaurants_id_seq RESTART WITH 1;
ALTER SEQUENCE dishes_id_seq RESTART WITH 1;
ALTER SEQUENCE orders_id_seq RESTART WITH 1;
ALTER SEQUENCE order_items_id_seq RESTART WITH 1;

-- Insert test users
INSERT INTO users (name, email, phone, address) VALUES
('John Doe', 'john@example.com', '+1234567890', '123 Main St, New York, NY 10001'),
('Jane Smith', 'jane@example.com', '+1987654321', '456 Park Ave, New York, NY 10002'),
('Mike Johnson', 'mike@example.com', '+1122334455', '789 Broadway, New York, NY 10003'),
('Sarah Williams', 'sarah@example.com', '+1555666777', '321 5th Ave, New York, NY 10004'),
('David Brown', 'david@example.com', '+1999888777', '654 Madison Ave, New York, NY 10005');

-- Insert test restaurants with image URLs
INSERT INTO restaurants (
    name, description, email, phone, address, city, state, country, postal_code, 
    cuisine, opening_time, closing_time, rating, is_active, 
    logo_url, cover_image_url
) VALUES
('Pasta Paradise', 'Authentic Italian cuisine with a modern twist', 'contact@pastaparadise.com', '+12125551234', '123 Italian St', 'New York', 'NY', 'USA', '10001', 'Italian', '11:00:00', '22:00:00', 4.5, true, '/images/restaurant_logo_1.jpg', '/images/restaurant_cover_1.jpg'),
('Sushi Sensation', 'Fresh and innovative Japanese cuisine', 'info@sushisensation.com', '+12125552345', '456 Sushi Ave', 'New York', 'NY', 'USA', '10002', 'Japanese', '12:00:00', '23:00:00', 4.8, true, '/images/restaurant_logo_2.jpg', '/images/restaurant_cover_2.jpg'),
('Burger Bliss', 'Gourmet burgers and comfort food', 'hello@burgerbliss.com', '+12125553456', '789 Burger Rd', 'New York', 'NY', 'USA', '10003', 'American', '10:00:00', '22:00:00', 4.3, true, '/images/restaurant_logo_3.jpg', '/images/restaurant_cover_3.jpg'),
('Taco Temple', 'Authentic Mexican street food', 'order@tacotemple.com', '+12125554567', '321 Taco St', 'New York', 'NY', 'USA', '10004', 'Mexican', '11:00:00', '23:00:00', 4.6, true, '/images/restaurant_logo_4.jpg', '/images/restaurant_cover_4.jpg'),
('Curry Corner', 'Traditional Indian cuisine', 'info@currycorner.com', '+12125555678', '654 Spice Ave', 'New York', 'NY', 'USA', '10005', 'Indian', '11:30:00', '22:30:00', 4.7, true, '/images/restaurant_logo_5.jpg', '/images/restaurant_cover_5.jpg');

-- Insert dishes for Pasta Paradise (Italian) with image URLs
INSERT INTO dishes (restaurant_id, name, description, price, category, is_available, image_url) VALUES
(1, 'Spaghetti Carbonara', 'Classic pasta with eggs, cheese, pancetta, and black pepper', 16.99, 'Pasta', true, '/images/dish_1.jpg'),
(1, 'Margherita Pizza', 'Fresh tomatoes, mozzarella, basil, and olive oil', 14.99, 'Pizza', true, '/images/dish_2.jpg'),
(1, 'Lasagna Bolognese', 'Layers of pasta, meat sauce, and cheese', 18.99, 'Pasta', true, '/images/dish_3.jpg'),
(1, 'Tiramisu', 'Classic Italian dessert with coffee-soaked ladyfingers', 8.99, 'Dessert', true, '/images/dish_4.jpg'),
(1, 'Bruschetta', 'Toasted bread with tomatoes, garlic, and basil', 7.99, 'Appetizer', true, '/images/dish_5.jpg');

-- Insert dishes for Sushi Sensation (Japanese) with image URLs
INSERT INTO dishes (restaurant_id, name, description, price, category, is_available, image_url) VALUES
(2, 'California Roll', 'Crab, avocado, and cucumber roll', 12.99, 'Sushi', true, '/images/dish_6.jpg'),
(2, 'Spicy Tuna Roll', 'Fresh tuna with spicy sauce', 13.99, 'Sushi', true, '/images/dish_7.jpg'),
(2, 'Miso Soup', 'Traditional Japanese soup with tofu and seaweed', 4.99, 'Soup', true, '/images/dish_8.jpg'),
(2, 'Tempura Shrimp', 'Crispy battered shrimp with dipping sauce', 9.99, 'Appetizer', true, '/images/dish_9.jpg'),
(2, 'Green Tea Ice Cream', 'Traditional Japanese dessert', 5.99, 'Dessert', true, '/images/dish_10.jpg');

-- Insert dishes for Burger Bliss (American) with image URLs
INSERT INTO dishes (restaurant_id, name, description, price, category, is_available, image_url) VALUES
(3, 'Classic Cheeseburger', 'Angus beef patty with cheese, lettuce, and tomato', 12.99, 'Burger', true, '/images/dish_11.jpg'),
(3, 'Truffle Fries', 'Crispy fries with truffle oil and parmesan', 6.99, 'Sides', true, '/images/dish_12.jpg'),
(3, 'Chicken Wings', 'Spicy buffalo wings with blue cheese dip', 10.99, 'Appetizer', true, '/images/dish_13.jpg'),
(3, 'Chocolate Milkshake', 'Rich chocolate milkshake with whipped cream', 5.99, 'Drinks', true, '/images/dish_14.jpg'),
(3, 'Veggie Burger', 'Plant-based patty with fresh vegetables', 11.99, 'Burger', true, '/images/dish_15.jpg');

-- Insert dishes for Taco Temple (Mexican) with image URLs
INSERT INTO dishes (restaurant_id, name, description, price, category, is_available, image_url) VALUES
(4, 'Street Tacos', 'Three authentic street tacos with choice of meat', 9.99, 'Tacos', true, '/images/dish_16.jpg'),
(4, 'Quesadilla', 'Grilled tortilla with cheese and choice of filling', 8.99, 'Main Course', true, '/images/dish_17.jpg'),
(4, 'Guacamole', 'Fresh avocado dip with chips', 6.99, 'Appetizer', true, '/images/dish_18.jpg'),
(4, 'Churros', 'Crispy fried dough with cinnamon sugar', 5.99, 'Dessert', true, '/images/dish_19.jpg'),
(4, 'Horchata', 'Traditional rice drink with cinnamon', 3.99, 'Drinks', true, '/images/dish_20.jpg');

-- Insert dishes for Curry Corner (Indian) with image URLs
INSERT INTO dishes (restaurant_id, name, description, price, category, is_available, image_url) VALUES
(5, 'Butter Chicken', 'Tender chicken in tomato-based curry sauce', 15.99, 'Main Course', true, '/images/dish_21.jpg'),
(5, 'Vegetable Biryani', 'Fragrant rice with mixed vegetables and spices', 13.99, 'Main Course', true, '/images/dish_22.jpg'),
(5, 'Garlic Naan', 'Soft bread with garlic and butter', 3.99, 'Bread', true, '/images/dish_23.jpg'),
(5, 'Mango Lassi', 'Sweet yogurt drink with mango', 4.99, 'Drinks', true, '/images/dish_24.jpg'),
(5, 'Gulab Jamun', 'Sweet milk dumplings in sugar syrup', 5.99, 'Dessert', true, '/images/dish_25.jpg');

-- Insert sample orders
INSERT INTO orders (user_id, restaurant_id, status, total_amount, delivery_address) VALUES
(1, 1, 'DELIVERED', 35.97, '123 Main St, New York, NY 10001'),
(2, 2, 'PREPARING', 26.98, '456 Park Ave, New York, NY 10002'),
(3, 3, 'CONFIRMED', 19.98, '789 Broadway, New York, NY 10003'),
(4, 4, 'PENDING', 16.98, '321 5th Ave, New York, NY 10004'),
(5, 5, 'DELIVERED', 19.98, '654 Madison Ave, New York, NY 10005');

-- Insert order items
INSERT INTO order_items (order_id, dish_id, quantity, price) VALUES
(1, 1, 1, 16.99), -- Spaghetti Carbonara
(1, 4, 1, 8.99),  -- Tiramisu
(1, 5, 1, 7.99),  -- Bruschetta
(2, 6, 2, 12.99), -- California Roll
(2, 8, 1, 9.99),  -- Tempura Shrimp
(3, 11, 1, 12.99), -- Classic Cheeseburger
(3, 12, 1, 6.99),  -- Truffle Fries
(4, 16, 1, 9.99),  -- Street Tacos
(4, 18, 1, 6.99),  -- Guacamole
(5, 21, 1, 15.99), -- Butter Chicken
(5, 23, 1, 3.99);  -- Garlic Naan 