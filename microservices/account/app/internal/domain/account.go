package domain

import (
	"errors"
)

type Account struct {
	id       string
	name     string
	email    string
	password string
	role     string
	age      int
}

// NewAccount provides ability to create new user domain application.
func NewAccount(id, name, email, password, role string, age int) (*Account, error) {
	if id == "" {
		return nil, errors.New("failed to validate 'id', id is empty")
	}
	if name == "" {
		return nil, errors.New("failed to validate 'name', name is empty")
	}
	if email == "" {
		return nil, errors.New("failed to validate 'email', email is empty")
	}
	if password == "" {
		return nil, errors.New("failed to validate 'password', password is empty")
	}
	if role == "" {
		return nil, errors.New("failed to validate 'role', role is empty")
	}
	if age <= 18 {
		return nil, errors.New("failed to validate 'age', age is not correct")
	}

	account := &Account{
		id:       id,
		name:     name,
		email:    email,
		password: password,
		role:     role,
		age:      age,
	}

	return account, nil
}

func (a *Account) GetId() string {
	return a.id
}
func (a *Account) GetName() string {
	return a.name
}
func (a *Account) GetEmail() string {
	return a.email
}
func (a *Account) GetPassword() string {
	return a.password
}
func (a *Account) GetRole() string {
	return a.role
}
func (a *Account) GetAge() int {
	return a.age
}

func (a *Account) SetId(id string) (*Account, error) {
	if id == "" {
		return nil, errors.New("failed to validate 'id', id is empty")
	}
	a.id = id
	return a, nil
}
func (a *Account) SetName(name string) (*Account, error) {
	if name == "" {
		return nil, errors.New("failed to validate 'name', name is empty")
	}
	a.name = name
	return a, nil
}
func (a *Account) SetEmail(email string) (*Account, error) {
	if email == "" {
		return nil, errors.New("failed to validate 'email', email is empty")
	}
	a.email = email
	return a, nil
}
func (a *Account) SetPassword(password string) (*Account, error) {
	if password == "" {
		return nil, errors.New("failed to validate 'password', password is empty")
	}
	a.password = password
	return a, nil
}
func (a *Account) SetRole(role string) (*Account, error) {
	if role == "" {
		return nil, errors.New("failed to validate 'role', role is empty")
	}
	a.role = role
	return a, nil
}
func (a *Account) SetAge(age int) (*Account, error) {
	if age <= 0 {
		return nil, errors.New("failed to validate 'age', age is correct")
	}
	a.age = age
	return a, nil
}

type GetAccountFilter struct {
	Id string
}

type DeleteAccountFilter struct {
	Id string
}
