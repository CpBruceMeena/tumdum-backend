# Tumdum Backend

Tumdum is a food delivery application backend service built with Go and PostgreSQL. This service provides APIs for user management, restaurant listings, menu management, and order processing.

## Features

- User authentication and authorization
- Restaurant management with image support (logo and cover images)
- Menu and dish management with image support
- Order processing and tracking
- RESTful API endpoints
- PostgreSQL database integration
- Swagger documentation
- Image upload and management

## Prerequisites

- Go 1.21 or higher
- PostgreSQL 15 or higher
- Make (optional, for using Makefile commands)
- Git

## Project Structure

```
tumdum_backend/
├── api/                 # API handlers and routes
├── business/            # Business logic layer
├── config/             # Configuration files
├── dao/                # Data Access Objects
├── database/           # Database related code
├── models/             # Data models
├── uploads/            # Image upload directory
├── config.yaml         # Application configuration
├── go.mod              # Go module file
├── go.sum              # Go dependencies checksum
└── main.go             # Application entry point
```

## Setup and Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/tumdum_backend.git
   cd tumdum_backend
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Configure the application:
   - Copy the example config file:
     ```bash
     cp config/config.yaml.example config/config.yaml
     ```
   - Update the database and server settings in `config/config.yaml`:
     ```yaml
     database:
       host: localhost
       port: 5432
       username: your_db_user
       password: your_db_password
       dbname: tumdum
       sslmode: disable

     server:
       port: 8080
     ```

4. Set up the database:
   ```bash
   # Create database
   createdb tumdum

   # Run migrations (if available)
   make migrate-up
   ```

5. Create uploads directory:
   ```bash
   mkdir -p uploads
   ```

6. Run the application:
   ```bash
   # Development mode
   go run main.go

   # Or build and run
   go build -o tumdum
   ./tumdum
   ```

## Development Workflow

### Running the Application

1. Start the application:
   ```bash
   go run main.go
   ```

2. The server will start on port 8080 (or the port specified in config.yaml)
   - API documentation will be available at: http://localhost:8080/swagger/index.html
   - API endpoints will be available at: http://localhost:8080/api/*
   - Uploaded images will be available at: http://localhost:8080/images/*

### Making Changes

1. Create a new branch for your changes:
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. Make your changes and test them locally

3. Stage your changes:
   ```bash
   git add .
   ```

4. Commit your changes:
   ```bash
   git commit -m "Description of your changes"
   ```

5. Push your changes:
   ```bash
   git push origin feature/your-feature-name
   ```

6. Create a Pull Request on GitHub

### Code Quality

Before committing your changes:

1. Format your code:
   ```bash
   go fmt ./...
   ```

2. Run tests:
   ```bash
   go test ./...
   ```

3. Check for linting issues:
   ```bash
   go vet ./...
   ```

## API Documentation

The API documentation is available at `http://localhost:8080/swagger/index.html` when the application is running.

### Available Endpoints

#### Image Management
- POST /api/images/upload
  - Upload an image for restaurants or dishes
  - Content-Type: multipart/form-data
  - Parameters:
    - file: Image file (required)
    - type: Image type (restaurant_logo, restaurant_cover, dish) (required)
  - Returns: Image URL and metadata

- DELETE /api/images
  - Delete an image
  - Parameters:
    - url: Image URL to delete (required)
  - Returns: Success message

#### Users
- POST /api/users
  - Create a new user
  - Body: User details (name, email, phone, address)
  - Returns: Created user details

- GET /api/users/{id}
  - Get user by ID
  - Returns: User details

- PUT /api/users/{id}
  - Update user
  - Body: Updated user details
  - Returns: Updated user details

- DELETE /api/users/{id}
  - Delete user
  - Returns: Success message

- GET /api/users/{id}/orders
  - Get user's orders
  - Returns: List of orders

#### Restaurants
- POST /api/restaurants
  - Create a new restaurant
  - Body: Restaurant details (name, description, email, phone, address, city, state, country, postal_code, cuisine, rating, is_active)
  - Optional: logo and cover image files
  - Returns: Created restaurant details

- GET /api/restaurants
  - List all restaurants
  - Query Parameters:
    - cuisine: Filter by cuisine type
    - city: Filter by city
    - is_active: Filter by active status
  - Returns: List of restaurants

- GET /api/restaurants/{id}
  - Get restaurant by ID
  - Returns: Restaurant details

- PUT /api/restaurants/{id}
  - Update restaurant
  - Body: Updated restaurant details
  - Optional: New logo and cover image files
  - Returns: Updated restaurant details

- DELETE /api/restaurants/{id}
  - Delete restaurant
  - Returns: Success message

#### Dishes
- POST /api/restaurant-dishes/{restaurant_id}
  - Create a new dish
  - Body: Dish details (name, description, price, category)
  - Optional: Dish image file
  - Returns: Created dish details

- GET /api/restaurant-dishes/{restaurant_id}
  - Get restaurant's dishes
  - Query Parameters:
    - category: Filter by category
    - is_available: Filter by availability
  - Returns: List of dishes

