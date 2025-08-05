package test_test

import (
	"bytes"
	"crud/test"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestCreateContact_Success(t *testing.T) {
	app := fiber.New()
	app.Post("/contacts", test.CreateContactMock) // ← pakai mock handler

	payload := map[string]interface{}{
		"name":   "John Doe",
		"email":  "john@example.com",
		"phone":  "1234567890",
		"gender": "male",
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest("POST", "/contacts", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

func TestCreateContact_InvalidEmail(t *testing.T) {
	app := fiber.New()
	app.Post("/api/contacts", test.CreateContactMock) // Tambahkan ini

	payload := map[string]string{
		"name":   "John Doe",
		"email":  "invalid-email", // Invalid email
		"phone":  "081234567890",
		"gender": "male",
	}

	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/api/contacts", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}
