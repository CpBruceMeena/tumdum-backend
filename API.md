# Tumdum API Documentation

This document provides detailed information about the Tumdum API endpoints, request/response formats, and examples.

## Base URL

```
http://localhost:8080/api
```

## Authentication

Currently, the API does not require authentication. This will be implemented in future versions.

## Image Management

### Upload Image

Upload an image file (restaurant logo, cover image, or dish image).

```http
POST /api/images/upload
Content-Type: multipart/form-data
```

**Form Data:**
- `file` (required): Image file (JPG, JPEG, or PNG)
- `type` (required): Image type (restaurant_logo, restaurant_cover, or dish)
- `id` (required): ID of the entity (restaurant_id or dish_id)

**Response:**
```json
{
    "url": "/images/restaurant_logo_1.jpg"
}
```

### Image URL Structure

Images are organized by type and entity ID:

1. Restaurant Logos:
   - URL: `/images/restaurant_logo_{restaurant_id}.jpg`
   - Example: `/images/restaurant_logo_1.jpg`

2. Restaurant Cover Images:
   - URL: `/images/restaurant_cover_{restaurant_id}.jpg`
   - Example: `/images/restaurant_cover_1.jpg`

3. Dish Images:
   - URL: `/images/dish_{dish_id}.jpg`
   - Example: `/images/dish_1.jpg`

### Delete Image

Delete an uploaded image.

```http
DELETE /api/images?url=/images/restaurant_logo_1.jpg
```

**Response:**
```json
{
    "message": "Image deleted successfully"
}
```

## Users

### Create User

Create a new user.

```http
POST /api/users
Content-Type: application/json
```

**Request Body:**
```json
{
    "name": "John Doe",
    "email": "john@example.com",
    "phone": "+1234567890",
    "address": "123 Main St",
    "city": "New York",
    "state": "NY",
    "country": "USA",
    "postal_code": "10001"
}
```

**Response:**
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
    "created_at": "2024-03-20T10:00:00Z",
    "updated_at": "2024-03-20T10:00:00Z"
}
```

### Get User by ID

Get user details by ID.

```http
GET /api/users/{id}
```

**Response:**
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
    "created_at": "2024-03-20T10:00:00Z",
    "updated_at": "2024-03-20T10:00:00Z"
}
```

### Update User

Update user details.

```http
PUT /api/users/{id}
Content-Type: application/json
```

**Request Body:**
```json
{
    "name": "John Smith",
    "email": "john.smith@example.com",
    "phone": "+1234567890",
    "address": "456 Oak St",
    "city": "Los Angeles",
    "state": "CA",
    "country": "USA",
    "postal_code": "90001"
}
```

**Response:**
```json
{
    "id": 1,
    "name": "John Smith",
    "email": "john.smith@example.com",
    "phone": "+1234567890",
    "address": "456 Oak St",
    "city": "Los Angeles",
    "state": "CA",
    "country": "USA",
    "postal_code": "90001",
    "created_at": "2024-03-20T10:00:00Z",
    "updated_at": "2024-03-20T11:00:00Z"
}
```

### Delete User

Delete a user.

```http
DELETE /api/users/{id}
```

**Response:**
```json
{
    "message": "User deleted successfully"
}
```

### Get User's Orders

Get all orders for a specific user.

```http
GET /users/{user_id}/orders
```

**Response:**
```json
[
    {
        "id": 1,
        "user_id": 1,
        "restaurant_id": 1,
        "status": "pending",
        "total_amount": 25.99,
        "created_at": "2024-03-20T10:00:00Z",
        "updated_at": "2024-03-20T10:00:00Z",
        "order_items": [
            {
                "id": 1,
                "order_id": 1,
                "dish_id": 1,
                "quantity": 2,
                "price": 12.99
            }
        ]
    }
]
```

## Restaurants

### Create Restaurant

Create a new restaurant with logo and cover image.

```http
POST /api/restaurants
Content-Type: multipart/form-data
```

