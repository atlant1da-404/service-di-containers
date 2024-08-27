package ahttp

import "net/http"

type Controller interface {
	CreateAccount(w http.ResponseWriter, r *http.Request) error
	UpdateAccount(w http.ResponseWriter, r *http.Request) error
	GetAccount(w http.ResponseWriter, r *http.Request) error
	DeleteAccount(w http.ResponseWriter, r *http.Request) error
}
