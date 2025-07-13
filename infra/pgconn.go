package infra

import (
	"context"
	"go_api_kumparan/config"
	"go_api_kumparan/internal/query"
	"log"

	"github.com/jackc/pgx/v5"
)

var PgxConn *pgx.Conn

func PgConn() (*query.Queries, error) {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, config.Cfg.DatabaseUrl)
	if err != nil {
		log.Fatal("Error connected to database")
		return nil, err
	}
	PgxConn = conn
	q := query.New(conn)
	return q, nil
}
