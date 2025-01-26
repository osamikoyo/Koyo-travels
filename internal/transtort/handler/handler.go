package handler

import (
	"github.com/osamikoyo/koyo-travels/internal/service"
	"github.com/osamikoyo/koyo-travels/pkg/loger"
	"net/http"
)

type Handler struct {
	loger loger.Logger
	Service *service.Service
}

func (h Handler) GetRoutes(mux *http.ServeMux){
	mux.Handle("/register", h.REAL(h.registerHandler))
}

