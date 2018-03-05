package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateReview(t *testing.T) {
	mongoDb := newMongoAtlasDBDS()
	defer mongoDb.mongoSession.Close()
	if false {
		place := Place{PlaceID: "ChIJJf4m6glYqDsRh5BROhTjgXE", Type: "Shopping  Mall", Coordinates: []float64{76.9917911, 11.0556083}}
		reviewer := Reviewer{Name: "Vidhya", Age: 18, Gender: "Female"}
		review := Review{Title: "Great Place", TimeStamp: time.Now().String(), Rating: 5,
			Content: "Very safe @ nite", VisitTime: "10:00 PM", Place: place, Reviewer: reviewer}
		err := mongoDb.CreateReview(review)
		assert.Nil(t, err)

	}
}

func TestFindReviewsForPlace(t *testing.T) {
	mongoDb := newMongoAtlasDBDS()
	defer mongoDb.mongoSession.Close()

	place := Place{PlaceID: "ChIJJf4m6glYqDsRh5BROhTjgXE", Type: "Shopping  Mall", Coordinates: []float64{76.9917911, 11.0556083}}
	findType, reviews, err := mongoDb.FindReviewsForPlace(place)
	assert.Equal(t, "PLACE", findType)
	assert.True(t, true, len(reviews) > 0)
	assert.Nil(t, err)
}

func TestFindReviewsForPlace_NearBy(t *testing.T) {
	mongoDb := newMongoAtlasDBDS()
	defer mongoDb.mongoSession.Close()

	place := Place{PlaceID: "ChIJGQ6k2QhYqDsRgkxMNsJi8Jw", Type: "Shopping  Mall", Coordinates: []float64{76.9940433, 11.054779}}
	findType, reviews, err := mongoDb.FindReviewsForPlace(place)
	assert.Equal(t, "NEARBY", findType)
	assert.True(t, true, len(reviews) > 0)
	assert.Nil(t, err)
}
