// handlers.review.go

package handlers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hackerearth/safety_check/config"
	"github.com/hackerearth/safety_check/models"
	"github.com/hackerearth/safety_check/resources"
)

type SafetyScore struct {
	Time  int     `json:"time,omitempty" binding:"omitempty,len=0"`
	Score float32 `json:"score,omitempty" binding:"omitempty,len=0"`
}

func ShowIndexPage(c *gin.Context) {
	// reviews := getAllreviews()

	// // Call the render function with the name of the template to render
	// render(c, gin.H{
	// 	"title":   "Home Page",
	// 	"payload": reviews}, "index.html")
	fmt.Printf("fb app id: ", config.GetFacebookApp())
	// render(c, gin.H{
	// 	"fb_appId": config.GetFacebookApp(), "CurrentPage": "home"}, "landing-page.html")
	render(c, gin.H{
		"fb_appId": config.GetFacebookApp(), "CurrentPage": "home"}, "index.html")
}

func RenderAssets(c *gin.Context) {
	file, _ := resources.Asset("static" + c.Param("path"))
	contentType := ""
	if path.Ext(c.Param("path")) == ".css" {
		contentType = "text/css"
	} else if path.Ext(c.Param("path")) == ".js" {
		contentType = "text/javascript"
	} else if path.Ext(c.Param("path")) == ".png" {
		contentType = "image/png"
	}
	c.Data(http.StatusOK, contentType, file)
	return
}

func ShowReviewCreationPage(c *gin.Context) {
	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title": "Create New review", "CurrentPage": "submit-review"}, "create-review.html")
}

func createPlaceFromParams(c *gin.Context) (models.Place, error) {
	var finalError error
	placeid := c.Param("placeid")

	long, err := strconv.ParseFloat(c.Param("long"), 64)
	if err != nil {
		finalError = errors.New(err.Error())

	}

	lat, err := strconv.ParseFloat(c.Param("lat"), 64)
	if err != nil {
		finalError = errors.New(finalError.Error() + err.Error())
		return models.Place{}, finalError
	}
	return models.Place{PlaceID: placeid, Coordinates: []float64{long, lat}}, nil
}

func GetReview(c *gin.Context) {
	// Check if the place ID is valid or long and lat is present

	if place, err := createPlaceFromParams(c); err == nil {
		// Check if the review exists
		if placeType, review, err := models.GetReviewsForPlace(place); err == nil {
			// Call the render function with the title, review and the name of the
			// template
			render(c, gin.H{
				"placeType": placeType,
				"payload":   review}, "display_reviews_new.html")

		} else {
			// If the review is not found, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		// If an invalid review ID is specified in the URL, abort with an error
		c.AbortWithError(http.StatusNotFound, errors.New("All Params were not passed. Hence couldnt get review"))
	}
}

func CreateReview(c *gin.Context) {
	// Obtain the POSTed title and content values

	var review models.Review
	err := c.BindJSON(&review)
	log.Println("Review Json is ", review)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	if err = models.CreateNewReview(review); err == nil {
		// If the review is created successfully, show success message
		render(c, gin.H{
			"title": "Submission Successful",
		}, "submission-successful.html")
	} else {
		// if there was an error while creating the review, abort with an error
		c.AbortWithError(http.StatusBadRequest, err)
	}

}

func GetSafetyScore(c *gin.Context) {
	if place, err := createPlaceFromParams(c); err == nil {

		// Check if the review exists
		if _, reviewArray, err := models.GetReviewsForPlace(place); err == nil {
			// Call the render function with the title, review and the name of the
			// template

			fmt.Println(reviewArray)
			timeGroupedScores := make(map[string][]float32)
			for _, r := range reviewArray {
				if len(timeGroupedScores[r.VisitTime]) == 0 {
					timeGroupedScores[r.VisitTime] = []float32{r.Rating}
				} else {
					timeGroupedScores[r.VisitTime] = append(timeGroupedScores[r.VisitTime], r.Rating)
				}
			}
			scores := []SafetyScore{}

			for k, v := range timeGroupedScores {
				sum := float32(0)
				for _, num := range v {
					sum += num
				}
				safetyScore := (sum / float32(len(v)))
				timeOfTheDay, _ := strconv.Atoi(k)
				scores = append(scores, SafetyScore{Time: timeOfTheDay, Score: safetyScore})

			}

			fmt.Println(scores)
			c.JSON(http.StatusOK, gin.H{"safety_scores": scores})
			return

		} else {
			// If the review is not found, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		// If an invalid review ID is specified in the URL, abort with an error
		c.AbortWithError(http.StatusNotFound, errors.New("All Params were not passed. Hence couldnt get review"))
	}
}
