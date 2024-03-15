package tests

import (
	"os"
	"service-employee/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	os.Setenv("POSTGRES_HOST", "127.0.0.1")
	os.Setenv("POSTGRES_DB", "apigateway")
	os.Setenv("POSTGRES_USER", "habi")
	os.Setenv("POSTGRES_PASSWORD", "habi123")
}

func TestConnectPostgreSQL(t *testing.T) {
	db := config.GetPostgresDatabase()

	assert.NotNil(t, db)
}