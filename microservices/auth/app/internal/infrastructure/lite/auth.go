package lite

import (
	"atlant1da-404/service-di-containers/auth/app/internal/application"
	"atlant1da-404/service-di-containers/auth/app/internal/domain"
	"atlant1da-404/service-di-containers/pkg/public/database"
	"context"
)

type auth struct {
	database database.Database
}

func NewAuthStorage(database database.Database) application.AuthStorage {
	return &auth{
		database: database,
	}
}

func (a auth) GetAccount(ctx context.Context, account domain.GetAccountFilter) (*domain.Account, error) {
	return nil, nil
}

func (a auth) GetToken(ctx context.Context, token domain.GetTokenFilter) (*domain.Token, error) {
	return nil, nil
}

func (a auth) CreateToken(ctx context.Context, token *domain.Token) error {
	return nil
}
