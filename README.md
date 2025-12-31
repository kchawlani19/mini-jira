# 🧩 Mini-JIRA Backend (Go)

A backend application inspired by JIRA, built using **pure Go (net/http)** with a **clean layered architecture**, **JWT authentication**, and **MySQL persistence**.

---

## 🚀 Features Implemented

### 👤 User Management
- User signup (public)
- Login using email & password
- Password hashing using bcrypt
- JWT generation on login
- Get all users (protected)
- Get user by ID (protected)
- Update user (protected)
- Delete user (protected)

### 🔐 Authentication & Security
- JWT-based stateless authentication
- Middleware-based route protection
- Public vs protected API separation
- Secure password storage (bcrypt)

### 🗄️ Database
- MySQL integration
- Repository pattern using interfaces

---

## 🏗️ Architecture

```
HTTP Request
   ↓
Handler (HTTP + JSON)
   ↓
Service (Business Logic & Rules)
   ↓
Repository (Interface)
   ↓
MySQL Database
```

---

## 📂 Project Structure

```
mini-jira/
│
├── main.go
├── go.mod
├── go.sum
│
├── model/
│   └── user.go
│
├── handler/
│   ├── user_handler.go
│   └── auth_handler.go
│
├── service/
│   └── user_service.go
│
├── repository/
│   ├── user_repository.go
│   ├── mysql_user_repository.go
│   └── db.go
│
├── middleware/
│   └── auth_middleware.go
│
└── auth/
    └── jwt.go
```

---

## 🗄️ Database Schema

```
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(150) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

---

## 🔐 Authentication Flow (JWT)

```
POST /login
   ↓
Validate email + password
   ↓
Generate JWT token
   ↓
Client stores token
   ↓
Token sent in Authorization header
```

Header format:
```
Authorization: Bearer <JWT_TOKEN>
```

---

## 🧪 API Endpoints

### Public
POST /users  
POST /login  

### Protected (JWT)
GET /users  
GET /users/{id}  
PUT /users/{id}  
DELETE /users/{id}  

---

## ▶️ How to Run

```
git clone https://github.com/<your-username>/mini-jira.git
cd mini-jira
go mod tidy
go run main.go
```

Server runs on:
```
http://localhost:8080
```

---

## 🔮 Upcoming
- Task / Issue management
- Task assignment
- Status transitions (OPEN → IN_PROGRESS → DONE)

---

## ✨ Author
Built with ❤️ while learning backend engineering the right way.
