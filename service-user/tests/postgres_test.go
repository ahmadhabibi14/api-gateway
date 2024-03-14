package tests

import (
	"service-employee/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectPostgreSQL(t *testing.T) {
	db := config.GetPostgresDatabase()

	assert.NotNil(t, db)
}