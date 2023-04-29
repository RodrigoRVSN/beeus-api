package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/rodrigoRVSN/beeus-api/src/infra/service"
)

func TestAuthMiddleware(t *testing.T) {
	app := fiber.New()

	app.Get("/protected", AuthMiddleware, func(c *fiber.Ctx) error {
		return c.SendString("Success")
	})

	req := httptest.NewRequest("GET", "/protected", nil)

	// Case 1: Missing authentication token
	resp, err := app.Test(req)
	if err != nil {
		t.Error("Failed to perform request")
	}
	if resp.StatusCode != http.StatusUnauthorized {
		t.Error("Expected status code 401")
	}

	// Case 2: Valid authentication token
	token, _ := service.CreateJwtToken(1)
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err = app.Test(req)
	if err != nil {
		t.Error("Failed to perform request")
	}
	if resp.StatusCode != http.StatusOK {
		t.Error("Expected status code 200")
	}
}
