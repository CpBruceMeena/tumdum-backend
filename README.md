# Tumdum Backend

A Go-based backend service for the Tumdum application, providing user management, authentication, and other core functionalities.

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
- PostgreSQL 12 or higher
- Make (optional, for using Makefile commands)

## Quick Start

The easiest way to get started is using the provided run script:

```bash
# Make the script executable (first time only)
chmod +x run.sh

# Run the application
./run.sh
```

The script will:
1. Check for required dependencies
2. Set up configuration files if they don't exist
3. Install dependencies
4. Run database migrations
5. Build and start the application

## Manual Setup

If you prefer to set up manually, follow these steps:

1. Clone the repository:
   ```bash
   git clone https://github.com/CpBruceMeena/tumdum-backend.git
   cd tumdum-backend
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Set up configuration:
   ```bash
   # Copy example configuration files
   cp .env.example .env
   cp config/config.yaml.example config/config.yaml
   
   # Update the configuration files with your values
   ```

4. Set up the database:
   ```bash
   # Create the database
   createdb tumdum
   
   # Run migrations
   psql -U postgres -f database/sql/schema.sql
   ```

5. Build and run:
   ```bash
   go build -o tumdum-backend
   ./tumdum-backend
   ```

## Configuration

The application uses two configuration files:

1. `.env` - Environment variables
2. `config/config.yaml` - Application configuration

Make sure to update both files with your specific values. Never commit the actual configuration files to version control.

### Environment Variables (.env)

```env
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=tumdum
DB_SSL_MODE=disable

# Server Configuration
SERVER_PORT=8080
SERVER_HOST=localhost

# JWT Configuration
JWT_SECRET=your_jwt_secret
JWT_EXPIRATION=24h

# API Configuration
API_VERSION=v1
API_PREFIX=/api

# Image Upload Configuration
UPLOAD_MAX_SIZE=5242880
UPLOAD_ALLOWED_TYPES=image/jpeg,image/png,image/gif
UPLOAD_DIR=uploads
```

### Application Configuration (config.yaml)

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
  secret: your_jwt_secret
  expiration: 24h

api:
  version: v1
  prefix: /api

upload:
  max_size: 5242880
  allowed_types:
    - image/jpeg
    - image/png
    - image/gif
  directory: uploads
```

## Security Notes

1. Never commit sensitive information like passwords or API keys
2. Use strong passwords in production
3. Enable SSL in production
4. Keep your dependencies updated
5. Use environment variables for sensitive data

## API Documentation

For detailed API documentation, please refer to [API.md](API.md).

## Authentication

For authentication details and JWT implementation, please refer to [docs/AUTH.md](docs/AUTH.md).

## Development

### Project Structure

```
.
├── api/            # API handlers and routes
├── auth/           # Authentication related code
├── business/       # Business logic layer
├── config/         # Configuration management
├── dao/            # Data Access Objects
├── database/       # Database related code
├── docs/           # Documentation
├── middleware/     # HTTP middleware
├── models/         # Data models
└── utils/          # Utility functions
```

### Running Tests

```bash
go test ./...
```

### Code Style

The project follows standard Go code style guidelines. Use `gofmt` to format your code:

```bash
gofmt -w .
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details. 