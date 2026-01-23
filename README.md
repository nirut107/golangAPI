# Golang Backend API

A backend REST API built with Golang following clean architecture principles, focusing on authentication, middleware, and secure API design.

This project demonstrates real-world backend development using Go’s standard library, JWT authentication, and PostgreSQL.

---

## 📌 Features
- RESTful API using net/http
- Clean architecture (handler / service / repository)
- JWT Authentication (HS256, HMAC)
- Secure password hashing with bcrypt
- Authentication middleware
- Public & Protected routes separation
- JSON request/response handling
- PostgreSQL database integration
- Environment-based configuration
- Logging middleware

---


## 🏗 Project Structure
```python
.
├── app/            # Application container / dependency wiring
├── handler/        # HTTP handlers (login, register, users)
├── middleware/     # Auth, JWT, logging middleware
├── model/          # Data models (User, etc.)
├── repository/     # Database access layer (PostgreSQL)
├── service/        # Business logic layer
├── routes/         # Route setup (public & protected routes)
├── main.go         # Application entry point
├── go.mod
├── go.sum
└── README.md
```
## 📌 Folder Responsibilities

### app/ 
จัดการ dependency injection และรวม handler ทั้งหมดไว้ใน application container
ช่วยให้ main.go สะอาดและขยายง่าย

### handler/
รับ HTTP request / response
ไม่ทำ business logic โดยตรง

### service/
business logic
ตรวจสอบเงื่อนไขและเรียก repository

### repository/
ติดต่อฐานข้อมูล (PostgreSQL)
แยก logic DB ออกจาก service

### middleware/
Authentication (JWT, HMAC), Logging
ใช้ http.Handler ตาม Go idiom

### routes/
แยก public และ protected routes
wrap middleware อย่างเป็นระบบ

### model/
struct สำหรับข้อมูล เช่น User

### main.go
bootstrap แอป:

- load env

- connect DB

- init app

- start HTTP server

---

## 🔐 Authentication Flow
### Register
```
POST /register
```

- Hash password using bcrypt

- Store user in database

- Does NOT auto-login (security best practice)

### Login
```
POST /login
```


- Validate username & password
- Generate JWT (HS256)
- Return token via HttpOnly cookie

### Protected Routes
```
GET /users
```


- Require valid JWT cookie
- Verified via AuthMiddleware
- User identity is stored in request context

## 🔒 Security Practices

- Passwords are hashed using bcrypt
- JWT signed with HMAC (HS256)
- Signing method is strictly validated to prevent alg=none attacks
- JWT stored in HttpOnly Cookie
- Environment variables for secrets

## ⚙️ Environment Variables

## Create a .env file:
```
DATABASE_URL=postgres://user:password@localhost:5432/dbname
JWT_SECRET_KEY=your-secret-key
```


## 🛠 How to Run

```golang
 go mod tidy
 go run main.go
 ```


---

### Server will start at:
```
http://localhost:8080
```

## 🧪 Example Request
### Login
```
POST /login
Content-Type: application/json


{
  "username": "testuser",
  "password": "password123"
}
```

## 🛠 Technologies Used

- Golang (net/http)
- PostgreSQL
- JWT (golang-jwt)
- bcrypt
- Docker (optional)
- Git & GitHub

## 📌 Purpose of This Project

#### This project was built to:
- Practice real-world backend development with Go
- Demonstrate authentication & security best practices
- Serve as a portfolio project for Backend Engineer (Golang) positions

## 👤 Author

    Nirut Somrod
    Backend / Software Engineer
    GitHub: https://github.com/nirut107

