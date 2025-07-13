package service

import (
	"context"
	"database/sql"
	"go_api_kumparan/internal/dto"
	"go_api_kumparan/internal/query"
	"go_api_kumparan/internal/utils"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type ArticleService struct {
	q       *query.Queries
	pgxConn *pgx.Conn
}

func NewArticleService(query *query.Queries, pgx *pgx.Conn) *ArticleService {
	return &ArticleService{
		q:       query,
		pgxConn: pgx,
	}
}

func (a ArticleService) CreateArticle(ctx context.Context, req *dto.RequestArticle) (interface{}, error) {
	// slugify title
	slugifyTitle := utils.Slugify(req.Title)
	tx, err := a.pgxConn.BeginTx(ctx, pgx.TxOptions{})
	qtx := a.q.WithTx(tx)
	if err != nil {
		return nil, err
	}
	result, err := qtx.CreateArticle(ctx, query.CreateArticleParams{
		ID:       slugifyTitle,
		AuthorID: pgtype.Text{String: req.AuthorId, Valid: true},
		Title:    pgtype.Text{String: req.Title, Valid: true},
		Body:     pgtype.Text{String: req.Body, Valid: true},
	})
	if err != nil {
		return nil, err
	}
	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}
	return result, nil

}

func (a ArticleService) GetAllArticle(ctx context.Context, req *dto.GetAllArticleReq) (interface{}, error) {
	result, err := a.q.ListArticlesFiltered(ctx, query.ListArticlesFilteredParams{
		Column1: sql.NullString{String: req.Query, Valid: true},
		Column2: sql.NullString{String: req.Author, Valid: true},
		Limit:   5,
		Offset:  int32(utils.PageToOffset(req.Page)),
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}
