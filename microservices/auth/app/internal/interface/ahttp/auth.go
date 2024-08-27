package ahttp

import (
	"atlant1da-404/service-di-containers/auth/app/internal/application"
	"encoding/json"
	"net/http"
)

type controller struct {
	service application.AuthService
}

func NewController(service application.AuthService) Controller {
	return &controller{service: service}
}

func (c controller) Login(w http.ResponseWriter, r *http.Request) error {
	var dto application.LoginReqDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return err
	}

	result, err := c.service.Login(r.Context(), &dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	w.WriteHeader(http.StatusCreated)
	return nil
}

func (c controller) Refresh(w http.ResponseWriter, r *http.Request) error {
	var dto application.RefreshReqDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return err
	}

	result, err := c.service.Refresh(r.Context(), &dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	w.WriteHeader(http.StatusCreated)
	return nil
}
