package models

import "gopkg.in/mgo.v2/bson"

type (
	// Product represents the structure of our resource
	Product struct {
		ID            bson.ObjectId `json:"id" bson:"_id"`
		ProductName   string        `json:"productName" bson:"productName"`
		Color         string        `json:"color" bson:"color"`
		Gender        string        `json:"gender" bson:"gender"`
		AgeGroup      string        `json:"ageGroup" bson:"ageGroup"`
		Size          string        `json:"size" bson:"size"`
		EventType     string        `json:"eventType" bson:"eventType"`
		TraditionType string        `json:"traditionType" bson:"traditionType"`
	}
)
