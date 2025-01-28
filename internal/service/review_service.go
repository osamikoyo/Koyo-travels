package service

import (
	"encoding/json"
	"github.com/osamikoyo/koyo-travels/internal/data"
	"github.com/osamikoyo/koyo-travels/internal/data/models"
	"github.com/osamikoyo/koyo-travels/pkg/loger"
	"io/ioutil"
	"net/http"
)

type ReviewService struct {
	Data *data.Storage
	Loger loger.Logger
}

func (r ReviewService) Add(req *http.Request)  error {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil{
		return err
	}

	var review models.Review

	title := req.FormValue("title")

	if err := json.Unmarshal(body, &review);err != nil{
		return err
	}

	return r.Data.ReviewAdd(title, review)
}