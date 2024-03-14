package config

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var postgresDB *sql.DB

func init() {
   PostgresHost := os.Getenv("POSTGRES_HOST")
   PostgresPort := os.Getenv("POSTGRES_HOST")
   PostgresDB := os.Getenv("POSTGRES_HOST")
   PostgresUser := os.Getenv("POSTGRES_HOST")
   PostgresPassword := os.Getenv("POSTGRES_HOST")

   dconn := fmt.Sprintf(
      "postgresql://%s:%s@%s:%s/%s?sslmode=disable",
      PostgresUser, PostgresPassword, PostgresHost, PostgresPort, PostgresDB,
   )

   db, err := sql.Open("postgres", dconn)
   if err != nil {
      panic(err)
   }

   postgresDB = db
}

func NewPostgresContext() (context.Context, context.CancelFunc) {
   return  context.WithTimeout(context.Background(), 10*time.Second)
}

func GetPostgresDatabase() *sql.DB {
   return postgresDB
}