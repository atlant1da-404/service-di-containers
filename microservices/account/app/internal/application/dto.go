package application

type CreateAccountReqDto struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Age      int    `json:"age"`
}

type CreateAccountResDto struct {
	Id string `json:"id"`
}

type UpdateAccountReqDto struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}
