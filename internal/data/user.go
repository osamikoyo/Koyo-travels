package data

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/osamikoyo/koyo-travels/internal/data/models"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

func (s Storage) UserRegister(user models.User) error {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil{
		return err
	}

	user.Password = string(password)

	return s.AddGorm(user)
}

func generateJWT(username, key string) (string, error) {
	claims := jwt.MapClaims{
		"username" : username,
		"exp" : time.Now().Add(180 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}

func (s Storage) UserLogin(user models.LoginUser) (string, error) {
	var u models.User
	key := os.Getenv("JWT_KEY")

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil{
		return "", err
	}

	if err := s.Gorm.Where(
			&models.User{
				Username: user.Username,
				Password: string(password),
			},
			).Find(&u).Error;err != nil{
		return "", err
	}

	token, err := generateJWT(user.Username, key)

	return token, err
}