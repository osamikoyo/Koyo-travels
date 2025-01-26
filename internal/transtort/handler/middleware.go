package handler

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
)

type HandlerFunc func (w http.ResponseWriter, r *http.Request) error

func (h Handler) AuthMW(handlerFunc http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" {
			http.Error(w, "token not exist", http.StatusUnauthorized)
		}

		tokens, err := jwt.Parse(token[len("Bearer "):], func(token *jwt.Token) (interface{}, error) {
            return os.Getenv("JWT_KEY"), nil
        })

		if err != nil || !tokens.Valid{
			http.Error(w, "token not valid", http.StatusUnauthorized)
		}

		handlerFunc.ServeHTTP(w, r)
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