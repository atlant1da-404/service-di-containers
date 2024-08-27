package ahttp

import "net/http"

type Controller interface {
	Login(w http.ResponseWriter, r *http.Request) error
	Refresh(w http.ResponseWriter, r *http.Request) error
}
