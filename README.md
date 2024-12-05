
# SQLC API Go

This project leverages the **Echo** framework and **SQLC** to build a Go API integrated with a PostgreSQL database.

## Versioning Table

| Date       |    Version    |                           Description                           |     Commit    | 
|------------|---------------|-----------------------------------------------------------------|---------------| 
| 2024-12-05 | 0.1.1         | x |   `xxxxxxx`   | 
| 2024-12-04 | 0.1.0 :tada:  | Initial project setup with basic structure and dependencies.    |   `xxxxxxx`   | 

---
---

## Project Structure

```plaintext
sqlc-api-go/
├── cmd/main.go             # Application entry point, initializes and starts the server
├── internal/user/          # Business logic and data access layer for user-related operations
│   ├── handler.go          # Handles HTTP requests and responses for user routes
│   ├── service.go          # Contains business rules and logic for user operations
│   └── repository.go       # Implements data access methods for user-related queries
├── database/               # Database interaction and generated code
│   ├── models.go           # Structs representing database entities
│   └── query.sql.go        # Auto-generated code by SQLC for database queries
├── config/
│   └── database.go         # Configuration for connecting to the database
├── .env                    # Environment variables for sensitive data (e.g., credentials, ports)
├── query.sql               # SQL queries used by SQLC to generate Go code
├── schema.sql              # SQL script defining the database schema
├── go.mod                  # Declares dependencies and module information for the project
└── go.sum                  # Dependency lock file, ensuring version consistency

```

---

## Dependencies

### Libraries Used

- **[Echo](https://echo.labstack.com/docs):** Web framework for routing and middleware management.
- **[SQLC](https://docs.sqlc.dev/en/latest/overview/install.html):** Generates Go code from SQL files, ensuring strongly typed queries.
- **[PGX Pool](https://github.com/jackc/pgx):** PostgreSQL database connection library with connection pooling support.
- **[Godotenv](https://github.com/joho/godotenv):** Loads environment variables from `.env` files, simplifying development and production setup.

---

## Installation from scratch

1. **Set Up Go Environment**:
   - Install Go and ensure `GOPATH` and `GOROOT` variables are properly configured.

2. **Initialize the Project**:
   ```bash
   mkdir -p $GOPATH/src/github.com/MaluMSiza/sqlc-api-go
   cd $GOPATH/src/github.com/MaluMSiza/sqlc-api-go
   code .
   go mod init github.com/MaluMSiza/sqlc-api-go
   ```

3. **Install Dependencies**:
   ```bash
   go get -u github.com/labstack/echo/v4
   go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
   go get github.com/jackc/pgx/v5/pgxpool
   go get github.com/joho/godotenv
   ```

4. **Prepare the Database**:
   - Use `schema.sql` to create the database schema.
   - Configure the connection settings in the `.env` file.

5. **Run the Server**:
   ```bash
   cd cmd
   go run main.go
   ```

---


