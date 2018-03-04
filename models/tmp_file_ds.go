package models

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type tmpFileDS struct {
	filePath string
}

func newTmpFileDS() tmpFileDS {
	path := "/tmp/safetyCheck-review.json"

	feedbackDS := tmpFileDS{filePath: path}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		ioutil.WriteFile(path, []byte("[]"), 0644)
	}

	return feedbackDS
}

func (r tmpFileDS) FindReviewsForPlace(place Place) ([]Review, error) {
	return nil, errors.New("UnSupportOperationInDev")
}

func (r tmpFileDS) CreateReview(review Review) error {
	bytes, err := json.Marshal(review)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.filePath, bytes, 0644)
	if err != nil {
		return err
	}

	return nil
}
