package ghttp

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/rotisserie/eris"
	"net/http"
)

type goChi struct {
	address string
	mux     *chi.Mux
}

func (g goChi) Get(pattern string, hx http.HandlerFunc) {
	g.mux.Get(pattern, hx)
}

func (g goChi) Post(pattern string, hx http.HandlerFunc) {
	g.mux.Post(pattern, hx)
}

func (g goChi) Put(pattern string, hx http.HandlerFunc) {
	g.mux.Put(pattern, hx)
}

func (g goChi) Delete(pattern string, hx http.HandlerFunc) {
	g.mux.Delete(pattern, hx)
}

func NewHTTPRouter(address string) HTTTPRouter {
	return &goChi{
		address: address,
		mux:     chi.NewRouter(),
	}
}

func (g goChi) ErrorWrapper(hx func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := hx(w, r)
		if err != nil {
			formatted := eris.ToCustomJSON(err, eris.JSONFormat{Options: eris.FormatOptions{InvertOutput: true, WithExternal: true, WithTrace: true}})
			errLog, _ := json.MarshalIndent(formatted, "", "  ")
			fmt.Println(string(errLog))
		}
		return
	}
}

func (g goChi) ListenAndServe() error {
	err := http.ListenAndServe(g.address, g.mux)
	if err != nil {
		return fmt.Errorf("failed to start ghttp server: %w", err)
	}

	return nil
}
