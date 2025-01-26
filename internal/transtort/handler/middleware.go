package handler

import (
	"net/http"
)

type HandlerFunc func (w http.ResponseWriter, r *http.Request) error

func (h Handler) AuthMW(handlerFunc http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h Handler) REAL(handler HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.loger.Info().Str("Method", r.Method).
			Str("Url", r.RequestURI).
			Msg("New request!")

		if err := handler(w,r);err != nil {
			h.loger.Error().Err(err)
		}
	}
}