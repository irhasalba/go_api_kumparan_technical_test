package main

import (
	"fmt"
	"go_api_kumparan/config"
	"go_api_kumparan/infra"
	"go_api_kumparan/internal/handler"
	"go_api_kumparan/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	app := fiber.New()

	config.Init()
	db, err := infra.PgConn()
	if err != nil {
		panic(err)
	}
	// init service
	articleService := service.NewArticleService(db, infra.PgxConn)
	// init handler
	articleHandler := handler.NewArticleHandler(*articleService)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Post("/article/new", articleHandler.CreateArticle)
	app.Get("/articles", articleHandler.GetAllArticle)

	port := fmt.Sprintf(":%d", config.Cfg.Port)
	app.Listen(port)
}
