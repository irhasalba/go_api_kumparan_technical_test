package handler

import (
	"go_api_kumparan/internal/dto"
	"go_api_kumparan/internal/service"
	"go_api_kumparan/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type ArticleHandler struct {
	s service.ArticleService
}

func NewArticleHandler(service service.ArticleService) *ArticleHandler {
	return &ArticleHandler{
		s: service,
	}
}

func (h *ArticleHandler) CreateArticle(ctx *fiber.Ctx) error {
	req := new(dto.RequestArticle)
	if err := ctx.BodyParser(req); err != nil {
		return utils.FailedResponse(ctx, 400, err.Error())
	}
	article, err := h.s.CreateArticle(ctx.Context(), req)
	if err != nil {
		return utils.FailedResponse(ctx, 400, err.Error())
	}
	return utils.SuccessResponse(ctx, 201, article)
}

func (h *ArticleHandler) GetAllArticle(ctx *fiber.Ctx) error {
	req := new(dto.GetAllArticleReq)
	if err := ctx.QueryParser(req); err != nil {
		return utils.FailedResponse(ctx, 400, err.Error())
	}
	article, err := h.s.GetAllArticle(ctx.Context(), req)
	if err != nil {
		return utils.FailedResponse(ctx, 400, err.Error())
	}
	return utils.SuccessResponse(ctx, 200, article)
}
