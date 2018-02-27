// models.review.go

package main

import "errors"

type review struct {
	ID             int    `json:"id"`
	Title          string `json:"title"`
	Rating         int    `json:"rating"`
	Content        string `json:"content"`
	ReviewerName   string `json:"reviewerName"`
	ReviewerAge    int    `json:"reviewerAge"`
	ReviewerGender string `json:"reviewerGender"`
}

// For this demo, we're storing the review list in memory
// In a real application, this list will most likely be fetched
// from a database or from static files
var reviewList = []review{
	review{ID: 1, Title: "Review 1", Content: "review 1 body", ReviewerGender: "female", Rating: 4, ReviewerName: "Person 1"},
	review{ID: 2, Title: "Review 2", Content: "review 2 body", ReviewerGender: "female", Rating: 4, ReviewerName: "Person 2"},
}

// Return a list of all the reviews
func getAllReviews() []review {
	return reviewList
}

// Fetch an review based on the ID supplied
func getReviewByID(id int) (*review, error) {
	for _, a := range reviewList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("review not found")
}

// Create a new review with the title and content provided
func createNewReview(title, content string) (*review, error) {
	// Set the ID of a new review to one more than the number of reviews
	a := review{ID: len(reviewList) + 1, Title: title, Content: content}

	// Add the review to the list of reviews
	reviewList = append(reviewList, a)

	return &a, nil
}
