package ahttp

import (
	"atlant1da-404/service-di-containers/account/app/internal/application"
	"encoding/json"
	"net/http"
)

type controller struct {
	service application.AccountService
}

func NewController(service application.AccountService) Controller {
	return &controller{service: service}
}

func (c *controller) CreateAccount(w http.ResponseWriter, r *http.Request) error {
	var dto application.CreateAccountReqDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return err
	}

	result, err := c.service.CreateAccount(r.Context(), &dto)
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

func (c *controller) UpdateAccount(w http.ResponseWriter, r *http.Request) error {
	var dto application.UpdateAccountReqDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return err
	}

	err := c.service.UpdateAccount(r.Context(), &dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	w.WriteHeader(http.StatusOK)
	return nil
}

func (c *controller) GetAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (c *controller) DeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
