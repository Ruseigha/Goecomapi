# 🛒 E-commerce Backend API (Golang)

![CI](https://github.com/<your-username>/<repo-name>/actions/workflows/ci.yml/badge.svg)
![Go](https://img.shields.io/badge/Go-1.20+-00ADD8?logo=go)
![MongoDB](https://img.shields.io/badge/MongoDB-Official_Driver-47A248?logo=mongodb)
![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?logo=docker)
![License](https://img.shields.io/badge/License-MIT-green)
![Coverage](https://img.shields.io/codecov/c/github/<your-username>/<repo-name>)

A **production-grade e-commerce backend API** built with **Go**, designed using **Clean Architecture**, the **Repository Pattern**, and **GitHub Flow**.  
The system implements **JWT-based authentication**, **role-based authorization (admin/user)**, and full **CRUD operations** for users, products, and orders.

This project reflects **real-world backend engineering practices**, including **manual dependency injection**, **context propagation**, **structured logging**, **graceful shutdown**, **testing at multiple layers**, **Dockerized deployment**, and **CI enforcement**.

---

## 📌 Key Features

- Clean Architecture (handlers → services → repositories)
- Manual Dependency Injection (no DI frameworks)
- JWT Authentication (access tokens)
- Role-Based Authorization (admin / user)
- MongoDB using the official Go driver
- RESTful API with versioning (`/api/v1`)
- Pagination for list endpoints
- Structured logging with Zap
- Graceful shutdown with context cancellation
- Unit & integration testing (manual mocks, no mocking libraries)
- Docker & docker-compose
- OpenAPI v3 (Swagger) documentation
- GitHub Flow + CI with GitHub Actions

---

## 🏗 Architecture Over
                           ┌──────────────┐
                           │     Client      │
                           │  (Web / App)    │
                           └──────┬───────┘
                                   │ HTTP Request (JSON)
                                   ▼
                           ┌──────────────┐
                           │   Router        │
                           │ (gorilla/mux)   │
                           └──────┬───────┘
                                   │
                                   ▼
                           ┌──────────────┐
                           │   Handler       │
                           │ (HTTP Layer)    │
                           └──────┬───────┘
                                   │ context.Context
                                   ▼
                           ┌──────────────┐
                           │   Service       │
                           │ (Business       │
                           │   Logic)        │
                           └──────┬───────┘
                                   │ context.Context
                                   ▼
                           ┌──────────────┐
                           │ Repository      │
                           │ (Data Access    │
                           │ Abstraction)    │
                           └──────┬───────┘
                                   │ context.Context
                                   ▼
                           ┌──────────────┐
                           │   MongoDB       │
                           │ (Persistence)   │
                           └──────────────┘

### Why this architecture?
- Keeps business logic independent of frameworks
- Makes the database swappable with minimal changes
- Enables fast, deterministic unit tests
- Scales well for real teams and long-term maintenance

---

## 📁 Project Structure
       ecommerce-api/
       ├── cmd/
       │   └── api/               # Application entry point (main.go)
       ├── internal/
       │   ├── config/            # Environment & config loading (load .env, app settings)
       │   ├── domain/            # Core domain models (User, Product, Order structs)
       │   ├── database/          # MongoDB connection setup and client provider
       │   ├── repository/        # Data access implementations (interfaces + MongoDB)
       │   ├── service/           # Business logic (AuthService, ProductService, etc.)
       │   ├── handler/           # HTTP handlers (handle requests & responses)
       │   ├── middleware/        # Auth & role middleware (JWT validation, admin-only routes)
       │   └── routes/            # API routing (versioned endpoints /api/v1)
       ├── pkg/                   # Shared utilities (JWT, hashing, JSON response helpers)
       ├── tests/                 # Tests (unit + integration for handlers, services, repositories)
       ├── docs/                  # Swagger/OpenAPI documentation
       ├── Dockerfile             # Docker image build instructions
       ├── docker-compose.yml     # Compose for API + MongoDB (+ optional Mongo Express)
       ├── Makefile               # Build, run, test, lint, docker commands
       ├── README.md              # Project overview, setup, architecture, API docs
       └── CONTRIBUTING.md        # Contribution guidelines, commit standards, workflow

## 🔐 Authentication & Authorization

### Authentication
- Users authenticate using **email + password**
- Passwords are hashed with **bcrypt**
- JWT access tokens are issued on login

### Authorization
- Every request carries a JWT
- Middleware validates the token
- Role-based middleware restricts admin routes

**Roles**
- `user` → can place orders, view own data
- `admin` → can manage products

---

## 🌐 API Versioning

All endpoints are versioned:

     /api/v1/...  
This allows safe future evolution without breaking clients.

## 📚 API Documentation (Swagger)
Swagger (OpenAPI v3) documentation is available and secured.

### Access locally

    http://localhost:8080/swagger 

##Swagger includes:
- All endpoints
- Request/response schemas
- JWT security definitions
- Example payloads

### Testing Principles
- Go `testing` package only
- Manual mocks (no mocking libraries)
- Table-driven tests
- Arrange → Act → Assert
- Context cancellation covered
- Deterministic & fast

Run all tests:
```bash
go test ./... -cover
