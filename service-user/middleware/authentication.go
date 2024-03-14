package middleware

import (
	"database/sql"
	"fmt"
	"service-user/config"
	"service-user/helpers"
	"service-user/model"

	"github.com/gofiber/fiber/v2"
)

func Authentication(c *fiber.Ctx) error {
	access_token := c.Get("access_token")

	if len(access_token) == 0 {
		return c.Status(401).SendString("Invalid token: Access token missing")
	}

	checkToken, err := helpers.VerifyToken(access_token)

	if err != nil {
		return c.Status(401).SendString("Invalid token: Failed to verify token")
	}

	fmt.Println(checkToken, "CEKKKK" ,checkToken["email"])

	var user model.User

	db := config.GetPostgresDatabase()
	
	ctx, cancel := config.NewPostgresContext()
	defer cancel()

	query := `SELECT email FROM users WHERE email = $1 LIMIT = 1`
	rows := db.QueryRowContext(ctx, query, checkToken["email"])
	if rows.Scan(&user.Email) == sql.ErrNoRows {
		fmt.Println(err, "Error fetching user from database")
		return c.Status(401).SendString("Invalid token: User not found")
	}

	// Set user data in context for future use
	c.Locals("user", user)

	// Continue processing if user is found
	return c.Next()
}