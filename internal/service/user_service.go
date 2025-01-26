package service

import (
	"encoding/json"
	"github.com/osamikoyo/koyo-travels/internal/data"
	"github.com/osamikoyo/koyo-travels/internal/data/models"
	"github.com/osamikoyo/koyo-travels/pkg/loger"
	"io/ioutil"
	"net/http"
)

type UserService struct {
	Data *data.Storage
	Loger loger.Logger
}

func (u *UserService) Register(r *http.Request) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		return err
	}

	u.Loger.Info().Msg("Unmarshaling...")

	var user models.User
	if err = json.Unmarshal(body, &user);err != nil{
		return err
	}

	u.Loger.Info().Msg("Success!")

	return u.Data.UserRegister(user)
}

func (u *UserService) Login(r *http.Request) (string, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		return "", err
	}

	var user models.LoginUser
	if err = json.Unmarshal(body, &user);err != nil{
		return "", err
	}

	return u.Data.UserLogin(user)
}