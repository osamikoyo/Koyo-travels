package models

import "time"

type TokenResponse struct {
	Token string
}
type User struct {
	ID uint `gorm:"primaryKey"`
	Username string
	Password string
	CreatedAt time.Time
}

type LoginUser struct {
	Username string
	Password string
}