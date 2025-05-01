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
   - Copy `config/config.yaml.example` to `config/config.yaml`
   - Update the configuration values in `config/config.yaml`

## Configuration

The application uses a YAML configuration file (`config/config.yaml`) with the following structure:

```yaml
database:
  host: localhost
  port: 5432
  user: postgres
  password: your_password
  name: tumdum
  ssl_mode: disable

server:
  port: 8080
  host: localhost

jwt:
  secret: your_jwt_secret_key
  expiration: 24h

api:
  version: v1
  prefix: /api
```

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