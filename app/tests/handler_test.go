package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	"news-feed-system-backend/app/handlers"
	"news-feed-system-backend/database"
)

func setupApp() *fiber.App {
	database.ConnectDB()
	app := fiber.New()

	app.Post("/api/register", handlers.Register)
	app.Post("/api/login", handlers.Login)
	app.Post("/api/posts", handlers.CreatePost)
	app.Post("/api/follow/:userid", handlers.Follow)
	app.Delete("/api/follow/:userid", handlers.Unfollow)
	app.Get("/api/feed", handlers.GetFeed)

	return app
}

func TestRegisterAndLogin(t *testing.T) {
	app := setupApp()

	body := []byte(`{"username": "budi", "password": "123456"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/register", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	reqDup := httptest.NewRequest(http.MethodPost, "/api/register", bytes.NewReader(body))
	reqDup.Header.Set("Content-Type", "application/json")
	respDup, _ := app.Test(reqDup)

	assert.Equal(t, http.StatusConflict, respDup.StatusCode)

	loginBody := []byte(`{"username": "budi", "password": "123456"}`)
	reqLogin := httptest.NewRequest(http.MethodPost, "/api/login", bytes.NewReader(loginBody))
	reqLogin.Header.Set("Content-Type", "application/json")
	respLogin, _ := app.Test(reqLogin)

	assert.Equal(t, http.StatusOK, respLogin.StatusCode)
}

func TestCreatePost(t *testing.T) {
	app := setupApp()

	token := "Bearer dummy_token"

	postBody := []byte(`{"content": "Hello world!"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/posts", bytes.NewReader(postBody))
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	longContent := make([]byte, 201)
	for i := 0; i < 201; i++ {
		longContent[i] = 'a'
	}
	body := map[string]string{"content": string(longContent)}
	jsonBody, _ := json.Marshal(body)

	reqLong := httptest.NewRequest(http.MethodPost, "/api/posts", bytes.NewReader(jsonBody))
	reqLong.Header.Set("Authorization", token)
	reqLong.Header.Set("Content-Type", "application/json")
	respLong, _ := app.Test(reqLong)

	assert.Equal(t, http.StatusUnprocessableEntity, respLong.StatusCode)
}

func TestFollowUnfollow(t *testing.T) {
	app := setupApp()
	token := "Bearer dummy_token"

	req := httptest.NewRequest(http.MethodPost, "/api/follow/2", nil)
	req.Header.Set("Authorization", token)
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	req404 := httptest.NewRequest(http.MethodPost, "/api/follow/9999", nil)
	req404.Header.Set("Authorization", token)
	resp404, _ := app.Test(req404)
	assert.Equal(t, http.StatusNotFound, resp404.StatusCode)

	reqDel := httptest.NewRequest(http.MethodDelete, "/api/follow/2", nil)
	reqDel.Header.Set("Authorization", token)
	respDel, _ := app.Test(reqDel)
	assert.Equal(t, http.StatusOK, respDel.StatusCode)
}

func TestFeed(t *testing.T) {
	app := setupApp()
	token := "Bearer dummy_token"

	req := httptest.NewRequest(http.MethodGet, "/api/feed?page=1&limit=10", nil)
	req.Header.Set("Authorization", token)
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
