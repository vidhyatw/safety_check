package models

import (
	"strings"

	"github.com/hackerearth/safetycheck/config"
)

type DBApi interface {
	CreateReview(review Review) error
	FindReviewsForPlace(place Place) (string, []Review, error)
}

func GetDatasource() DBApi {
	if strings.ToLower(config.GetEnvironment()) == "DEVELOPMENT" {
		return newTmpFileDS()
	} else {
		return newMongoAtlasDBDS()
	}
}
