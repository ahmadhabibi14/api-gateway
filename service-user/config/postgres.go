package config

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

var (
  once sync.Once
  postgresDB *sql.DB
)

func NewPostgresDatabase() {
  once.Do(func() {
    LoadEnv() // prevent unloaded .env file
    PostgresHost := os.Getenv("POSTGRES_HOST")
    PostgresDbName := os.Getenv("POSTGRES_DB")
    PostgresUser := os.Getenv("POSTGRES_USER")
    PostgresPassword := os.Getenv("POSTGRES_PASSWORD")

    PostgresInfo := fmt.Sprintf(
      "host=%v user=%v password=%v dbname=%v sslmode=disable",
      PostgresHost, PostgresUser, PostgresPassword, PostgresDbName,
    )

    db, err := sql.Open("postgres", PostgresInfo)
    if err != nil {
      panic(err)
    }

    db.SetMaxIdleConns(10)
    db.SetMaxOpenConns(100)
    db.SetConnMaxIdleTime(5 * time.Minute)
    db.SetConnMaxLifetime(60 * time.Minute)

    postgresDB = db
  })
}

func NewPostgresContext() (context.Context, context.CancelFunc) {
  return context.WithTimeout(context.Background(), 10*time.Second)
}

func GetPostgresDatabase() *sql.DB {
  return postgresDB
}