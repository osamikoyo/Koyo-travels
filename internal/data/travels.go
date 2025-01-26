package data

import (
	"github.com/osamikoyo/koyo-travels/internal/data/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Storage) TravelGet(title string) (models.Travel, error) {
	var travels models.Travel

	cursor, err := s.Mongo.Collection.Find(s.Mongo.ctx, bson.M{"title" : title})
	if err != nil{
		return models.Travel{}, err
	}

	if err = cursor.All(s.Mongo.ctx, &travels);err != nil{
		return models.Travel{}, err
	}

	return travels, nil
}

func (s Storage) TravelUpdate(value interface{},key string, title string) error {
	filter := bson.M{
		"title" : title,
	}
	update := bson.M{
		key : value,
	}

	_, err := s.Mongo.Collection.UpdateOne(s.Mongo.ctx, filter, update)
	return err
}