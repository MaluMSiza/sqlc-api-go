package config

import (
    "context"
    "fmt"
    "log"
    "os"

    "github.com/golang-migrate/migrate/v4"
    _ "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"
    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/joho/godotenv"
)

var DBPool *pgxpool.Pool

func InitDB() {
    envFilePath := ".env" 
    if os.Getenv("DOCKER") == "true" {
        envFilePath = "/app/.env" 
    }

    if err := godotenv.Load(envFilePath); err != nil {
        log.Fatal("Error loading .env file")
    }

    connString := fmt.Sprintf(
        "postgresql://%s:%s@%s:%s/%s?sslmode=disable",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_NAME"),
    )

    applyMigrations(connString)

    var err error
    DBPool, err = pgxpool.New(context.Background(), connString)
    if err != nil {
        log.Fatal("Error connecting to the database:", err)
    }

    log.Println("Database connection successful!")
}

func CloseDB() {
    DBPool.Close()
}

func applyMigrations(connString string) {
    migrationsPath := "file://./migrations"

    if os.Getenv("DOCKER") == "true" {
        migrationsPath = "file:///app/migrations"
    }

    m, err := migrate.New(migrationsPath, connString)
    if err != nil {
        log.Fatalf("Could not start migrations: %v", err)
    }

    if err := m.Up(); err != nil && err != migrate.ErrNoChange {
        log.Fatalf("Could not apply migrations: %v", err)
    }

    log.Println("Migrations applied successfully!")
}
