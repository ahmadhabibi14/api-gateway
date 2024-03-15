package tests

import (
	"io"
	"net/http/httptest"
	"os"
	"service-user/controller"
	"service-user/middleware"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func init() {
	os.Setenv("POSTGRES_HOST", "127.0.0.1")
	os.Setenv("POSTGRES_DB", "apigateway")
	os.Setenv("POSTGRES_USER", "habi")
	os.Setenv("POSTGRES_PASSWORD", "habi123")
}


func TestUserAuth(t *testing.T) {
	app := fiber.New()
	app.Get("/user/auth", middleware.Authentication, controller.Auth)

	t.Run("Authenticated", func(t *testing.T) {
		// obtain this token by login
		access_token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNhdG9ydTk5QHByb3Rvbi5tZSJ9.8tpN7S9LtxvzfIHzchb4QsFJcFm8K1f0rEzh2lgCThM"
		req := httptest.NewRequest(fiber.MethodGet, "/user/auth", nil)
		req.Header.Set("access_token", access_token)

		resp, err := app.Test(req, 1000000000000)
		if err != nil {
			t.Error(err.Error())
		}

		assert.Equal(t, fiber.StatusOK, resp.StatusCode, "user authenticated")

		body, _ := io.ReadAll(resp.Body)
		t.Logf(`Response: %v`, string(body))
	})

	t.Run("InvalidAccessToken", func(t *testing.T) {
		access_token := "ini-error-haha"
		req := httptest.NewRequest(fiber.MethodGet, "/user/auth", nil)
		req.Header.Set("access_token", access_token)

		resp, err := app.Test(req, 1000000000000)
		if err != nil {
			t.Error(err.Error())
		}

		assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode, "invalid access token")

		body, _ := io.ReadAll(resp.Body)
		t.Logf(`Response: %v`, string(body))
	})
}