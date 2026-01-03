# 🧩 Mini-JIRA Backend (Go)

A lightweight backend project built in **Go (without any framework)** to understand
clean backend architecture, MySQL integration, and real-world feature design.

This project implements:
- User management
- JWT-based authentication
- Core Mini-JIRA Task module

---

## 🏗️ Architecture Overview

HTTP Request  
↓  
Handler (HTTP + JSON)  
↓  
Service (Business Rules)  
↓  
Repository (Interface)  
↓  
MySQL Database  

---

## 📁 Project Structure

mini-jira/  
│  
├── main.go  
├── go.mod  
├── go.sum  
│  
├── model/  
│   ├── user.go  
│   └── task.go  
│  
├── handler/  
│   ├── user_handler.go  
│   ├── auth_handler.go  
│   └── task_handler.go  
│  
├── service/  
│   ├── user_service.go  
│   └── task_service.go  
│  
├── repository/  
│   ├── user_repository.go  
│   ├── mysql_user_repository.go  
│   ├── task_repository.go  
│   ├── mysql_task_repository.go  
│   └── db.go  
│  
├── middleware/  
│   └── auth_middleware.go  
│  
└── auth/  
    └── jwt.go  

---

## 🗄️ Database Schema

### Users Table

CREATE TABLE users (  
    id INT AUTO_INCREMENT PRIMARY KEY,  
    name VARCHAR(100) NOT NULL,  
    email VARCHAR(150) NOT NULL UNIQUE,  
    password VARCHAR(255) NOT NULL,  
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP  
);  

### Tasks Table

CREATE TABLE tasks (  
    id INT AUTO_INCREMENT PRIMARY KEY,  
    title VARCHAR(200) NOT NULL,  
    description TEXT,  
    status VARCHAR(20) NOT NULL,  
    assignee_id INT NULL,  
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP  
);  

---

## 🔐 Authentication (JWT)

POST /login  
↓  
Validate email & password  
↓  
Generate JWT token  
↓  
Client stores token  
↓  
Token sent in Authorization header  

Header format:  
Authorization: Bearer <JWT_TOKEN>  

Note: JWT is currently applied **only to user APIs**.  
Task APIs are intentionally kept **JWT-free**.

---

## 🧪 API Endpoints

### User APIs

Create User  
POST /users  

Request Body:  
{  
  "name": "Krushna",  
  "email": "abhishek@gmail.com",  
  "password": "123456"  
}  

Login  
POST /login  

Request Body:  
{  
  "email": "abhishek@gmail.com",  
  "password": "123456"  
}  

---

## 🧩 Task Module (Mini-JIRA Core)

### Implemented Features

- Create task  
- View all tasks  
- View task by ID  
- Assign task to user  
- Update task status  
- Delete task  

### Task Status Rules

OPEN → IN_PROGRESS → DONE  

Invalid transitions are blocked at **service layer**.

---

### Task APIs

Create Task  
POST /tasks  

Request Body:  
{  
  "title": "Fix login bug",  
  "description": "Login fails on invalid password"  
}  

Get All Tasks  
GET /tasks-list  

Get Task by ID  
GET /tasks/{id}  

Assign Task  
PUT /tasks/assign/{taskId}  

Request Body:  
{  
  "assignee_id": 2  
}  

Update Task Status  
PUT /tasks/status/{taskId}  

Request Body:  
{  
  "status": "IN_PROGRESS"  
}  

Delete Task  
DELETE /tasks/delete/{taskId}  

---

## ▶️ How to Run

Clone repository:  
git clone https://github.com/<your-username>/mini-jira.git  
cd mini-jira  

Install dependencies:  
go mod tidy  

Configure database in:  
repository/db.go  

Run server:  
go run main.go  

Server runs on:  
http://localhost:8080  

---

## 🧠 Key Learnings

- Layered backend architecture  
- Repository pattern with interfaces  
- MySQL integration in Go  
- JWT authentication without frameworks  
- Business rules enforced at service layer  
- Nullable DB fields handling  
- Manual HTTP routing  

---

## 🔮 Future Enhancements

- Protect task APIs with JWT  
- Task comments & priorities  
- Role-based access control  
- Pagination & filters  
- gRPC-based service  
- Concurrency & channels demo  

---

Built with ❤️ while learning backend engineering the right way.
