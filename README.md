Description:

This repository contains a Go-lang REST API project built using popular frameworks such as GORM, Gorilla Mux, and Go JWT. The project provides a robust foundation for developing scalable and secure RESTful APIs using Go.

Features:

1. GORM: A powerful ORM (Object-Relational Mapping) library for Go, enabling seamless database interactions. GORM simplifies data persistence by providing an intuitive API for querying, creating, updating, and deleting records in the database.
2. Gorilla Mux: A flexible and powerful HTTP request router and dispatcher for Go. Gorilla Mux allows you to define and handle routes, perform URL parameter matching, and handle various HTTP methods with ease.
3. Go JWT: A JSON Web Token (JWT) implementation in Go. Go JWT enables secure authentication and authorization by generating and validating JWTs, allowing you to protect your API endpoints and manage user sessions effectively.


Key Components:

1. API Routes: Define your API endpoints and corresponding handlers using Gorilla Mux's powerful routing capabilities. Create routes for various HTTP methods like GET, POST, PUT, DELETE, etc., and map them to appropriate handler functions.
2. Database Integration: Utilize GORM to connect to your preferred database (MySQL, PostgreSQL, SQLite, etc.) and perform database operations, including querying, creating, updating, and deleting records.
3. Authentication and Authorization: Implement secure user authentication and authorization using Go JWT. Generate and validate JWTs to authenticate API requests and protect sensitive routes and resources.


Usage:

1. Clone the repository: git clone https://github.com/muhardiansyah15/go-rest-api.git
2. Install the project dependencies: go mod download
3. Configure your database connection in the project's configuration files.
4. Run the API server: go run main.go
