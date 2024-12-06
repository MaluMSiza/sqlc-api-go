package config

import (
    "context"
    "fmt"
    "log"
    "os"

    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/joho/godotenv"
)

var DBPool *pgxpool.Pool

func InitDB() {
    err := godotenv.Load(".env") 
    if err != nil {
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

    DBPool, err = pgxpool.New(context.Background(), connString)
    if err != nil {
        log.Fatal("Error connecting to the database:", err)
    }
    log.Println("Database connection successful!")
}

func CloseDB() {
    DBPool.Close()
}
