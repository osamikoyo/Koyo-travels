package data

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type Mongodb struct {
	ctx context.Context
	Collection *mongo.Collection
}

type Storage struct {
	Mongo *Mongodb
	Gorm *gorm.DB
}

func (s Storage) AddGorm(value interface{}) error {
	return  s.Gorm.Create(&value).Error
}


