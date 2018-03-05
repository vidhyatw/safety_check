// handlers.review.go

package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/hackerearth/safetycheck/models"

	"github.com/gin-gonic/gin"
)

func ShowIndexPage(c *gin.Context) {
	// reviews := getAllreviews()

	// // Call the render function with the name of the template to render
	// render(c, gin.H{
	// 	"title":   "Home Page",
	// 	"payload": reviews}, "index.html")
	render(c, gin.H{}, "landing-page.html")
}

func ShowReviewCreationPage(c *gin.Context) {
	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title": "Create New review"}, "create-review.html")
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
				"payload":   review}, "index.html")

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

	if err := c.BindJSON(&review); err == nil {
		if err := models.CreateNewReview(review); err == nil {
			// If the review is created successfully, show success message
			render(c, gin.H{
				"title": "Submission Successful",
			}, "submission-successful.html")
		} else {
			// if there was an error while creating the review, abort with an error
			c.AbortWithStatus(http.StatusBadRequest)
		}

	} else {

	}
}
