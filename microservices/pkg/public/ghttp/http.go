package ghttp

import "net/http"

type HTTTPRouter interface {
	Get(pattern string, hx http.HandlerFunc)
	Post(pattern string, hx http.HandlerFunc)
	Put(pattern string, hx http.HandlerFunc)
	Delete(pattern string, hx http.HandlerFunc)
	ListenAndServe() error
	ErrorWrapper(func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc
}
