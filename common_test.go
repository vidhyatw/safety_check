package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/hackerearth/safetycheck/middleware"
)

var tmpUserList []user
var tmpreviewList []review

// This function is used to do setup before executing the test functions
func TestMain(m *testing.M) {
	//Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Run the other tests
	os.Exit(m.Run())
}

// Helper function to create a router during testing
func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("templates/*")
		r.Use(middleware.SetUserStatus())
	}
	return r
}

// Helper function to process a request and test its response
func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {

	// Create a response recorder
	w := httptest.NewRecorder()

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}

// This is a helper function that allows us to reuse some code in the above
// test methods
func testMiddlewareRequest(t *testing.T, r *gin.Engine, expectedHTTPCode int) {
	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/", nil)

	// Process the request and test the response
	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		return w.Code == expectedHTTPCode
	})
}

// This function is used to store the main lists into the temporary one
// for testing
func saveLists() {
	tmpUserList = userList
	tmpreviewList = reviewList
}

// This function is used to restore the main lists from the temporary one
func restoreLists() {
	userList = tmpUserList
	reviewList = tmpreviewList
}
