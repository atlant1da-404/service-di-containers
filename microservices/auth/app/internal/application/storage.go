package application

import (
	"atlant1da-404/service-di-containers/auth/app/internal/domain"
	"context"
)

type AuthStorage interface {
	GetAccount(ctx context.Context, account domain.GetAccountFilter) (*domain.Account, error)
	GetToken(ctx context.Context, token domain.GetTokenFilter) (*domain.Token, error)
	CreateToken(ctx context.Context, token *domain.Token) error
}
