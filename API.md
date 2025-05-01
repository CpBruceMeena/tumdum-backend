# Tumdum API Documentation

This document provides detailed information about all available API endpoints in the Tumdum backend service.

## Base URL

All API endpoints are prefixed with `/api`

## Authentication

Currently, the API does not require authentication. This will be implemented in future versions.

## API Endpoints

### Users

#### Create User
```http
POST /api/users
```

Request Body:
```json
{
    "name": "John Doe",
    "email": "john@example.com",
    "phone": "+1234567890",
    "address": "123 Main St, City, Country"
}
```

Response (201 Created):
```json
{
    "id": "user_123",
    "name": "John Doe",
    "email": "john@example.com",
    "phone": "+1234567890",
    "address": "123 Main St, City, Country",
    "created_at": "2024-03-20T10:00:00Z"
}
```

#### Get User by ID
```http
GET /api/users/{id}
```

Response (200 OK):
```json
{
    "id": "user_123",
    "name": "John Doe",
    "email": "john@example.com",
    "phone": "+1234567890",
    "address": "123 Main St, City, Country",
    "created_at": "2024-03-20T10:00:00Z"
}
```

#### Update User
```http
PUT /api/users/{id}
```

Request Body:
```json
{
    "name": "John Updated",
    "phone": "+1987654321",
    "address": "456 New St, City, Country"
}
```

Response (200 OK):
```json
{
    "id": "user_123",
    "name": "John Updated",
    "email": "john@example.com",
    "phone": "+1987654321",
    "address": "456 New St, City, Country",
    "updated_at": "2024-03-20T11:00:00Z"
}
```

#### Delete User
```http
DELETE /api/users/{id}
```

Response (204 No Content)

#### Get User Orders
```http
GET /api/users/{id}/orders
```

Query Parameters:
- `page` (optional): Page number for pagination (default: 1)
- `limit` (optional): Number of items per page (default: 10)

Response (200 OK):
```json
{
    "orders": [
        {
            "id": "order_123",
            "user_id": "user_123",
            "restaurant_id": "rest_456",
            "status": "DELIVERED",
            "total_amount": 25.99,
            "created_at": "2024-03-20T10:00:00Z",
            "items": [
                {
                    "dish_id": "dish_789",
                    "quantity": 2,
                    "price": 12.99
                }
            ]
        }
    ],
    "total": 1,
    "page": 1,
    "limit": 10
}
```

### Restaurants

#### List All Restaurants
```http
GET /api/restaurants
```

Query Parameters:
- `page` (optional): Page number for pagination (default: 1)
- `limit` (optional): Number of items per page (default: 10)
- `cuisine` (optional): Filter by cuisine type
- `is_active` (optional): Filter by active status (true/false)

Response (200 OK):
```json
{
    "restaurants": [
        {
            "id": "rest_123",
            "name": "Tasty Bites",
            "description": "Delicious food for everyone",
            "address": "789 Food St, City, Country",
            "cuisine": "Italian",
            "rating": 4.5,
            "is_active": true
        }
    ],
    "total": 1,
    "page": 1,
    "limit": 10
}
```

#### Get Restaurant by ID
```http
GET /api/restaurants/{id}
```

Response (200 OK):
```json
{
    "id": "rest_123",
    "name": "Tasty Bites",
    "description": "Delicious food for everyone",
    "address": "789 Food St, City, Country",
    "cuisine": "Italian",
    "rating": 4.5,
    "is_active": true,
    "created_at": "2024-03-20T10:00:00Z"
}
```

#### Get Restaurant Dishes
```http
GET /api/restaurants/{id}/dishes
```

Query Parameters:
- `page` (optional): Page number for pagination (default: 1)
- `limit` (optional): Number of items per page (default: 10)
- `category` (optional): Filter by dish category
- `is_available` (optional): Filter by availability (true/false)

Response (200 OK):
```json
{
    "dishes": [
        {
            "id": "dish_123",
            "restaurant_id": "rest_456",
            "name": "Margherita Pizza",
            "description": "Classic tomato and mozzarella pizza",
            "price": 12.99,
            "category": "Pizza",
            "is_available": true
        }
    ],
    "total": 1,
    "page": 1,
    "limit": 10
}
```

### Dishes

#### Get Dish by ID
```http
GET /api/dishes/{id}
```

Response (200 OK):
```json
{
    "id": "dish_123",
    "restaurant_id": "rest_456",
    "name": "Margherita Pizza",
    "description": "Classic tomato and mozzarella pizza",
    "price": 12.99,
    "category": "Pizza",
    "is_available": true,
    "created_at": "2024-03-20T10:00:00Z"
}
```

### Orders

#### Create Order
```http
POST /api/orders
```

Request Body:
```json
{
    "user_id": "user_123",
    "restaurant_id": "rest_456",
    "items": [
        {
            "dish_id": "dish_789",
            "quantity": 2
        }
    ],
    "delivery_address": "123 Main St, City, Country"
}
```

Response (201 Created):
```json
{
    "id": "order_123",
    "user_id": "user_123",
    "restaurant_id": "rest_456",
    "status": "PENDING",
    "total_amount": 25.98,
    "items": [
        {
            "dish_id": "dish_789",
            "quantity": 2,
            "price": 12.99
        }
    ],
    "delivery_address": "123 Main St, City, Country",
    "created_at": "2024-03-20T10:00:00Z"
}
```

#### Get Order by ID
```http
GET /api/orders/{id}
```

Response (200 OK):
```json
{
    "id": "order_123",
    "user_id": "user_123",
    "restaurant_id": "rest_456",
    "status": "DELIVERED",
    "total_amount": 25.98,
    "items": [
        {
            "dish_id": "dish_789",
            "quantity": 2,
            "price": 12.99
        }
    ],
    "delivery_address": "123 Main St, City, Country",
    "created_at": "2024-03-20T10:00:00Z",
    "updated_at": "2024-03-20T11:00:00Z"
}
```

#### Update Order Status
```http
PUT /api/orders/{id}/status
```

Request Body:
```json
{
    "status": "PREPARING"
}
```

Available Status Values:
- `PENDING`
- `CONFIRMED`
- `PREPARING`
- `READY_FOR_DELIVERY`
- `OUT_FOR_DELIVERY`
- `DELIVERED`
- `CANCELLED`

Response (200 OK):
```json
{
    "id": "order_123",
    "status": "PREPARING",
    "updated_at": "2024-03-20T11:00:00Z"
}
```

## Error Responses

All endpoints may return the following error responses:

### 400 Bad Request
```json
{
    "error": "Invalid request parameters",
    "details": "Field 'email' is required"
}
```

### 404 Not Found
```json
{
    "error": "Resource not found",
    "details": "User with ID 'user_123' not found"
}
```

### 500 Internal Server Error
```json
{
    "error": "Internal server error",
    "details": "Database connection failed"
}
```

## Rate Limiting

Currently, there are no rate limits implemented. This will be added in future versions.

## Versioning

The current API version is v1. All endpoints are prefixed with `/api`. Future versions will be prefixed with their version number (e.g., `/api/v2`).

## Support

For API support or to report issues, please create an issue in the GitHub repository. 