**Form Data:**
- `name` (required): Restaurant name
- `description` (required): Restaurant description
- `email` (required): Restaurant email
- `phone` (required): Restaurant phone
- `address` (required): Restaurant address
- `city` (required): Restaurant city
- `state` (required): Restaurant state
- `country` (required): Restaurant country
- `postal_code` (required): Restaurant postal code
- `cuisine` (required): Restaurant cuisine
- `opening_time` (required): Opening time (HH:MM:SS)
- `closing_time` (required): Closing time (HH:MM:SS)
- `logo` (optional): Restaurant logo image
- `cover_image` (optional): Restaurant cover image

**Response:**
```json
{
    "id": 1,
    "name": "Tasty Bites",
    "description": "Delicious food for everyone",
    "email": "info@tastybites.com",
    "phone": "+1234567890",
    "address": "789 Food St",
    "city": "Chicago",
    "state": "IL",
    "country": "USA",
    "postal_code": "60601",
    "cuisine": "Italian",
    "opening_time": "09:00:00",
    "closing_time": "22:00:00",
    "logo_url": "/images/restaurant_logo/1.jpg",
    "cover_image_url": "/images/restaurant_cover/1.jpg",
    "created_at": "2024-03-20T10:00:00Z",
    "updated_at": "2024-03-20T10:00:00Z"
}
```

### Get All Restaurants

Get a list of all restaurants with optional filters.

```http
GET /api/restaurants?cuisine=Italian&is_active=true&city=Chicago
```

**Query Parameters:**
- `cuisine` (optional): Filter by cuisine
- `is_active` (optional): Filter by active status
- `city` (optional): Filter by city

**Response:**
```json
[
    {
        "id": 1,
        "name": "Tasty Bites",
        "description": "Delicious food for everyone",
        "email": "info@tastybites.com",
        "phone": "+1234567890",
        "address": "789 Food St",
        "city": "Chicago",
        "state": "IL",
        "country": "USA",
        "postal_code": "60601",
        "cuisine": "Italian",
        "opening_time": "09:00:00",
        "closing_time": "22:00:00",
        "logo_url": "/images/restaurant_logo/1.jpg",
        "cover_image_url": "/images/restaurant_cover/1.jpg",
        "created_at": "2024-03-20T10:00:00Z",
        "updated_at": "2024-03-20T10:00:00Z"
    }
]
```

### Get Restaurant by ID

Get restaurant details by ID.

```http
GET /api/restaurants/{id}
```

**Response:**
```json
{
    "id": 1,
    "name": "Tasty Bites",
    "description": "Delicious food for everyone",
    "email": "info@tastybites.com",
    "phone": "+1234567890",
    "address": "789 Food St",
    "city": "Chicago",
    "state": "IL",
    "country": "USA",
    "postal_code": "60601",
    "cuisine": "Italian",
    "opening_time": "09:00:00",
    "closing_time": "22:00:00",
    "logo_url": "/images/restaurant_logo/1.jpg",
    "cover_image_url": "/images/restaurant_cover/1.jpg",
    "created_at": "2024-03-20T10:00:00Z",
    "updated_at": "2024-03-20T10:00:00Z"
}
```

### Update Restaurant

Update restaurant details.

```http
PUT /api/restaurants/{id}
Content-Type: multipart/form-data
```

**Form Data:**
- `name` (optional): Restaurant name
- `description` (optional): Restaurant description
- `email` (optional): Restaurant email
- `phone` (optional): Restaurant phone
- `address` (optional): Restaurant address
- `city` (optional): Restaurant city
- `state` (optional): Restaurant state
- `country` (optional): Restaurant country
- `postal_code` (optional): Restaurant postal code
- `cuisine` (optional): Restaurant cuisine
- `opening_time` (optional): Opening time (HH:MM:SS)
- `closing_time` (optional): Closing time (HH:MM:SS)
- `logo` (optional): Restaurant logo image
- `cover_image` (optional): Restaurant cover image

