package application

import (
	"atlant1da-404/service-di-containers/auth/app/internal/domain"
	"atlant1da-404/service-di-containers/auth/app/internal/pkg/auth"
	"atlant1da-404/service-di-containers/pkg/public/hash"
	"context"
	"github.com/rotisserie/eris"
)

type service struct {
	storage AuthStorage
	hash    hash.Hash
	auth    auth.Authenticator
}

func NewAuthService(storage AuthStorage, hash hash.Hash, auth auth.Authenticator) AuthService {
	return &service{storage: storage, hash: hash, auth: auth}
}

func (s service) Login(ctx context.Context, dto *LoginReqDto) (*LoginResDto, error) {
	account, err := s.storage.GetAccount(ctx, domain.GetAccountFilter{Id: dto.AccountId})
	if err != nil {
		return nil, eris.Wrapf(err, "failed to get account")
	}

	err = s.hash.CompareHash([]byte(account.GetPassword()), []byte(dto.Password))
	if err != nil {
		return nil, eris.Wrapf(err, "something went wrong")
	}

	accessToken, err := s.auth.GenerateToken(&auth.GenerateTokenClaimsOptions{AccountId: account.GetId()})
	if err != nil {
		return nil, eris.Wrapf(err, "failed to generate token")
	}

	refreshToken, err := s.auth.RefreshToken(&auth.RefreshTokenClaimsOptions{AccountId: account.GetId()})
	if err != nil {
		return nil, eris.Wrapf(err, "failed to generate token")
	}

	err = s.storage.CreateToken(ctx, domain.NewToken(account.GetId(), accessToken, refreshToken))
	if err != nil {
		return nil, eris.Wrapf(err, "failed to save tokens")
	}

	return &LoginResDto{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func (s service) Refresh(ctx context.Context, dto *RefreshReqDto) (*RefreshResDto, error) {
	token, err := s.storage.GetToken(ctx, domain.GetTokenFilter{RefreshToken: dto.RefreshToken, AccountId: dto.AccountId})
	if err != nil {
		return nil, eris.Wrapf(err, "failed to get token")
	}

	accessToken, err := s.auth.GenerateToken(&auth.GenerateTokenClaimsOptions{AccountId: token.GetAccountId()})
	if err != nil {
		return nil, eris.Wrapf(err, "failed to generate token")
	}

	err = s.storage.CreateToken(ctx, domain.NewToken(token.GetAccountId(), accessToken, token.GetRefresh()))
	if err != nil {
		return nil, eris.Wrapf(err, "failed to save tokens")
	}

	return &RefreshResDto{AccessToken: accessToken, RefreshToken: token.GetRefresh()}, nil
}
