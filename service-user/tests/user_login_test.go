package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"
	"os"
	"service-user/config"
	"service-user/controller"
	"service-user/model"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func init() {
	os.Setenv("POSTGRES_HOST", "127.0.0.1")
	os.Setenv("POSTGRES_DB", "apigateway")
	os.Setenv("POSTGRES_USER", "habi")
	os.Setenv("POSTGRES_PASSWORD", "habi123")
	config.NewPostgresDatabase()
}

func TestUserLogin(t *testing.T) {
	app := fiber.New()
	app.Post("/user/login", controller.Login)

	payload := model.User{
		Email: "satoru99@proton.me",
		Password: "gojo12345678",
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		t.Error(`failed to convert JSON data`)
	}

	req := httptest.NewRequest(fiber.MethodPost, "/user/login", bytes.NewBuffer(jsonData))
	resp, err := app.Test(req, 1000000000000)
	if err != nil {
		t.Error(err.Error())
	}

	assert.Equal(t, fiber.StatusOK, resp.StatusCode, "login successful")

	body, _ := io.ReadAll(resp.Body)
	t.Logf(`Response: %v`, string(body))
}