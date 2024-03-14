package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"
	"service-user/controller"
	"service-user/model"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic("Cannot load .env file")
	}
}

func TestUserRegister(t *testing.T) {
	app := fiber.New()
	app.Post("/user/register", controller.Register)

	t.Run("Success", func(t *testing.T) {
		payload := model.User{
			Email: "satoru99@proton.me",
			Password: "gojo12345678",
		}
		jsonData, err := json.Marshal(payload)
		if err != nil {
			t.Error(`failed to convert JSON data`)
		}

		req := httptest.NewRequest(fiber.MethodPost, "/user/register", bytes.NewBuffer(jsonData))
		resp, err := app.Test(req)
		if err != nil {
			t.Error(err.Error())
		}

		assert.Equal(t, fiber.StatusCreated, resp.StatusCode, "register successful")

		body, _ := io.ReadAll(resp.Body)
		t.Logf(`Response: %v`, string(body))
	})

	t.Run("Failed", func(t *testing.T) {
		payload := model.User{
			Email: "emailgakada@proton.me",
			Password: "bukanpassword",
		}
		jsonData, err := json.Marshal(payload)
		if err != nil {
			t.Error(`failed to convert JSON data`)
		}

		req := httptest.NewRequest(fiber.MethodPost, "/user/register", bytes.NewBuffer(jsonData))
		resp, err := app.Test(req)
		if err != nil {
			t.Error(err.Error())
		}

		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode, "register failed")

		body, _ := io.ReadAll(resp.Body)
		t.Logf(`Response: %v`, string(body))
	})
}