# GoFiber Clean Architecture API

A clean and modular REST API boilerplate built with GoFiber, MongoDB, and JWT Authentication using Clean Architecture principles.

## 📁 Project Structure

│   .air.toml
│   .dockerignore
│   .env
│   .gitignore
│   docker-clean.bat
│   docker-compose.yml
│   docker-down.bat
│   docker-up.bat
│   dockerfile
│   go.mod
│   go.sum
│   main.go
├───.vscode
│       launch.json
├───configuration
│       config.go
├───container
│       container.go
├───controller
│       user_controller.go
├───database
│       mongodb.go
├───dto
│       user_dto.go
├───handler
│       auth_handler.go
│       root_handler.go
│       user_handler.go
├───mapper
│       user_mapper.go
├───middleware
│       jwt.go
├───model
│       api_model.go
│       user_model.go
├───repository
│       user_repository.go
├───routes
│       routes.go
├───service
│       auth_service.go
│       user_service.go
├───tmp
        air.log
└───validators
        validators.go

## 🚀 Features

- ✅ JWT authentication with middleware protection
- ✅ MongoDB integration
- ✅ Modular clean architecture
- ✅ DTOs and model mappers
- ✅ Dependency injection container
- ✅ Docker support (`docker-compose up --build` or docker compose down)

---

## 🔧 Prerequisites

- [Go](https://go.dev/dl/) 1.20+
- [Docker](https://www.docker.com/)
- [MongoDB](https://www.mongodb.com/)

---

## 🐳 Run with Docker

```bash up
docker-compose up --build

```bash down
docker compose down


## Test

GET / HTTP/1.1
Host: localhost:5000
User-Agent: HTTPie


HTTP/1.1 200 OK
{
  "message": "Welcome to the GoFiber Clean Architecture API!"
}

POST /auth/register
Content-Type: application/json

{
  "username": "john_doe",
  "email": "john@example.com",
  "password": "secret"
}

Http/1.1 201 Created
{
  "status": "success",
  "message": "User registered successfully"
}

HTTP/1.1 409 Conflict
{
  "status": "error",
  "message": "",
  "error": "Email already exists"
}


POST /auth/login
Content-Type: application/json

{
  "email": "john@example.com",
  "password": "secret"
}



Http/1.1 200
{
  "status": "success",
  "message": "Login successful",
  "data": {
    "email": "john.doe1@example.com",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDgwOTg5NDYsInVzZXJpZCI6IjY4MmRhZjE5M2Y1YThmY2IxODk0M2EyZSIsInVzZXJuYW1lIjoiam9obl9kb2UxIn0.8ZIKM1eB0IpMp8vSuqfw1p6BKo7WOnn8w0MjhSLG3Io",
    "userid": "682daf193f5a8fcb18943a2e",
    "username": "john_doe1"
  }
}

HTTP/1.1 401 Unauthorized
{
  "status": "error",
  "message": "",
  "error": "Invalid email or password"
}


POST /user/682daf193f5a8fcb18943a2e HTTP/1.1
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDgwOTY3MzQsInVzZXJpZCI6IjY4MmRhZjE5M2Y1YThmY2IxODk0M2EyZSIsInVzZXJuYW1lIjoiam9obl9kb2UxIn0.8Q9zB8OLQnaOHTBqT_nHIKvuMfOUJqWQPSJ_oiyFbjc
Content-Length: 0
Content-Type: application/json
Host: localhost:5000
User-Agent: HTTPie

HTTP/1.1 200 OK
{
  "status": "success",
  "message": "User fetched successfully",
  "data": {
    "email": "john.doe1@example.com",
    "userid": "682daf193f5a8fcb18943a2e",
    "username": "john_doe1"
  }
}

POST /user/error HTTP/1.1
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDgwOTY3MzQsInVzZXJpZCI6IjY4MmRhZjE5M2Y1YThmY2IxODk0M2EyZSIsInVzZXJuYW1lIjoiam9obl9kb2UxIn0.8Q9zB8OLQnaOHTBqT_nHIKvuMfOUJqWQPSJ_oiyFbjc
Content-Length: 0
Content-Type: application/json
Host: localhost:5000
User-Agent: HTTPie

HTTP/1.1 403 Forbidden
{
  "status": "error",
  "message": "Access denied",
  "error": "You are not authorized to access this resource"
}


## 🐳 Todos
Error Handler ,Code Refactoring more solid robotus way
