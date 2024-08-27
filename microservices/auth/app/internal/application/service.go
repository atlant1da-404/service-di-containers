package application

import "context"

type AuthService interface {
	Login(ctx context.Context, dto *LoginReqDto) (*LoginResDto, error)
	Refresh(ctx context.Context, dto *RefreshReqDto) (*RefreshResDto, error)
}