- GET /api/restaurant-dishes/{restaurant_id}/{dish_id}
  - Get dish by ID
  - Returns: Dish details

- PUT /api/restaurant-dishes/{restaurant_id}/{dish_id}
  - Update dish
  - Body: Updated dish details
  - Optional: New dish image file
  - Returns: Updated dish details

- DELETE /api/restaurant-dishes/{restaurant_id}/{dish_id}
  - Delete dish
  - Returns: Success message

#### Orders
- POST /api/orders
  - Create a new order
  - Body: Order details (user_id, restaurant_id, items, delivery_address)
  - Returns: Created order details

- GET /api/orders
  - List all orders
  - Query Parameters:
    - status: Filter by order status
    - user_id: Filter by user
    - restaurant_id: Filter by restaurant
  - Returns: List of orders

- GET /api/orders/{id}
  - Get order by ID
  - Returns: Order details

- PUT /api/orders/{id}
  - Update order
  - Body: Updated order details
  - Returns: Updated order details

- DELETE /api/orders/{id}
  - Delete order
  - Returns: Success message

### Image Upload Guidelines

1. Supported file types: JPG, JPEG, PNG
2. Maximum file size: 5MB
3. Image types and locations:
   - Restaurant logo (type: restaurant_logo)
     - URL format: `/images/restaurant_logo/{restaurant_id}.jpg`
     - Recommended size: 400x400 pixels
     - Square aspect ratio
   - Restaurant cover image (type: restaurant_cover)
     - URL format: `/images/restaurant_cover/{restaurant_id}.jpg`
     - Recommended size: 1200x400 pixels
     - Wide aspect ratio
   - Dish image (type: dish)
     - URL format: `/images/dish/{dish_id}.jpg`
     - Recommended size: 400x400 pixels
     - Square aspect ratio

4. Image Storage:
   - Images are stored in the `uploads` directory with subdirectories for each type
   - Served through the `/images` endpoint
   - Old images are automatically deleted when replaced
   - Directory structure:
     ```
     uploads/
     ├── restaurant_logo/
     │   ├── 1.jpg
     │   ├── 2.jpg
     │   └── ...
     ├── restaurant_cover/
     │   ├── 1.jpg
     │   ├── 2.jpg
     │   └── ...
     └── dish/
         ├── 1.jpg
         ├── 2.jpg
         └── ...
     ```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

### Restaurant Endpoints

#### Get All Restaurants
- **GET** `/api/restaurants`
- **Description**: Get a list of all restaurants
- **Query Parameters**:
  - `cuisine` (optional): Filter by cuisine type
  - `rating` (optional): Filter by minimum rating
  - `is_active` (optional): Filter by active status
- **Response**: List of restaurants with their details

#### Get Restaurant by ID
- **GET** `/api/restaurants/:id`
- **Description**: Get details of a specific restaurant
- **Response**: Restaurant details including name, description, contact info, location, cuisine, and rating

#### Get Restaurant Dishes
- **GET** `/api/restaurants/:id/dishes`
- **Description**: Get all dishes for a specific restaurant
- **Response**: List of dishes with their details including name, description, price, category, and availability

#### Create Restaurant
- **POST** `/api/restaurants`
- **Description**: Create a new restaurant
- **Body**:
  ```json
  {
    "name": "string",
    "description": "string",
    "email": "string",
    "phone": "string",
    "address": "string",
    "city": "string",
    "state": "string",
    "country": "string",
    "postal_code": "string",
    "cuisine": "string",
    "rating": "number",
    "is_active": "boolean"
  }
  ```
- **Response**: Created restaurant details

#### Update Restaurant
- **PUT** `/api/restaurants/:id`
- **Description**: Update an existing restaurant
- **Body**: Same as Create Restaurant
- **Response**: Updated restaurant details

#### Delete Restaurant
- **DELETE** `/api/restaurants/:id`
- **Description**: Delete a restaurant
- **Response**: Success message

### Dish Endpoints

#### Get Dishes by Restaurant ID
- **GET** `/api/restaurant-dishes/:restaurant_id`
- **Description**: Get all dishes for a specific restaurant
- **Response**: List of dishes with their details

#### Get Dish by ID
- **GET** `/api/restaurant-dishes/:restaurant_id/:dish_id`
- **Description**: Get details of a specific dish
- **Response**: Dish details including name, description, price, and category

#### Create Dish
- **POST** `/api/restaurant-dishes/:restaurant_id`
- **Description**: Create a new dish for a restaurant
- **Body**:
  ```json
  {
    "name": "string",
    "description": "string",
    "price": "number",
    "category": "string",
    "is_available": "boolean",
    "image_url": "string"
  }
  ```
- **Response**: Created dish details

#### Update Dish
- **PUT** `/api/restaurant-dishes/:restaurant_id/:dish_id`
- **Description**: Update an existing dish
- **Body**: Same as Create Dish
- **Response**: Updated dish details

#### Delete Dish
- **DELETE** `/api/restaurant-dishes/:restaurant_id/:dish_id`
- **Description**: Delete a dish
- **Response**: Success message 