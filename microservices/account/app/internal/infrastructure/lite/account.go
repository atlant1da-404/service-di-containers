package lite

import (
	"atlant1da-404/service-di-containers/account/app/internal/application"
	"atlant1da-404/service-di-containers/account/app/internal/domain"
	"atlant1da-404/service-di-containers/pkg/public/database"
	"context"
)

type account struct {
	database database.Database
}

func NewAccountStorage(database database.Database) application.AccountStorage {
	return &account{
		database: database,
	}
}

func (a account) CreateAccount(ctx context.Context, entity *domain.Account) (string, error) {
	return "", nil
}

func (a account) UpdateAccount(ctx context.Context, account *domain.Account) error {
	return nil
}

func (a account) GetAccount(ctx context.Context, filter domain.GetAccountFilter) (*domain.Account, error) {
	return nil, nil
}

func (a account) DeleteAccount(ctx context.Context, filter domain.DeleteAccountFilter) error {
	return nil
}
