🛒 E-commerce Backend API (Golang)

A production-grade e-commerce backend API built with Go, designed using Clean Architecture, the Repository Pattern, and GitHub Flow.
The system implements JWT-based authentication, role-based authorization (admin/user), and full CRUD (create, read, update, delete) operations for users, products, and orders.

This project reflects real-world backend engineering practices, including manual dependency injection, context propagation, structured logging, graceful shutdown, testing at multiple layers, Dockerized deployment, and CI enforcement.

📌 Key Features
Clean Architecture (handlers → services → repositories)
Manual Dependency Injection (no DI frameworks)
JWT Authentication (access tokens)
Role-Based Authorization (admin/user)
MongoDB using the official Go driver
RESTful API with versioning (/api/v1)
Pagination for list endpoints
Structured logging with Zap
Graceful shutdown with context cancellation
Unit & integration testing (manual mocks, no mocking libraries)
Docker & docker-compose
OpenAPI v3 (Swagger) documentation
GitHub Flow + CI with GitHub Actions


🏗 Architecture Overview
HTTP Request
   ↓
Handler (HTTP layer)
   ↓
Service (Business logic)
   ↓
Repository (Data access abstraction)
   ↓
MongoDB


Why this architecture?
Keeps business logic independent of frameworks
Makes the database swappable with minimal changes
Enables fast, deterministic unit tests
Scales well for real teams and long-term maintenance


📁 Project Structure
ecommerce-api/
├── cmd/api/               # Application entry point
├── internal/
│   ├── config/            # Environment & config loading
│   ├── domain/            # Core domain models
│   ├── database/          # MongoDB connection
│   ├── repository/        # Data access implementations
│   ├── service/           # Business logic
│   ├── handler/           # HTTP handlers
│   ├── middleware/        # Auth & role middleware
│   └── routes/            # API routing
├── pkg/                   # Shared utilities (JWT, hashing, responses)
├── tests/                 # Tests (unit + integration)
├── docs/                  # Swagger (OpenAPI)
├── Dockerfile
├── docker-compose.yml
├── Makefile
└── README.md


🔐 Authentication & Authorization
Authentication
Users authenticate using email + password
Passwords are hashed with bcrypt
JWT access tokens are issued on login
Authorization
Every request carries a JWT
Middleware validates the token
Role-based middleware restricts admin routes
Roles
user → can place orders, view own data
Admin → can manage products

🌐 API Versioning
All endpoints are versioned:
/api/v1/...
This enables safe future evolution without compromising client stability.

📚 API Documentation (Swagger)
Swagger (OpenAPI v3) documentation is available and secured.
Access locally
http://localhost:8080/swagger
Swagger includes:
All endpoints
Request/response schemas
JWT security definitions
Example payloads
