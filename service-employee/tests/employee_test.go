package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http/httptest"
	"service-employee/controller"
	"service-employee/model"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func init() {
	err := godotenv.Load("../../.env.dev")
	if err != nil {
		log.Println("Cannot load .env file")
	}
}

func TestCreateEmployee(t *testing.T) {
	app := fiber.New()
	app.Post("/employee", controller.CreateEmployee)

	payload := model.Employee{
		Name: "Gojo Satoru",
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		t.Error(`failed to convert JSON data`)
	}

	access_token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNhdG9ydTk5QHByb3Rvbi5tZSJ9.8tpN7S9LtxvzfIHzchb4QsFJcFm8K1f0rEzh2lgCThM"
	req := httptest.NewRequest(fiber.MethodPost, "/employee", bytes.NewBuffer(jsonData))
	req.Header.Set("access_token", access_token)

	resp, err := app.Test(req, 1000000000000)
	if err != nil {
		t.Error(err.Error())
	}

	assert.Equal(t, fiber.StatusCreated, resp.StatusCode, "login successful")

	body, _ := io.ReadAll(resp.Body)
	t.Logf(`Response: %v`, string(body))
}