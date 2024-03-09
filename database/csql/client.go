package csql

import (
	"context"
	"database/sql"
)

func NewSQLClientWithDriver(ctx context.Context, dataSourceName, driverName string) (*sql.DB, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}
	return db, nil
}
