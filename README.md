# 🧩 Mini-JIRA Backend (Go)

A backend application inspired by JIRA, built using **pure Go (net/http)** with **clean layered architecture**, **JWT authentication**, and **MySQL persistence**.

This project is built incrementally to demonstrate **real backend engineering concepts** like authentication, authorization, repository abstraction, and business-logic separation.

---

## 🚀 Features Implemented (Current)

### 👤 User Management
- Create user (Signup – public)
- Login with email & password
- Password hashing using **bcrypt**
- JWT generation on login
- Get all users (protected)
- Get user by ID (protected)
- Update user (protected)
- Delete user (protected)

### 🔐 Authentication & Security
- JWT-based stateless authentication
- Middleware-based route protection
- Secure password storage (bcrypt)
- Public vs protected route separation

### 🗄️ Database
- MySQL integration
- Repository pattern using interfaces
- Easy switch between in-memory & MySQL implementations

---

## 🏗️ Architecture

The project follows a **layered architecture**:

