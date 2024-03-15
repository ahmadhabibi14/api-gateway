package controller

import (
	"fmt"
	"net/http"
	"service-employee/config"
	"service-employee/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var user_uri string = "http://service-user:3001/user"

type WebResponse struct {
	Code int
	Status string
	Data interface{}
}

func CreateEmployee(c *fiber.Ctx) error {
	db := config.GetPostgresDatabase()
	var requestBody model.Employee

	c.BodyParser(&requestBody)

	requestBody.Id = uuid.New().String()

	access_token := c.Get("access_token")
	if len(access_token) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(WebResponse{
			Code: 401,
			Status: "BAD_REQUEST",
			Data: "Invalid token: Access token missing",
		})
	}

	req, err := http.NewRequest("GET", user_uri + "/auth", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		panic(err)
	}

	// Set headers
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("access_token", access_token)

	// Send the request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		panic(err)
	}
	defer resp.Body.Close()

	if resp.Status != "200 OK" {
		return c.Status(fiber.StatusBadRequest).JSON(WebResponse{
			Code: 401,
			Status: "BAD_REQUEST",
			Data: "invalid token",
		})
	}

	ctx, cancel := config.NewPostgresContext()
	defer cancel()

	query := `INSERT INTO employee (id, name) VALUES ($1, $2)`
	_, err = db.ExecContext(ctx, query,
		requestBody.Id,
		requestBody.Name,
	)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(WebResponse{
			Code: 401,
			Status: "BAD_REQUEST",
			Data: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(WebResponse{
		Code: 201,
		Status: "OK",
		Data: requestBody,
	})
}