# Mini Jira Backend (Go + MySQL)

Mini Jira is a backend-only implementation of a Jira-like issue tracking system.
The project is built using pure Go (`net/http`) and MySQL, without using any web frameworks.

The focus of this project is correct backend design, authentication and authorization,
data integrity, and real-world API behavior.

---

## Project Goals

- Build a production-style backend using standard Go libraries
- Implement role-based access control and ownership rules
- Follow real Jira-like workflows instead of shortcuts
- Handle edge cases correctly (validation, authorization, deletion)
- Keep the architecture simple, readable, and extensible

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
- USER and ADMIN can register and log in
- USER can only see and modify tasks assigned to them
- ADMIN can view and manage all tasks
- ADMIN-only access to users list

---

## Task Workflow (Jira-like)

- Tasks are created **unassigned**
- Only ADMIN can assign a task to a user
- A USER can modify a task only after it is assigned to them
- Ownership is enforced using database state, not request body
- Task lifecycle:
  - OPEN
  - IN_PROGRESS
  - DONE

---

## Task Management Features

- Create tasks
- Assign tasks (ADMIN only)
- Update tasks (ownership enforced)
- Update task status
- Soft delete tasks
- Paginated task listing

---

## Soft Delete Strategy

Tasks are never permanently removed from the database.

Instead:
- A `deleted_at` timestamp is used
- Deleted tasks:
  - Do not appear in GET APIs
  - Cannot be updated
  - Preserve historical data
- DELETE returns:
  - 200 for successful delete
  - 404 if task does not exist or is already deleted

---

## Pagination

Task listing APIs support pagination using query parameters:

?page=1&limit=10

- Pagination is applied at database level using LIMIT and OFFSET
- Defaults are applied if parameters are missing or invalid

---

## Input Validation & Error Handling

- Empty task titles are rejected
- Invalid JSON payloads return proper errors
- Unauthorized and forbidden actions are handled consistently
- HTTP status codes follow REST conventions

---

## Tech Stack

- Language: Go
- HTTP: net/http
- Database: MySQL
- Authentication: JWT
- No third-party web frameworks

---

## Database Schema

CREATE DATABASE mini_jira;
USE mini_jira;

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role ENUM('USER','ADMIN') DEFAULT 'USER'
);

CREATE TABLE tasks (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status ENUM('OPEN','IN_PROGRESS','DONE') DEFAULT 'OPEN',
    assignee_id INT,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    FOREIGN KEY (assignee_id) REFERENCES users(id) ON DELETE SET NULL
);

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

## API Endpoints

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

The API was tested using:
- Postman
- PowerShell (Invoke-WebRequest)
- curl.exe

Test cases covered:
- Authentication failures
- Role-based authorization
- Ownership violations
- Pagination
- Soft delete behavior
- Validation errors

---

## Design Decisions

- Primary keys are immutable and never reused
- No public users API for security reasons
- Middleware is implemented using http.Handler (idiomatic Go)
- Ownership checks rely on database state
- No hidden side effects in request payloads

---




