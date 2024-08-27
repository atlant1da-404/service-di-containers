package database

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type sqlSQL struct {
	conn *pgx.Conn
}

func NewSQLDatabase(address string) (Database, error) {
	return &sqlSQL{}, nil
}

func (p sqlSQL) Insert(ctx context.Context, pattern string, body ...interface{}) error {
	return nil
}

func (p sqlSQL) Get(ctx context.Context, pattern string) (interface{}, error) {
	return nil, nil
}

func (p sqlSQL) Update(ctx context.Context, pattern string, body interface{}) error {
	return nil
}

func (p sqlSQL) Delete(ctx context.Context, pattern string, body interface{}) error {
	return nil
}
