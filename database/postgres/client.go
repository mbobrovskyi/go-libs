package postgres

import (
	"database/sql"
	"github.com/mbobrovskyi/go-libs/database/csql"
	"golang.org/x/net/context"

	_ "github.com/lib/pq"
)

func NewPostgresClient(ctx context.Context, dsn string) (*sql.DB, error) {
	return csql.NewSQLClientWithDriver(ctx, dsn, "postgres")
}
