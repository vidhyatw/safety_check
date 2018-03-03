// handlers.review.go

package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	// reviews := getAllreviews()

	// // Call the render function with the name of the template to render
	// render(c, gin.H{
	// 	"title":   "Home Page",
	// 	"payload": reviews}, "index.html")
	render(c, gin.H{}, "landing-page.html")
}

func showreviewCreationPage(c *gin.Context) {
	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title": "Create New review"}, "create-review.html")
}

func getreview(c *gin.Context) {
	// Check if the review ID is valid
	if reviewID, err := strconv.Atoi(c.Param("review_id")); err == nil {
		// Check if the review exists
		if review, err := getReviewByID(reviewID); err == nil {
			// Call the render function with the title, review and the name of the
			// template
			render(c, gin.H{
				"title":   review.Title,
				"payload": review}, "review.html")

		} else {
			// If the review is not found, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		// If an invalid review ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func createreview(c *gin.Context) {
	// Obtain the POSTed title and content values
	title := c.PostForm("title")
	content := c.PostForm("content")

	if a, err := createNewReview(title, content); err == nil {
		// If the review is created successfully, show success message
		render(c, gin.H{
			"title":   "Submission Successful",
			"payload": a}, "submission-successful.html")
	} else {
		// if there was an error while creating the review, abort with an error
		c.AbortWithStatus(http.StatusBadRequest)
	}
}
