package utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type ResponseStructure struct {
	Status    int         `json:"status"`
	Message   string      `json:"message"`
	Timestamp time.Time   `json:"timestamp"`
	Data      interface{} `json:"data"`
}

func SuccessResponse(ctx *fiber.Ctx, status int, data interface{}) error {
	res := ResponseStructure{
		Status:    status,
		Message:   "SUCCESS",
		Timestamp: time.Now(),
		Data:      data,
	}
	return ctx.Status(status).JSON(res)
}

func FailedResponse(ctx *fiber.Ctx, status int, message string) error {
	res := ResponseStructure{
		Status:    status,
		Message:   message,
		Timestamp: time.Now(),
	}
	return ctx.Status(status).JSON(res)
}
