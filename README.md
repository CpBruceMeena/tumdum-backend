# Tumdum Backend

Tumdum is a food delivery application backend service built with Go and PostgreSQL. This service provides APIs for user management, restaurant listings, menu management, and order processing.

## Features

- User authentication and authorization
- Restaurant management
- Menu and dish management
- Order processing and tracking
- RESTful API endpoints
- PostgreSQL database integration
- Swagger documentation

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

5. Run the application:
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

- Users:
  - POST /api/users - Create a new user
  - GET /api/users/{id} - Get user by ID
  - PUT /api/users/{id} - Update user
  - DELETE /api/users/{id} - Delete user
  - GET /api/users/{user_id}/orders - Get user's orders

- Restaurants:
  - GET /api/restaurants - List all restaurants
  - GET /api/restaurants/{id} - Get restaurant by ID
  - GET /api/restaurants/{restaurant_id}/dishes - Get restaurant's dishes

- Dishes:
  - GET /api/dishes/{id} - Get dish by ID

- Orders:
  - POST /api/orders - Create a new order
  - GET /api/orders/{id} - Get order by ID
  - PUT /api/orders/{id}/status - Update order status

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details. 