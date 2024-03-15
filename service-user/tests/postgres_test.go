package tests

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	err := godotenv.Load("../../.env.dev")
	if err != nil {
		log.Println("Cannot load .env file")
	}
}

func TestConnectPostgreSQL(t *testing.T) {
	PostgresHost := os.Getenv("POSTGRES_HOST")
	PostgresDbName := os.Getenv("POSTGRES_DB")
	PostgresUser := os.Getenv("POSTGRES_USER")
	PostgresPassword := os.Getenv("POSTGRES_PASSWORD")
	PostgresInfo := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v sslmode=disable",
		PostgresHost, PostgresUser, PostgresPassword, PostgresDbName,
	)

	_, err := sql.Open("postgres", PostgresInfo)
	if err != nil {
		t.Error("error connect to PostgreSQL:", err.Error())
	}

	t.Log("Connected to PostgreSQL")
}