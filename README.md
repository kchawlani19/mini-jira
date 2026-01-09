# Mini Jira Backend (Go + MySQL)

Mini Jira is a backend-only implementation of a Jira-like issue tracking system,
built using **pure Go (net/http)** and **MySQL**, without relying on any web frameworks.

The project focuses on **backend fundamentals** such as authentication,
authorization, ownership enforcement, database integrity, and clean architecture.

---

## Project Objective

The goal of this project is to demonstrate how a real-world backend system is designed,
where **rules are enforced at multiple layers** (API + database) and responsibilities
are clearly separated.

This is not a UI-driven project — it is intentionally backend-focused.

---

## Key Highlights (Why this project stands out)

- No web frameworks — only standard Go libraries
- Clear separation of concerns (HTTP, business logic, database)
- Role-based access control (ADMIN vs USER)
- Strict task ownership enforcement
- Database-level constraints for data integrity
- Soft delete strategy instead of hard deletes
- Pagination implemented at database level
- Designed with interview discussions in mind

---

## Authentication & Authorization

- User registration and login
- Password hashing
- JWT-based authentication
- Middleware-based request protection
- Claims include user ID and role

### Roles
- USER
- ADMIN

---

## Role-Based Access Control

USER:
- Can create tasks (unassigned)
- Can view only tasks assigned to them
- Can update only their own tasks

ADMIN:
- Can view all tasks
- Can assign tasks to users
- Can update or delete any task
- Can access users listing

Authorization is enforced **before any database mutation**.

---

## Jira-like Task Workflow

- Tasks are created **unassigned**
- Only ADMIN can assign tasks
- A USER can update a task only after it is assigned to them
- Ownership is verified using database state
- Task lifecycle:
  - OPEN
  - IN_PROGRESS
  - DONE

This mirrors how task ownership works in real Jira-style systems.

---

## Task Management Features

- Create tasks
- Assign tasks (ADMIN only)
- Update task details
- Update task status
- Soft delete tasks
- Paginated task listing

---

## Soft Delete Strategy

Tasks are never physically removed from the database.

Instead:
- A `deleted_at` timestamp is used
- Deleted tasks do not appear in GET APIs
- Deleted tasks cannot be updated
- DELETE returns 404 if the task does not exist

This preserves data integrity and auditability.

---

## Database Design (Critical Part)

The database is responsible for enforcing core rules, not just the application code.

Key constraints:
- Unique constraint on user email
- Foreign key constraint on task assignee
- ENUM-based task status
- Indexes on frequently queried columns
- Soft delete handled at schema level

These constraints protect the system from race conditions and invalid data states.

---

## Database Schema

CREATE DATABASE IF NOT EXISTS mini_jira;
USE mini_jira;

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role ENUM('USER','ADMIN') NOT NULL DEFAULT 'USER',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tasks (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status ENUM('OPEN','IN_PROGRESS','DONE') NOT NULL DEFAULT 'OPEN',
    assignee_id INT NULL,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (assignee_id) REFERENCES users(id) ON DELETE SET NULL
);

CREATE INDEX idx_tasks_assignee ON tasks(assignee_id);
CREATE INDEX idx_tasks_status ON tasks(status);
CREATE INDEX idx_tasks_deleted ON tasks(deleted_at);

---

## Pagination

Task listing APIs support pagination using query parameters:

?page=1&limit=10

Pagination is applied at the database level using LIMIT and OFFSET.

---

## Input Validation & Error Handling

- Empty task titles are rejected
- Invalid JSON payloads return proper error responses
- Unauthorized and forbidden actions are handled consistently
- REST-appropriate HTTP status codes are used

---

## Tech Stack

- Language: Go
- HTTP Server: net/http
- Database: MySQL
- Authentication: JWT
- No third-party web frameworks

---

## Running the Project

1. Install Go (version 1.22 or higher recommended)

2. Configure MySQL credentials in:
   db/mysql.go

3. Install dependencies:
   go mod tidy

4. Start the server:
   go run main.go

The server runs on:
http://localhost:8080

---

## API Overview

Authentication:
- POST /register
- POST /login

Users (ADMIN only):
- GET /users

Tasks:
- POST   /tasks
- GET    /tasks
- PUT    /tasks/{id}
- PATCH  /tasks/{id}
- DELETE /tasks/{id}

---

## Testing

The APIs were tested using:
- Postman
- PowerShell (Invoke-WebRequest)
- curl.exe

Test scenarios include:
- Authentication failures
- Role-based authorization
- Ownership violations
- Pagination correctness
- Soft delete behavior
- Validation errors

---

## Design Decisions

- Database constraints are used to enforce correctness
- Primary keys are immutable and never reused
- No public users API for security reasons
- Middleware is implemented using http.Handler (idiomatic Go)
- Business rules are validated before database updates

---

## Future Improvements

- Database migrations
- Externalized configuration using environment variables
- Containerization using Docker and Docker Compose
- Observability and structured logging

---

## Project Status

✔ Feature complete  
✔ Database integrity enforced  
✔ Tested  
✔ Interview-ready  

---

## License

MIT
