package models

import "time"

type User struct {
	ID uint `gorm:"primaryKey"`
	Username string
	Password string
	CreatedAt time.Time
}