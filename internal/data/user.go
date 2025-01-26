package data

import "github.com/osamikoyo/koyo-travels/internal/data/models"

func (s Storage) UserRegister(user models.User) error {
	return s.AddGorm(user)
}

func (s Storage) UserLogin(user models.LoginUser) error {
	var u models.User

	password, err := 

	if err := s.Gorm.Where(
			&models.User{
				Username: user.Username,

			}
		)
}