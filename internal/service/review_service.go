package service

import (
	"github.com/osamikoyo/koyo-travels/internal/data"
	"github.com/osamikoyo/koyo-travels/pkg/loger"
)

type ReviewService struct {
	Data *data.Storage
	Loger loger.Logger
}