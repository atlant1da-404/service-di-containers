package database

import "context"

type Database interface {
	Insert(ctx context.Context, pattern string, body ...interface{}) error
	Get(ctx context.Context, pattern string) (interface{}, error)
	Update(ctx context.Context, pattern string, body interface{}) error
	Delete(ctx context.Context, pattern string, body interface{}) error
}
