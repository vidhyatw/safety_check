package models

import (
	"log"

	"github.com/hackerearth/safety_check/config"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	collectionName = "reviews"
	databaseName   = "heroku_38bq7nx5"
)

type MongoDBDS struct {
	mongoSession *mgo.Session
}

func newMongoAtlasDBDS() MongoDBDS {
	session, err := mgo.Dial(config.GetDatabaseURI())

	if err != nil {
		log.Fatal("could not connect to db: ", err)
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	return MongoDBDS{session}
}

func (mongoDS MongoDBDS) CreateReview(review Review) error {
	session := mongoDS.mongoSession.Copy()
	defer session.Close()
	c := session.DB(databaseName).C(collectionName)

	err := c.Insert(review)

	if err != nil {
		log.Fatal("Unable to createReview", err)
		return err
	}

	return nil
}
func (mongoDS MongoDBDS) FindReviewsForPlace(place Place) (string, []Review, error) {
	session := mongoDS.mongoSession.Copy()
	defer session.Close()
	c := session.DB(databaseName).C(collectionName)

	var reviews []Review
	err := c.Find(bson.M{"place.placeid": place.PlaceID}).All(&reviews)
	if err == nil && len(reviews) == 0 {
		return mongoDS.FindReviewsForNearByPlace(place)
	}
	return "PLACE", reviews, err
}

func (mongoDS MongoDBDS) FindReviewsForNearByPlace(place Place) (string, []Review, error) {
	session := mongoDS.mongoSession.Copy()
	defer session.Close()
	c := session.DB(databaseName).C(collectionName)

	var reviews []Review
	long := place.Coordinates[0]
	lat := place.Coordinates[1]
	scope := 2000
	err := c.Find(bson.M{
		"place.coordinates": bson.M{
			"$nearSphere": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{long, lat},
				},
				"$maxDistance": scope,
			},
		},
	}).All(&reviews)

	// err := c.Find(bson.M{"place.coordinates": bson.M{"$near": []float64{long, lat}, "$maxDistance": 2000}}).All(&reviews)
	return "NEARBY", reviews, err
}
