package models

import "gopkg.in/mgo.v2/bson"

type (
	// Weight represents the structure of our resource
	Weight struct {
		ID            bson.ObjectId `json:"id" bson:"_id"`
		ProductName   int           `json:"productName" bson:"productName"`
		Color         int           `json:"color" bson:"color"`
		Gender        int           `json:"gender" bson:"gender"`
		AgeGroup      int           `json:"ageGroup" bson:"ageGroup"`
		Size          int           `json:"size" bson:"size"`
		EventType     int           `json:"eventType" bson:"eventType"`
		TraditionType int           `json:"traditionType" bson:"traditionType"`
	}
)
