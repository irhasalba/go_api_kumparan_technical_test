package main

import (
	"fmt"
	"go_api_kumparan/config"
	"go_api_kumparan/infra"
	"go_api_kumparan/internal/handler"
	"go_api_kumparan/internal/service"
	"go_api_kumparan/internal/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
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

	app.Use(limiter.New(limiter.Config{
		Max:        100,               // max 100 requests
		Expiration: 300 * time.Second, // per 5 minute
		LimitReached: func(ctx *fiber.Ctx) error {
			return utils.FailedResponse(ctx, 429, "Too many requests. Please try again later.")
		},
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Post("/article/new", articleHandler.CreateArticle)
	app.Get("/articles", articleHandler.GetAllArticle)

	port := fmt.Sprintf(":%d", config.Cfg.Port)
	app.Listen(port)
}
