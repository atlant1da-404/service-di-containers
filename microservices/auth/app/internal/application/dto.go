package application

type LoginReqDto struct {
	AccountId string `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type LoginResDto struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshReqDto struct {
	AccountId    string `json:"id"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshResDto struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
