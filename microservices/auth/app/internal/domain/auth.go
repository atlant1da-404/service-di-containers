package domain

type Account struct {
	id       string
	email    string
	password string
}

func (a *Account) GetId() string {
	return a.id
}

func (a *Account) GetEmail() string {
	return a.email
}

func (a *Account) GetPassword() string {
	return a.password
}

type GetAccountFilter struct {
	Id string
}

type Token struct {
	accountId    string
	accessToken  string
	refreshToken string
}

func NewToken(accountId, accessToken, refreshToken string) *Token {
	return &Token{
		accountId:    accountId,
		accessToken:  accessToken,
		refreshToken: refreshToken,
	}
}

func (a *Token) GetAccess() string {
	return a.accessToken
}

func (a *Token) GetRefresh() string {
	return a.refreshToken
}

func (a *Token) GetAccountId() string {
	return a.accountId
}

type GetTokenFilter struct {
	AccountId    string
	RefreshToken string
}
