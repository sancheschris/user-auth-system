# User Authentication System

This project is a **User Authentication System** built with Go. It provides APIs for user registration and login, using MongoDB as the database. The project is designed with modularity and scalability in mind, following clean architecture principles.

---

## Features
- **User Registration**: Allows users to register with a username and password.
- **User Login**: Authenticates users and generates JWT tokens.
- **MongoDB Integration**: Stores user data securely in a MongoDB database.
- **Environment Configuration**: Uses `.env` files for managing sensitive credentials.
- **Integration Tests**: Includes integration tests using `testcontainers-go` for MongoDB.

---

## Project Structure
```
user-auth-system/
├── cmd/
│   └── api/                # Main entry point for the application
├── internal/
│   ├── auth/               # Authentication logic (service, repository, handlers)
│   ├── config/             # Configuration management
│   ├── database/           # Database connection logic
│   └── model/              # Data models
├── test/
│   └── integration/        # Integration tests
├── .env                    # Environment variables
└── README.md               # Project documentation
```


---

## Prerequisites
- Go 1.20 or later
- Docker (for running MongoDB in a container)
- MongoDB (if running locally)

---

## Setup Instructions

### 1. Clone the Repository
```bash
git clone https://github.com/sancheschris/user-auth-system.git
cd user-auth-system
```

### 2. Set Up Environment Variables
```
MONGO_USER=root
MONGO_PASSWORD=secret
MONGO_HOST=localhost
MONGO_PORT=27017
MONGO_DB=example_db
```

### 3. Run MongoDB

- **Run MongoDB Locally**: Ensure MongoDB is running on
`localhost:27017`
- **Run MongoDB in Docker**:
```bash 
docker run --rm --name mongo -d -p 27017:27017 mongo:latest
```

### 4. Install Dependencies
```bash
go mod tidy
```

### 5. Run the Application
```bash
go run cmd/api/main.go
```

- The server will start on `http://localhost:8080`

- **Endpoint**: `POST /register`
- **Request Body**:
```
{
  "username": "testuser",
  "password": "testpassword"
}
```
- **Response**:
```
{
  "message": "User registered successfully"
}
```

- **Endpoint**: `POST /login`
- **Request Body**:
```
{
  "username": "testuser",
  "password": "testpassword"
}
```

- **Response**:
```
{
  "token": "your-jwt-token"
}
```

### Running Tests

#### 1 Unit Tests
Run all unit tests:

`go test ./... -v`

### 2 Integration Tests
Integration tests use testcontainers-go to spin up a temporary MongoDB container. Run them with:

`go test ./test/integration -v`

