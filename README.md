Here's a complete README for your Go project:

---

# Go CRUD API

This project is a simple in-memory CRUD (Create, Read, Update, Delete) API built using Go. The API allows managing user data, including first names, last names, and biographies. It uses the [go-chi/chi](https://github.com/go-chi/chi) router for handling HTTP requests and the [go-crud-in-memory](https://github.com/FelipeBelloDultra/go-crud-in-memory) package for in-memory database operations.

## Table of Contents

-   [Installation](#installation)
-   [Usage](#usage)
-   [API Endpoints](#api-endpoints)
-   [Request and Response Structures](#request-and-response-structures)
-   [Error Handling](#error-handling)
-   [Logging](#logging)

## Installation

To install this project, you need to have [Go](https://go.dev/doc/install) installed on your machine. Clone the repository and install the dependencies:

```bash
git clone git@github.com:FelipeBelloDultra/go-crud-in-memory.git
cd go-crud-api
go mod tidy
```

## Usage

To start the API server, run:

```bash
go run main.go
```

The server will start on `http://localhost:8080`. You can use tools like [Postman](https://www.postman.com/) or `curl` to interact with the API.

## API Endpoints

The API provides the following endpoints:

### 1. Create User

-   **Endpoint**: `POST /api/users`
-   **Description**: Creates a new user.
-   **Request Body**:
    ```json
    {
        "first_name": "John",
        "last_name": "Doe",
        "biography": "Software developer with a passion for open-source."
    }
    ```
-   **Response**:
    ```json
    {
        "data": "user_id"
    }
    ```

### 2. List Users

-   **Endpoint**: `GET /api/users`
-   **Description**: Retrieves a list of all users.
-   **Response**:
    ```json
    {
        "data": [
            {
                "id": "user_id",
                "first_name": "John",
                "last_name": "Doe",
                "biography": "Software developer with a passion for open-source."
            }
        ]
    }
    ```

### 3. Get User by ID

-   **Endpoint**: `GET /api/users/{id}`
-   **Description**: Retrieves a user by their ID.
-   **Response**:
    ```json
    {
        "data": {
            "id": "user_id",
            "first_name": "John",
            "last_name": "Doe",
            "biography": "Software developer with a passion for open-source."
        }
    }
    ```

### 4. Update User by ID

-   **Endpoint**: `PUT /api/users/{id}`
-   **Description**: Updates a user's information by their ID.
-   **Request Body**:
    ```json
    {
        "first_name": "John",
        "last_name": "Doe",
        "biography": "Updated biography."
    }
    ```
-   **Response**:
    ```json
    {
        "data": {
            "id": "user_id",
            "first_name": "John",
            "last_name": "Doe",
            "biography": "Updated biography."
        }
    }
    ```

### 5. Delete User by ID

-   **Endpoint**: `DELETE /api/users/{id}`
-   **Description**: Deletes a user by their ID.
-   **Response**:
    ```json
    {
        "data": "id user_id deleted"
    }
    ```

## Request and Response Structures

### UserRequestBody

The structure of the request body for creating or updating a user:

```go
type UserRequestBody struct {
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    Biography string `json:"biography"`
}
```

### UserResponse

The structure of the response when retrieving user information:

```go
type UserResponse struct {
    ID        string `json:"id"`
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    Biography string `json:"biography"`
}
```

### Response

General structure of the API response:

```go
type Response struct {
    Error string `json:"error,omitempty"`
    Data  any    `json:"data,omitempty"`
}
```

## Error Handling

Errors are returned in the response with an appropriate HTTP status code. For example:

-   `400 Bad Request` for validation errors.
-   `404 Not Found` when a user ID does not exist.
-   `500 Internal Server Error` for unexpected errors.

## Logging

The API uses structured logging through `log/slog` to log errors and significant events.
