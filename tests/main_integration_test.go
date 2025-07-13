package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"go_api_kumparan/config"
	"go_api_kumparan/infra"
	"go_api_kumparan/internal/handler"
	"go_api_kumparan/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func setupApp() *fiber.App {
	_ = godotenv.Load()
	config.Init()
	db, err := infra.PgConn()
	if err != nil {
		panic(err)
	}
	articleService := service.NewArticleService(db, infra.PgxConn)
	articleHandler := handler.NewArticleHandler(*articleService)

	app := fiber.New()
	app.Post("/article/new", articleHandler.CreateArticle)
	app.Get("/articles", articleHandler.GetAllArticle)
	return app
}

func TestCreateArticleIntegration(t *testing.T) {
	app := setupApp()

	payload := map[string]interface{}{
		"title":     "Judul Test",
		"body":      "Isi Test",
		"author_id": "1",
	}
	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/article/new", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	if resp.StatusCode != fiber.StatusOK && resp.StatusCode != fiber.StatusCreated {
		t.Errorf("expected status 200/201, got %d", resp.StatusCode)
	}
	respBody, _ := io.ReadAll(resp.Body)
	if !bytes.Contains(respBody, []byte("Judul Test")) {
		t.Errorf("response body does not contain article title: %s", respBody)
	}
}

func TestGetArticlesIntegration(t *testing.T) {
	app := setupApp()
	req := httptest.NewRequest(http.MethodGet, "/articles", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
	respBody, _ := io.ReadAll(resp.Body)
	if !bytes.Contains(respBody, []byte("title")) && !bytes.Contains(respBody, []byte("Judul")) {
		t.Errorf("response body does not contain articles: %s", respBody)
	}
}
