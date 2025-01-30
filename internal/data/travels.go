package data

import (
	"github.com/osamikoyo/koyo-travels/internal/data/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Storage) TravelAdd(travel models.Travel) error {
	_, err := s.Mongo.Collection.InsertOne(s.Mongo.ctx, travel)
	return err
}

func (s *Storage) TravelUpdateMark(newMark float32, title string) (float32, error) {
	travel, err := s.TravelGet(title)
	if err != nil{
		return 0.0, err
	}

	counter := 0
	for _, r := range travel.Reviews {
		counter =+ int(r.Count)
	}

	counter =+ int(newMark)

	result := float32(counter/(len(travel.Reviews)+1))

	err = s.TravelUpdate(result, "mark", title)
	return result, err
}

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

func (s Storage) TravelExcursAdd(title string, excurs models.Excurs) error {
	filter := bson.M{
		"title" : title,
	}

	update := bson.M{
		"$push": bson.M{
			"excurs" : excurs,
		},
	}

	_, err := s.Mongo.Collection.UpdateOne(s.Mongo.ctx, filter, update)
	return err
}

func (s Storage) TravelDeleteExcurs(title string, excurs models.Excurs) error {
	filter := bson.M{
		"title" : title,
	}

	update := bson.M{
		"$pull" : bson.M{
			"excurs" : bson.M{
				"title" : excurs.Title,
			},
		},
	}

	_, err := s.Mongo.Collection.UpdateOne(s.Mongo.ctx, filter, update)
	return err
}

func (s Storage) UpdateExcurs(travel_title string, excurs models.Excurs) error {
	filter := bson.M{
		"title" : travel_title,
		"excurs.title" : excurs.Title,
	}

	update := bson.M{
		"$set" : bson.M{
			"excurs.$" : excurs,
		},
	}

	_, err := s.Mongo.Collection.UpdateOne(s.Mongo.ctx, filter, update)
	return err
}

