package auth

type Authenticator interface {
	GenerateToken(options *GenerateTokenClaimsOptions) (string, error)
	RefreshToken(options *RefreshTokenClaimsOptions) (string, error)
	ParseToken(accessToken string) (*ParseTokenClaimsOutput, error)
}

type GenerateTokenClaimsOptions struct {
	AccountId string
}

type ParseTokenClaimsOutput struct {
	AccountId string
}

type RefreshTokenClaimsOptions struct {
	AccountId string
}