**Response:**
```json
{
    "id": 1,
    "name": "Tasty Bites Updated",
    "description": "Updated description",
    "email": "info@tastybites.com",
    "phone": "+1234567890",
    "address": "789 Food St",
    "city": "Chicago",
    "state": "IL",
    "country": "USA",
    "postal_code": "60601",
    "cuisine": "Italian",
    "opening_time": "09:00:00",
    "closing_time": "22:00:00",
    "logo_url": "/images/restaurant_logo/9876543210.jpg",
    "cover_image_url": "/images/restaurant_cover/9876543210.jpg",
    "created_at": "2024-03-20T10:00:00Z",
    "updated_at": "2024-03-20T11:00:00Z"
}
```

### Delete Restaurant

Delete a restaurant.

```http
DELETE /api/restaurants/{id}
```

**Response:**
```json
{
    "message": "Restaurant deleted successfully"
}
```

### Get Restaurant's Dishes

Get all dishes for a specific restaurant.

```http
GET /api/restaurants/{id}/dishes
```

**Response:**
```json
[
    {
        "id": 1,
        "restaurant_id": 1,
        "name": "Margherita Pizza",
        "description": "Classic tomato and mozzarella pizza",
        "price": 12.99,
        "category": "Pizza",
        "image_url": "/images/dish/1.jpg",
        "is_available": true,
        "created_at": "2024-03-20T10:00:00Z",
        "updated_at": "2024-03-20T10:00:00Z"
    }
]
```

## Dishes

### Create Dish

Create a new dish with image.

```http
POST /api/restaurant-dishes/{restaurant_id}
Content-Type: multipart/form-data
```

**Form Data:**
- `name` (required): Dish name
- `description` (required): Dish description
- `price` (required): Dish price
- `category` (required): Dish category
- `image` (optional): Dish image

**Response:**
```json
{
    "id": 1,
    "restaurant_id": 1,
    "name": "Margherita Pizza",
    "description": "Classic tomato and mozzarella pizza",
    "price": 12.99,
    "category": "Pizza",
    "image_url": "/images/dish/1.jpg",
    "is_available": true,
    "created_at": "2024-03-20T10:00:00Z",
    "updated_at": "2024-03-20T10:00:00Z"
}
```

### Get Restaurant's Dishes

Get all dishes for a specific restaurant.

```http
GET /api/restaurant-dishes/{restaurant_id}
```

**Response:**
```json
[
    {
        "id": 1,
        "restaurant_id": 1,
        "name": "Margherita Pizza",
        "description": "Classic tomato and mozzarella pizza",
        "price": 12.99,
        "category": "Pizza",
        "image_url": "/images/dish/1.jpg",
        "is_available": true,
        "created_at": "2024-03-20T10:00:00Z",
        "updated_at": "2024-03-20T10:00:00Z"
    }
]
```

### Get Dish by ID

Get dish details by ID.

```http
GET /api/restaurant-dishes/{restaurant_id}/{dish_id}
```

**Response:**
```json
{
    "id": 1,
    "restaurant_id": 1,
    "name": "Margherita Pizza",
    "description": "Classic tomato and mozzarella pizza",
    "price": 12.99,
    "category": "Pizza",
    "image_url": "/images/dish/1.jpg",
    "is_available": true,
    "created_at": "2024-03-20T10:00:00Z",
    "updated_at": "2024-03-20T10:00:00Z"
}
```

### Update Dish

Update dish details.

```http
PUT /api/restaurant-dishes/{restaurant_id}/{dish_id}
Content-Type: multipart/form-data
```

**Form Data:**
- `name` (optional): Dish name
- `description` (optional): Dish description
- `price` (optional): Dish price
- `category` (optional): Dish category
- `image` (optional): Dish image

**Response:**
```json
{
    "id": 1,
    "restaurant_id": 1,
    "name": "Margherita Pizza Updated",
    "description": "Updated description",
    "price": 13.99,
    "category": "Pizza",
    "image_url": "/images/dish/9876543210.jpg",
    "is_available": true,
    "created_at": "2024-03-20T10:00:00Z",
    "updated_at": "2024-03-20T11:00:00Z"
}
```

