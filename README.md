# Tumdum Backend

A Go-based backend service for the Tumdum food delivery platform.

## Features

- User authentication with JWT
- Restaurant management
- Dish management
- Order processing
- Image upload and management
- PostgreSQL database
- RESTful API

## Prerequisites

- Go 1.21 or higher
- PostgreSQL 15 or higher
- Make (optional, for using Makefile commands)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/tumdum-backend.git
   cd tumdum-backend
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Create a PostgreSQL database:
   ```sql
   CREATE DATABASE tumdum;
   ```

4. Configure the application:
   - Copy the example config file:
     ```bash
     cp config/config.yaml.example config/config.yaml
     ```
   - Update the configuration values in `config/config.yaml`:
     - Set your database credentials
     - Set a secure JWT secret key
     - Adjust other settings as needed

## Configuration

The application uses a YAML configuration file (`config/config.yaml`). A sample configuration file (`config.yaml.example`) is provided as a template. Never commit your actual `config.yaml` file to version control.

### Configuration Structure

```yaml
# Database Configuration
database:
  host: localhost
  port: 5432
  user: your_db_user
  password: your_db_password
  name: tumdum
  ssl_mode: disable

# Server Configuration
server:
  port: 8080
  host: localhost

# JWT Configuration
jwt:
  secret: your_jwt_secret_key_here
  expiration: 24h

# API Configuration
api:
  version: v1
  prefix: /api

# Image Upload Configuration
upload:
  max_size: 5242880  # 5MB in bytes
  allowed_types:
    - image/jpeg
    - image/png
  upload_dir: uploads
```

### Security Notes

1. Never commit `config.yaml` to version control
2. Use strong, unique passwords for database access
3. Use a strong, random string for JWT secret
4. In production:
   - Use environment variables for sensitive data
   - Enable SSL for database connections
   - Use HTTPS for API endpoints

## Running the Application

1. Start the server:
   ```bash
   go run main.go
   ```

   Or using Make:
   ```bash
   make run
   ```

2. The server will start on `http://localhost:8080`

## API Documentation

The API documentation is available in [API.md](API.md) and [docs/AUTH.md](docs/AUTH.md). Here's a quick overview of the main endpoints:

### Authentication

1. Register a new user:
   ```bash
   curl -X POST http://localhost:8080/api/users/register \
     -H "Content-Type: application/json" \
     -d '{
       "name": "John Doe",
       "email": "john@example.com",
       "password": "Tumdum@123",
       "phone": "+1234567890",
       "address": "123 Main St",
       "city": "New York",
       "state": "NY",
       "country": "USA",
       "postal_code": "10001"
     }'
   ```

2. Login:
   ```bash
   curl -X POST http://localhost:8080/api/users/login \
     -H "Content-Type: application/json" \
     -d '{
       "email": "john@example.com",
       "password": "Tumdum@123"
     }'
   ```

3. Use the JWT token for authenticated requests:
   ```bash
   curl -X GET http://localhost:8080/api/restaurants \
     -H "Authorization: Bearer <your_token>"
   ```

### Protected Endpoints

All endpoints except registration, login, and image upload require authentication. Include the JWT token in the Authorization header:

```bash
Authorization: Bearer <your_token>
```

## Development

### Project Structure

```
.
├── api/            # API handlers and routes
├── auth/           # Authentication package
├── business/       # Business logic layer
├── config/         # Configuration
├── dao/            # Data Access Objects
├── database/       # Database initialization and migrations
├── middleware/     # HTTP middleware
├── models/         # Data models
├── uploads/        # Uploaded images
├── main.go         # Application entry point
└── Makefile        # Build and development commands
```

### Available Make Commands

- `make run`: Run the application
- `make build`: Build the application
- `make test`: Run tests
- `make clean`: Clean build artifacts
- `make migrate`: Run database migrations
- `make seed`: Seed the database with sample data

## Testing

Run the test suite:
```bash
go test ./...
```

Or using Make:
```bash
make test
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details. 