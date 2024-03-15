package controller

import (
	"service-user/helpers"
	"service-user/model"

	"service-user/config"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type WebResponse struct {
	Code int
	Status string
	Data interface{}
}

func Register(c *fiber.Ctx) error {
	var requestBody model.User

	db := config.GetPostgresDatabase()

	requestBody.Id = uuid.New().String()

	ctx, cancel := config.NewPostgresContext()
	defer cancel()

	c.BodyParser(&requestBody)

	query := `INSERT INTO users (id, email, password) VALUES ($1, $2, $3)`
	_, err := db.ExecContext(ctx, query,
		requestBody.Id,
		requestBody.Email,
		helpers.HashPassword([]byte(requestBody.Password)),
	)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(WebResponse{
			Code: 401,
			Status: "BAD_REQUEST",
			Data: "user already exist",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(WebResponse{
		Code: 201,
		Status: "OK",
		Data: struct{
			Email string `json:"email"`
		} {
			Email: requestBody.Email,
		},
	})
}

func Login(c *fiber.Ctx) error {
	var requestBody model.User
	var result model.User

	db := config.GetPostgresDatabase()
 
	c.BodyParser(&requestBody)

	ctx, cancel := config.NewPostgresContext()
	defer cancel()

	query := `SELECT id, email, password FROM users WHERE email = $1 LIMIT 1`
	rows, err := db.QueryContext(ctx, query, requestBody.Email)
	if err != nil || !rows.Next() {
		return c.Status(fiber.StatusBadRequest).JSON(WebResponse{
			Code: 401,
			Status: "BAD_REQUEST",
			Data: "user not found",
		})
	}

	defer rows.Close()
	rows.Scan(
		&result.Id,
		&result.Email,
		&result.Password,
	)
	
	checkPassword := helpers.ComparePassword([]byte(result.Password), []byte(requestBody.Password))
	if !checkPassword {
		return c.Status(fiber.StatusBadRequest).JSON(WebResponse{
			Code: 401,
			Status: "BAD_REQUEST",
			Data: "invalid password",
		})
	}

	access_token := helpers.SignToken(requestBody.Email)

	return c.Status(fiber.StatusOK).JSON(struct{
		Code int 
		Status string
		AccessToken string
		Data interface{}
	}{
		Code: 200,
		Status: "OK",
		AccessToken: access_token,
		Data: struct{
			Id string `json:"id"`
			Email string `json:"email"`
		} {
			Id: result.Id,
			Email: result.Email,
		},
	})
}

func Auth(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON("OK")
}
