package models

type Review struct {
	ID        int      `json:"id"`
	Title     string   `form:"title" json:"title" binding:"omitempty,len=0"`
	Rating    int      `form:"rating" json:"rating" binding:"required"`
	Content   string   `form:"content" json:"content" binding:"omitempty,len=0"`
	TimeStamp string   `form:"timestamp" json:"timestamp" binding:"required"`
	VisitTime string   `form:"visitTime" json:"visitTime" binding:"required"`
	Place     Place    `form:"place" json:"place" binding:"required"`
	Reviewer  Reviewer `form:"reviewer" json:"reviewer" binding:"required"`
	// Votes     Votes    `form:"user" json:"votes,omitempty" binding:"omitempty,len=0"`
}
type Reviewer struct {
	Name   string `form:"name" json:"name" binding:"omitempty,len=0"`
	Gender string `form:"gender" json:"gender" binding:"omitempty,len=0"`
	Age    int    `form:"age" json:"age" binding:"omitempty,len=0"`
	Email  string `form:"email" json:"email" binding:"omitempty,len=0"`
}

type Place struct {
	PlaceID     string    `form:"user" json:"placeid" binding:"required"`
	Type        string    `form:"user" json:"type" binding:"omitempty,len=0"`
	Coordinates []float64 `form:"user" json:"coordinates" binding:"required"`
}

type Votes struct {
	Thumbsup   int `form:"user" json:"title,omitempty" binding:"omitempty,len=0"`
	Thumbsdown int `form:"user" json:"title,omitempty" binding:"omitempty,len=0"`
}

// Fetch an review based on the ID supplied
func GetReviewsForPlace(place Place) ([]Review, error) {
	ds := GetDatasource()
	return ds.FindReviewsForPlace(place)
}

// Create a new review with the title and content provided
func CreateNewReview(newReview Review) error {
	ds := GetDatasource()
	return ds.CreateReview(newReview)
}
