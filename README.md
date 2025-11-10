
# News Feed System Backend

A lightweight social feed backend built with **Go Fiber**, **GORM**, and **PostgreSQL**.  This service provides APIs for user authentication (JWT), posting content, following/unfollowing users, and fetching a personalized feed.

## Demo
https://news-feed-system-backend-production.up.railway.app/

## Features

- User registration and login with JWT authentication  
- Create, read, and list posts with pagination  
- Follow and unfollow other users  
- Retrieve feed filtered by followed users  
- Database migration and ORM using GORM  
- Docker support for backend and PostgreSQL  
- CI/CD integration with GitHub Actions  
- CORS support for frontend connection  

## Tech

- Framework: Go Fiber v2
- ORM: Gorm
- Database: PostgreSQL
- Authentication: JWT
- Container: Docker & Docker Compose
- CI/CD: GitHub Actions


## Installation

### Prerequisites
- **Go** ≥ 1.25.1  
- **PostgreSQL** ≥ 15  
- **Docker** & **Docker Compose** (latest recommended)

### 1. Clone Project
```bash
git clone https://github.com/your-username/news-feed-system.git
cd news-feed-system
```
### 2. Setup Environment
```bash
cp .env.example .env
# Edit .env file with appropriate configuration
```
Example .env:

```bash
DATABASE_URL=postgres://postgres:postgres@postgres:5432/newsfeed?sslmode=disable
JWT_SECRET=secretkey12345
PORT=8080
```

### 3. Run Locally
```bash
go mod download
go run main.go
```

### 3. Run with Docker (Recommended)
```bash
docker-compose up -d --build
```

### 4. Access Application
```text
http://localhost:8080/api
```
