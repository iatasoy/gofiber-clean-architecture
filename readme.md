# GoFiber Clean Architecture API

A clean and modular REST API boilerplate built with GoFiber, MongoDB, and JWT Authentication using Clean Architecture principles.

---
## 🔧 Prerequisites

- [Go](https://go.dev/dl/) 1.20+
- [Docker](https://www.docker.com/)
- [MongoDB](https://www.mongodb.com/)

---

## 🚀 Features

- ✅ JWT authentication with middleware protection  
- ✅ MongoDB integration  
- ✅ Modular clean architecture  
- ✅ DTOs and model mappers  
- ✅ Dependency injection container  
- ✅ Docker support (`docker-compose up --build` or `docker-compose down`)

---

## 📁 Project Structure

- `.air.toml`
- `.env`
- `main.go`
- `docker-compose.yml`
- `configuration/`
- `container/`
- `controller/`
- `database/`
- `dto/`
- `handler/`
- `mapper/`
- `middleware/`
- `model/`
- `repository/`
- `routes/`
- `service/`
- `validators/`
- `tmp/`

---

## 🐳 Run with Docker

```bash
docker-compose up --build
```

```bash
docker compose down
```

---

## Use Case

``` bash
GET / HTTP/1.1
Host: localhost:5000
User-Agent: HTTPie


HTTP/1.1 200 OK
{
  "message": "Welcome to the GoFiber Clean Architecture API!"
}

```

``` bash
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
```

``` bash
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
```

``` bash
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
```

## 🐳 Todos
- [ ] ✅ **Centralized Error Handling**: Implement a global error handler middleware for consistent and clean error responses.
- [ ] ✅ **Refactor to Follow SOLID Principles**: 
  - Single Responsibility (SRP): Ensure each package/component has one well-defined responsibility.
  - Open/Closed (OCP): Make modules extensible without modifying existing code.
  - Interface Segregation (ISP): Use smaller, focused interfaces for better testability.
  - Dependency Inversion (DIP): Depend on abstractions, not concrete implementations.
- [ ] ✅ **Validation Layer**: Add reusable request validators using a middleware pattern.
- [ ] ✅ **DTO & Mapper Enhancements**: Improve separation between internal models and external API contracts.
- [ ] ✅ **Swagger/OpenAPI Documentation**: Generate and serve API documentation automatically.
- [ ] ✅ **Logging Middleware**: Add structured logging for requests, responses, and errors.
- [ ] ✅ **Request ID Tracing**: Correlate logs with unique request IDs for debugging.
- [ ] ✅ **Unit and Integration Tests**: Cover services, handlers, and middleware with tests.
- [ ] ✅ **Config Improvements**: Centralize configuration management with profiles (dev/test/prod).
- [ ] ✅ **Makefile / CLI Tooling**: Add helper commands for build, lint, and test.


