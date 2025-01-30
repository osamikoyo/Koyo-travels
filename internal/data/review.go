package data

import (
	"github.com/osamikoyo/koyo-travels/internal/data/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Storage) ReviewAdd(title string, review models.Review) error {
	filter := bson.M{
		"title" : title,
	}

	err := s.TravelUpdateMark(float32(review.Count), title)
	if err != nil{
		return err
	}

	update := bson.M{
		"$push": bson.M{
			"reviews" : review,
		},
	}
	_, err = s.Mongo.Collection.UpdateOne(s.Mongo.ctx, filter, update)
	return err
}

func (s *Storage) CalculateMarks(title string) (float32, error) {
	travels, err := s.TravelGet(title)
	if err != nil{
		return 0, err
	}

	var counter uint64

	for _, t := range travels.Reviews{
		counter =  counter + uint64(t.Count)
	}
	return float32(counter/uint64(len(travels.Reviews))), nil
}