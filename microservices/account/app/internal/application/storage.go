package application

import (
	"atlant1da-404/service-di-containers/account/app/internal/domain"
	"context"
)

type AccountStorage interface {
	CreateAccount(ctx context.Context, account *domain.Account) (string, error)
	UpdateAccount(ctx context.Context, account *domain.Account) error
	GetAccount(ctx context.Context, filter domain.GetAccountFilter) (*domain.Account, error)
	DeleteAccount(ctx context.Context, filter domain.DeleteAccountFilter) error
}
