package service

import (
	"encoding/json"
	"github.com/osamikoyo/koyo-travels/internal/data"
	"github.com/osamikoyo/koyo-travels/internal/data/models"
	"github.com/osamikoyo/koyo-travels/pkg/loger"
	"io/ioutil"
	"net/http"
)

type TravelService struct {
	Data *data.Storage
	Loger loger.Logger
}

func (t *TravelService) AddTravel()

func (t *TravelService) UpdateHotel(r *http.Request) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		return err
	}

	title := r.FormValue("title")

	var hotel models.Hotel
	if err := json.Unmarshal(body, &hotel);err != nil{
		return err
	}

	return t.Data.TravelUpdate(hotel, "hotel", title)
}

func (t *TravelService) AddExcurs(r *http.Request) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		return err
	}

	title := r.FormValue("title")

	var excurs models.Excurs
	if err := json.Unmarshal(body, &excurs);err != nil{
		return err
	}
	
	return t.Data.TravelExcursAdd(title, excurs)
}