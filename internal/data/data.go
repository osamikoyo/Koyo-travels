package data

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"time"
)

type Mongodb struct {
	ctx context.Context
	Collection *mongo.Collection
}

type Storage struct {
	Mongo *Mongodb
	Gorm *gorm.DB
}

func New() (*Storage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20 * time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
 	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil{
		return nil, err
	}

	collection := client.Database("travel").Collection("travels")

	mong := Mongodb{
		ctx: ctx,
		Collection: collection,
	}
	gr, err := gorm.Open(postgres.Open(os.Getenv("PSQL_URL")))
	if err != nil{
		return nil, err
	}

	return &Storage{
		Mongo: &mong,
		Gorm: gr,
	}, nil
}

func (s Storage) AddGorm(value interface{}) error {
	return  s.Gorm.Create(&value).Error
}


