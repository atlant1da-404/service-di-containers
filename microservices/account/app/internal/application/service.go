package application

import (
	"atlant1da-404/service-di-containers/account/app/internal/domain"
	"context"
)

type AccountService interface {
	CreateAccount(ctx context.Context, account *CreateAccountReqDto) (*CreateAccountResDto, error)
	UpdateAccount(ctx context.Context, account *UpdateAccountReqDto) error
	GetAccount(ctx context.Context, filter domain.GetAccountFilter) (*domain.Account, error)
	DeleteAccount(ctx context.Context, filter domain.DeleteAccountFilter) error
}
