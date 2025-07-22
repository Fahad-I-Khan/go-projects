````markdown
# 🚀 Google OAuth2 Authentication with Go (Gin + GORM + Docker)

This project demonstrates how to implement **Google OAuth2 Login** using **Golang's Gin framework**, along with **PostgreSQL**, **Docker**, **JWT**, and **Swagger** for API documentation.

---

## 📚 Table of Contents

- [Features](#-features)
- [Technologies Used](#-technologies-used)
- [Setup Instructions](#-setup-instructions)
- [API Endpoints](#-api-endpoints)
- [Author](#-author)

---

## ✅ Features

- 🔐 Google OAuth2 Login
- 🧾 JWT Token Generation
- ⚙️ Protected Routes
- 🧵 GORM Integration with PostgreSQL
- 🔄 Docker Support for Backend and Database
- 📘 Swagger UI for API Testing

---

## 🛠️ Technologies Used

- [Go (Gin Framework)](https://gin-gonic.com/)
- [GORM](https://gorm.io/)
- [PostgreSQL](https://www.postgresql.org/)
- [Docker](https://www.docker.com/)
- [Swagger (Swaggo)](https://github.com/swaggo/swag)

---

## ⚙️ Setup Instructions

### 📌 Step 1: Clone the Repository

```bash
git clone https://github.com/Fahad-I-Khan/go-oauth2-gin.git
cd go-oauth2-gin
````

### 📦 Step 2: Install Dependencies

Make sure you have [Go](https://go.dev/dl/), [Docker](https://www.docker.com/), and [Git](https://git-scm.com/) installed.

Then install Go dependencies:

```bash
go mod tidy
go install github.com/swaggo/swag/cmd/swag@latest
```

### 📦 Step 3: Install Go Dependencies

Use the following `go get` commands to install required packages:

```bash
go get github.com/gin-gonic/gin                 # Web framework
go get github.com/lib/pq                        # PostgreSQL driver for database connection
go get gorm.io/gorm                             # ORM for PostgreSQL
go get gorm.io/driver/postgres                  # PostgreSQL driver for GORM
go get github.com/joho/godotenv                 # Load .env files
go get golang.org/x/oauth2                      # OAuth2 core package
go get golang.org/x/oauth2/google               # Google-specific OAuth2 config
go get github.com/swaggo/gin-swagger            # Swagger UI for Gin
go get github.com/swaggo/files                  # Swagger files handler
go get github.com/gin-contrib/cors              # CORS middleware
go install github.com/swaggo/swag/cmd/swag@latest # CLI to generate Swagger docs
```

Then tidy up:

```bash
go mod tidy
```

---

The API will run at:
📌 `http://localhost:8080`
📘 Swagger UI: `http://localhost:8080/swagger/index.html`

---

### 🐳 Step 4: Run with Docker

```bash
docker-compose up -d go_db
docker-compose build
docker-compose up -d go-app
```

## 📮 API Endpoints

| Method | Endpoint                | Description                      |
| ------ | ----------------------- | -------------------------------- |
| GET    | `/api/v1/auth/login`    | Redirects to Google login        |
| GET    | `/api/v1/auth/callback` | Google callback handler          |
<!-- | GET    | `/api/v1/protected`     | Protected route (requires token) | -->

---
