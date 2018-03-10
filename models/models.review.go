package models

type Review struct {
	Title     string   `form:"title" json:"title" binding:"omitempty"`
	Rating    float32  `form:"rating" json:"rating" binding:"required"`
	Content   string   `form:"content" json:"content" binding:"omitempty"`
	TimeStamp float32  `form:"timestamp" json:"timestamp" binding:"required"`
	VisitTime string   `form:"visitTime" json:"visitTime" binding:"required"`
	Place     Place    `form:"place" json:"place" binding:"required"`
	Reviewer  Reviewer `form:"reviewer" json:"reviewer" binding:"required"`
	// Votes     Votes    `form:"user" json:"votes,omitempty" binding:"omitempty"`
}
type Reviewer struct {
	Name   string `form:"name" json:"name" binding:"omitempty"`
	Gender string `form:"gender" json:"gender" binding:"omitempty"`
	Age    int    `form:"age" json:"age" binding:"omitempty"`
	Email  string `form:"email" json:"email" binding:"omitempty"`
}

type Place struct {
	PlaceID     string    `form:"user" json:"placeid" binding:"required"`
	Type        string    `form:"user" json:"type" binding:"omitempty"`
	Coordinates []float64 `form:"user" json:"coordinates" binding:"required"`
}

type Votes struct {
	Thumbsup   int `form:"user" json:"title,omitempty" binding:"omitempty"`
	Thumbsdown int `form:"user" json:"title,omitempty" binding:"omitempty"`
}

// Fetch an review based on the ID supplied
func GetReviewsForPlace(place Place) (string, []Review, error) {
	ds := GetDatasource()
	return ds.FindReviewsForPlace(place)
}

// Create a new review with the title and content provided
func CreateNewReview(newReview Review) error {
	ds := GetDatasource()
	return ds.CreateReview(newReview)
}
