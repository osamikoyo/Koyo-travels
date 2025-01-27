package data

import (
	"github.com/osamikoyo/koyo-travels/internal/data/models"
	"go.mongodb.org/mongo-driver/bson"
	"sort"
	"strings"
)

type Match struct {
	Travel models.Travel
	Count int
}

type ByCount []Match

func (a ByCount) Len() int           { return len(a) }
func (a ByCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCount) Less(i, j int) bool { return a[i].Count > a[j].Count }

func countMatches(mainStr, sideStr string) int {
    mainWords := strings.Fields(mainStr)
    sideWords := strings.Fields(sideStr)

    matchCount := 0
    sideWordsMap := make(map[string]struct{})

    for _, word := range sideWords {
        sideWordsMap[word] = struct{}{}
    }

   	for _, word := range mainWords {
        if _, exists := sideWordsMap[word]; exists {
            matchCount++
        }
    }

    return matchCount
}

func (s Storage) Search(title string) ([]Match, error) {
	var travels []models.Travel
	var sortTravels []Match

	cursor, err := s.Mongo.Collection.Find(s.Mongo.ctx, bson.M{})
	if err != nil{
		return nil, err
	}
	defer cursor.Close(s.Mongo.ctx)

	for cursor.Next(s.Mongo.ctx){
		var travel models.Travel
		if err := cursor.Decode(&travel);err != nil{
			return nil, err
		}
		travels = append(travels, travel)
	}

	if err := cursor.Err();err != nil{
		return nil, err
	}

	for _, t := range travels{
		count := countMatches(title, t.Title)
		sortTravels = append(sortTravels, Match{Travel: t, Count: count})
	}

	sort.Sort(ByCount(sortTravels))
	return sortTravels, nil
}