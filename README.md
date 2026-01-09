# Mini Jira Backend (Go + MySQL)

Mini Jira is a backend-only implementation of a Jira-like issue tracking system
built using **pure Go (net/http)** and **MySQL**, without using any web frameworks.

The project focuses on **backend fundamentals**, including authentication,
authorization, ownership enforcement, and database integrity.

---

## Why this project

This project was built to demonstrate:
- How real-world backend systems enforce rules
- Clean separation between HTTP, business logic, and database layers
- Secure handling of users, roles, and tasks
- Correct use of database constraints instead of relying only on application code

---

## Core Features

### Authentication & Authorization
- User registration and login
- Password hashing
- JWT-based authentication
- Middleware-based request protection

### Roles
- USER
- ADMIN

### Role-Based Access Control
- USER can view and modify only tasks assigned to them
- ADMIN can manage all tasks
- ADMIN-only access to users list

---

## Jira-like Task Workflow

- Tasks are created **unassigned**
- Only ADMIN can assign a task to a user
- A USER can update a task only after it is assigned to them
- Ownership is enforced using **database state**, not request payloads
- Task lifecycle:
  - OPEN
  - IN_PROGRESS
  - DONE

---

## Task Management

- Create tasks
- Assign tasks (ADMIN only)
- Update tasks (ownership enforced)
- Update task status
- Soft delete tasks
- Paginated task listing

---

## Database Design (Important)

Database constraints are used to enforce correctness:

- Unique constraint on user email
- Foreign key constraint on task assignee
- ENUM-based task status
- Indexes for frequently queried columns
- Soft delete using `deleted_at` timestamp

These constraints ensure data integrity even in concurrent or failure scenarios.

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
- Invalid JSON payloads return appropriate errors
- Unauthorized and forbidden actions are handled consistently
- REST-appropriate HTTP status codes are used

---

## Tech Stack

- Language: Go
- HTTP: net/http
- Database: MySQL
- Authentication: JWT
- No web frameworks

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

APIs were tested using:
- Postman
- PowerShell (Invoke-WebRequest)
- curl.exe

Test coverage includes:
- Authentication failures
- Role-based authorization
- Ownership enforcement
- Pagination behavior
- Soft delete behavior
- Validation errors

---

## Design Decisions

- Database constraints are used to enforce critical rules
- Primary keys are immutable and never reused
- No public users API for security reasons
- Middleware is implemented using http.Handler (idiomatic Go)
- Ownership is enforced before any state change

---

## Future Improvements

- Database migrations
- Externalized configuration using environment variables
- Improved observability and metrics

---


MIT