### Delete Dish

Delete a dish.

```http
DELETE /api/restaurant-dishes/{restaurant_id}/{dish_id}
```

**Response:**
```json
{
    "message": "Dish deleted successfully"
}
```

## Orders

### Create Order

Create a new order.

```http
POST /orders
Content-Type: application/json
```

**Request Body:**
```json
{
    "user_id": 1,
    "restaurant_id": 1,
    "order_items": [
        {
            "dish_id": 1,
            "quantity": 2
        }
    ]
}
```

**Response:**
```json
{
    "id": 1,
    "user_id": 1,
    "restaurant_id": 1,
    "status": "pending",
    "total_amount": 25.98,
    "created_at": "2024-03-20T10:00:00Z",
    "updated_at": "2024-03-20T10:00:00Z",
    "order_items": [
        {
            "id": 1,
            "order_id": 1,
            "dish_id": 1,
            "quantity": 2,
            "price": 12.99
        }
    ]
}
```

### Get All Orders

Get a list of all orders.

```http
GET /orders
```

**Response:**
```json
[
    {
        "id": 1,
        "user_id": 1,
        "restaurant_id": 1,
        "status": "pending",
        "total_amount": 25.98,
        "created_at": "2024-03-20T10:00:00Z",
        "updated_at": "2024-03-20T10:00:00Z",
        "order_items": [
            {
                "id": 1,
                "order_id": 1,
                "dish_id": 1,
                "quantity": 2,
                "price": 12.99
            }
        ]
    }
]
```

### Get Order by ID

Get order details by ID.

```http
GET /orders/{id}
```

**Response:**
```json
{
    "id": 1,
    "user_id": 1,
    "restaurant_id": 1,
    "status": "pending",
    "total_amount": 25.98,
    "created_at": "2024-03-20T10:00:00Z",
    "updated_at": "2024-03-20T10:00:00Z",
    "order_items": [
        {
            "id": 1,
            "order_id": 1,
            "dish_id": 1,
            "quantity": 2,
            "price": 12.99
        }
    ]
}
```

### Update Order

Update order details.

```http
PUT /orders/{id}
Content-Type: application/json
```

**Request Body:**
```json
{
    "user_id": 1,
    "restaurant_id": 1,
    "status": "confirmed",
    "order_items": [
        {
            "dish_id": 1,
            "quantity": 3
        }
    ]
}
```

**Response:**
```json
{
    "id": 1,
    "user_id": 1,
    "restaurant_id": 1,
    "status": "confirmed",
    "total_amount": 38.97,
    "created_at": "2024-03-20T10:00:00Z",
    "updated_at": "2024-03-20T11:00:00Z",
    "order_items": [
        {
            "id": 1,
            "order_id": 1,
            "dish_id": 1,
            "quantity": 3,
            "price": 12.99
        }
    ]
}
```

### Delete Order

Delete an order.

```http
DELETE /orders/{id}
```

**Response:**
```json
{
    "message": "Order deleted successfully"
}
```

## Error Responses

All endpoints may return the following error responses:

### 400 Bad Request
```json
{
    "error": "Invalid request parameters"
}
```

### 404 Not Found
```json
{
    "error": "Resource not found"
}
```

### 500 Internal Server Error
```json
{
    "error": "Internal server error"
}
```

## Order Status Flow

Orders follow this status flow:
1. `pending` → `confirmed` or `cancelled`
2. `confirmed` → `preparing`
3. `preparing` → `ready`
4. `ready` → `delivered`

## Image Guidelines

1. Supported file types: JPG, JPEG, PNG
2. Maximum file size: 5MB
3. Image types:
   - Restaurant logo (type: restaurant_logo)
   - Restaurant cover image (type: restaurant_cover)
   - Dish image (type: dish) 