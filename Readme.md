# GoBank - Bank Management System

**GoBank** is a bank management system built using Go, leveraging Gorilla Mux for routing, JWT for authentication, PostgreSQL for the database, and Docker for containerization. This project provides basic bank management functionalities such as account creation, balance checking, and CRUD operations on user data.

## Tech Stack

- **Go (v1.24.6)**
- **Gorilla Mux** for routing
- **PostgreSQL** as the database
- **JWT** for authentication
- **Docker** for containerization

---

## Features

- Create, read, update, and delete (CRUD) bank accounts.
- JWT-based authentication for secure API access.
- Dockerized setup for development and production environments.
- RESTful APIs to interact with the bank system.

---

## Requirements

- **Go 1.24.6** (or compatible version)
- **Docker** (for containerization)

---

## Installation

### 1. Clone the repository:

```bash
git clone https://github.com/yourusername/gobank.git
cd gobank
```

### 2. Build the project:

The project uses a Makefile to manage builds and dependencies.

```bash
make build
```

This will compile the Go code and output an executable `gobank` in the `bin/` directory.

### 3. Run the project:

After building, you can run the application:

```bash
make run
```

This will start the server. By default, it will be available on `localhost:8080`.

---

## Docker Setup

### 1. Build the Docker containers:

```bash
docker-compose build
```

### 2. Start the Docker containers:

```bash
docker-compose up
```

This will start both the application and the PostgreSQL container. The application will be available on `localhost:8080`.

---

## API Endpoints

The GoBank API exposes the following endpoints:

- **POST /account**: Create a new bank account.
- **GET /account/{id}**: Get the details of a specific account.
- **PUT /account/{id}**: Update account details.
- **DELETE /account/{id}**: Delete an account.
- **POST /auth/login**: Authenticate user and get a JWT token.

---

## Testing

To run the tests for the application:

```bash
make test
```

This will run all unit tests in the project.

---

## Example Usage

### 1. Create an Account:

```bash
curl -X POST http://localhost:8080/account \
     -H "Content-Type: application/json" \
     -d '{"firstName": "John", "lastName": "Doe", "number": 123456789, "balance": 1000.50}'
```

### 2. Get Account Details:

```bash
curl -X GET http://localhost:8080/account/{id}
```

Replace `{id}` with the actual account ID.

### 3. Update Account:

```bash
curl -X PUT http://localhost:8080/account/{id} \
     -H "Content-Type: application/json" \
     -d '{"balance": 1500.75}'
```

### 4. Delete Account:

```bash
curl -X DELETE http://localhost:8080/account/{id}
```

---

## Makefile Commands

- **build**: Compile the Go code into the `bin/gobank` executable.
- **run**: Run the application after building it.
- **test**: Run the unit tests for the application.
