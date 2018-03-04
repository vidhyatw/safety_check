package models

import (
	"strings"

	"github.com/hackerearth/safetycheck/config"
)

type PersistDS interface {
	CreateReview(review Review) error
	FindReviewsForPlace(place Place) ([]Review, error)
}

func GetDatasource() PersistDS {
	if strings.ToLower(config.GetEnvironment()) == "DEVELOPMENT" {
		return newTmpFileDS()
	} else {
		return newMongoDBDS()
	}
}
