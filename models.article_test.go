// models.review_test.go

package main

import "testing"

// Test the function that fetches all reviews
func TestGetAllreviews(t *testing.T) {
	alist := getAllreviews()

	// Check that the length of the list of reviews returned is the
	// same as the length of the global variable holding the list
	if len(alist) != len(reviewList) {
		t.Fail()
	}

	// Check that each member is identical
	for i, v := range alist {
		if v.Content != reviewList[i].Content ||
			v.ID != reviewList[i].ID ||
			v.Title != reviewList[i].Title {

			t.Fail()
			break
		}
	}
}

// Test the function that fetche an review by its ID
func TestGetreviewByID(t *testing.T) {
	a, err := getreviewByID(1)

	if err != nil || a.ID != 1 || a.Title != "review 1" || a.Content != "review 1 body" {
		t.Fail()
	}
}

// Test the function that creates a new review
func TestCreateNewreview(t *testing.T) {
	// get the original count of reviews
	originalLength := len(getAllreviews())

	// add another review
	a, err := createNewreview("New test title", "New test content")

	// get the new count of reviews
	allreviews := getAllreviews()
	newLength := len(allreviews)

	if err != nil || newLength != originalLength+1 ||
		a.Title != "New test title" || a.Content != "New test content" {

		t.Fail()
	}
}
