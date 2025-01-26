package handler

import (
	"encoding/json"
	"errors"
	"github.com/osamikoyo/koyo-travels/internal/data/models"
	"net/http"
)

func (h Handler) registerHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		http.Error(w, "Method must to be post", http.StatusBadRequest)
	}

	return h.Service.User.Register(r)
}

func (h Handler) loginHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		http.Error(w, "Method must to be post", http.StatusBadRequest)
		return errors.New("method must to be post")
	}

	token, err := h.Service.User.Login(r)
	if err != nil{
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return err
	}

	resp := models.TokenResponse{
		Token: token,
	}

	body, err := json.Marshal(resp)
	if err != nil{
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return err
	}

	r.Response.Status = http.StatusText(200)
	_, err = w.Write(body)
	return err
}