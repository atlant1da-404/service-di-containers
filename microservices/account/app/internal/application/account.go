package application

import (
	"atlant1da-404/service-di-containers/account/app/internal/domain"
	"atlant1da-404/service-di-containers/pkg/public/hash"
	"context"
	"github.com/rotisserie/eris"
)

type account struct {
	storage AccountStorage
	hash    hash.Hash
}

func NewAccountService(storage AccountStorage, hash hash.Hash) AccountService {
	return &account{storage: storage, hash: hash}
}

func (a account) CreateAccount(ctx context.Context, dto *CreateAccountReqDto) (*CreateAccountResDto, error) {
	user, err := domain.NewAccount(dto.Id, dto.Name, dto.Email, dto.Password, dto.Role, dto.Age)
	if err != nil {
		return nil, eris.Wrapf(err, "failed to make new account")
	}

	hashedPassword, err := a.hash.GenerateHash(user.GetPassword())
	if err != nil {
		return nil, eris.Wrapf(err, "failed to generate hash")
	}

	user, err = user.SetPassword(hashedPassword)
	if err != nil {
		return nil, eris.Wrapf(err, "failed to set user password")
	}

	id, err := a.storage.CreateAccount(ctx, user)
	if err != nil {
		return nil, eris.Wrapf(err, "failed to create account: account")
	}
	return &CreateAccountResDto{Id: id}, nil
}

func (a account) UpdateAccount(ctx context.Context, dto *UpdateAccountReqDto) error {
	user, err := a.storage.GetAccount(ctx, domain.GetAccountFilter{Id: dto.Id})
	if err != nil {
		return eris.Wrapf(err, "failed to find account")
	}

	user, err = user.SetName(dto.Name)
	if err != nil {
		return eris.Wrapf(err, "failed to update name")
	}
	user, err = user.SetEmail(dto.Email)
	if err != nil {
		return eris.Wrapf(err, "failed to update email")
	}
	user, err = user.SetAge(dto.Age)
	if err != nil {
		return eris.Wrapf(err, "failed to set age")
	}

	err = a.storage.UpdateAccount(ctx, user)
	if err != nil {
		return eris.Wrapf(err, "failed to update account")
	}
	return nil
}

func (a account) GetAccount(ctx context.Context, filter domain.GetAccountFilter) (*domain.Account, error) {
	user, err := a.storage.GetAccount(ctx, filter)
	if err != nil {
		return nil, eris.Wrapf(err, "failed to find account")
	}
	return user, nil
}

func (a account) DeleteAccount(ctx context.Context, filter domain.DeleteAccountFilter) error {
	err := a.storage.DeleteAccount(ctx, filter)
	if err != nil {
		return eris.Wrapf(err, "failed to delete account")
	}
	return nil
}
