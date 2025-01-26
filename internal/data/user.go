package data

import (
	"github.com/osamikoyo/koyo-travels/internal/data/models"
	"golang.org/x/crypto/bcrypt"
)

func (s Storage) UserRegister(user models.User) error {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil{
		return err
	}

	user.Password = string(password)

	return s.AddGorm(user)
}

func (s Storage) UserLogin(user models.LoginUser) error {
	var u models.User

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil{
		return err
	}

	if err := s.Gorm.Where(
			&models.User{
				Username: user.Username,
				Password: string(password),
			},
			).Find(&u).Error;err != nil{
		return err
	}

	return nil
}