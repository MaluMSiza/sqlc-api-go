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
    err := godotenv.Load("../.env") 
    if err != nil {
        log.Fatal("Erro ao carregar arquivo .env")
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
        log.Fatal("Erro ao conectar no banco:", err)
    }
    log.Println("Conexão com o banco bem-sucedida!")
}

func CloseDB() {
    DBPool.Close()
}
