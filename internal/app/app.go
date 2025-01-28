package app

import (
	"github.com/osamikoyo/koyo-travels/internal/transtort/handler"
	"github.com/osamikoyo/koyo-travels/pkg/loger"
)

type App struct {
	loger loger.Logger
	handler handler.Handler
}

func New() App{
	
}