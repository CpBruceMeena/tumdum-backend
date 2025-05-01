# Authentication API Documentation

This document provides detailed information about the authentication-related endpoints in the Tumdum API.

## Base URL

```
http://localhost:8080/api
```

## Authentication Flow

1. Register a new user
2. Login to get a JWT token
3. Use the JWT token in subsequent requests
4. Token expires after 24 hours (configurable in config.yaml)

## API Endpoints

### 1. User Registration

Register a new user account.

```http
POST /api/users/register
Content-Type: application/json
```

**Request Body:**
```json
{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "Tumdum@123",
    "phone": "+1234567890",
    "address": "123 Main St",
    "city": "New York",
    "state": "NY",
    "country": "USA",
    "postal_code": "10001"
}
```

**Response (201 Created):**
```json
{
    "user": {
        "id": 1,
        "name": "John Doe",
        "email": "john@example.com",
        "phone": "+1234567890",
        "address": "123 Main St",
        "city": "New York",
        "state": "NY",
        "country": "USA",
        "postal_code": "10001",
        "created_at": "2024-05-01T12:00:00Z",
        "updated_at": "2024-05-01T12:00:00Z"
    },
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

**Example using curl:**
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

### 2. User Login

Login with email and password to get a JWT token.

```http
POST /api/users/login
Content-Type: application/json
```

**Request Body:**
```json
{
    "email": "john@example.com",
    "password": "Tumdum@123"
}
```

**Response (200 OK):**
```json
{
    "user": {
        "id": 1,
        "name": "John Doe",
        "email": "john@example.com",
        "phone": "+1234567890",
        "address": "123 Main St",
        "city": "New York",
        "state": "NY",
        "country": "USA",
        "postal_code": "10001",
        "created_at": "2024-05-01T12:00:00Z",
        "updated_at": "2024-05-01T12:00:00Z"
    },
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

**Example using curl:**
```bash
curl -X POST http://localhost:8080/api/users/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "Tumdum@123"
  }'
```

### 3. Get Current User

Get the current user's details using the JWT token.

```http
GET /api/users/{id}
Authorization: Bearer <token>
```

**Response (200 OK):**
```json
{
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "phone": "+1234567890",
    "address": "123 Main St",
    "city": "New York",
    "state": "NY",
    "country": "USA",
    "postal_code": "10001",
    "created_at": "2024-05-01T12:00:00Z",
    "updated_at": "2024-05-01T12:00:00Z"
}
```

**Example using curl:**
```bash
curl -X GET http://localhost:8080/api/users/1 \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
```

### 4. Update User

Update the current user's details.

```http
PUT /api/users/{id}
Authorization: Bearer <token>
Content-Type: application/json
```

**Request Body:**
```json
{
    "name": "John Updated",
    "phone": "+1987654321",
    "address": "456 New St",
    "city": "Los Angeles",
    "state": "CA",
    "country": "USA",
    "postal_code": "90001"
}
```

**Response (200 OK):**
```json
{
    "id": 1,
    "name": "John Updated",
    "email": "john@example.com",
    "phone": "+1987654321",
    "address": "456 New St",
    "city": "Los Angeles",
    "state": "CA",
    "country": "USA",
    "postal_code": "90001",
    "created_at": "2024-05-01T12:00:00Z",
    "updated_at": "2024-05-01T12:30:00Z"
}
```

**Example using curl:**
```bash
curl -X PUT http://localhost:8080/api/users/1 \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Updated",
    "phone": "+1987654321",
    "address": "456 New St",
    "city": "Los Angeles",
    "state": "CA",
    "country": "USA",
    "postal_code": "90001"
  }'
```

## Error Responses

### 400 Bad Request
```json
{
    "error": "invalid request body"
}
```

### 401 Unauthorized
```json
{
    "error": "unauthorized"
}
```

### 404 Not Found
```json
{
    "error": "user not found"
}
```

### 500 Internal Server Error
```json
{
    "error": "internal server error"
}
```

## Using the JWT Token

1. After successful login or registration, you'll receive a JWT token
2. Include this token in the Authorization header for all protected endpoints:
   ```
   Authorization: Bearer <your_token>
   ```
3. The token expires after 24 hours
4. Protected endpoints will return 401 Unauthorized if:
   - No token is provided
   - Token is invalid
   - Token has expired

## Security Best Practices

1. Always use HTTPS in production
2. Store tokens securely (e.g., in memory for web applications)
3. Never store tokens in localStorage or cookies
4. Implement token refresh mechanism for long-running sessions
5. Use strong passwords (minimum 8 characters, mix of uppercase, lowercase, numbers, and special characters)
6. Implement rate limiting for login attempts
7. Log out users when they close the application
8. Implement proper error handling for authentication failures 