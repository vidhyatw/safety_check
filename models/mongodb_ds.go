package models

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	collectionName = "reviews"
	databaseName   = "safetycheck"
	mongoDBHost    = "localhost"
)

type MongoDBDS struct {
	mongoSession *mgo.Session
}

func newMongoDBDS() MongoDBDS {
	cluster := mongoDBHost // mongodb host
	// connect to mongo
	session, err := mgo.Dial(cluster)
	if err != nil {
		log.Fatal("could not connect to db: ", err)
		panic(err)
	}
	defer session.Close()
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
func (mongoDS MongoDBDS) FindReviewsForPlace(place Place) ([]Review, error) {
	session := mongoDS.mongoSession.Copy()
	defer session.Close()
	c := session.DB(databaseName).C(collectionName)

	var reviews []Review
	long := place.Coordinates[0]
	lat := place.Coordinates[1]
	err := c.Find(bson.M{"place.coordinates": bson.M{"$near": []float64{long, lat}, "$maxDistance": 0.056}}).All(&reviews)
	return reviews, err
}
