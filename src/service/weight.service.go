package service

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"models"
)

type (
	// WeightService represents the service for operating on the Weight resource
	WeightService struct {
		session *mgo.Session
	}
)

//NewWeightService creating a new WeightService instance
func NewWeightService(s *mgo.Session) *WeightService {
	return &WeightService{s}
}

// GetWeight retrieves an individual Weight resource
func (us WeightService) GetWeight(id string) (models.Weight, bool) {

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		return models.Weight{}, true
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Stub Weight
	u := models.Weight{}

	// Fetch Weight
	if err := us.session.DB("eCommerce").C("weights").FindId(oid).One(&u); err != nil {
		return models.Weight{}, true
	}

	return u, false
}

// CreateWeight creates a new Weight resource
func (us WeightService) CreateWeight(Weight models.Weight) (models.Weight, bool) {

	// Add an Id
	Weight.ID = bson.NewObjectId()
	c := us.session.DB("eCommerce").C("weights")
	// Write the Weight to mongo
	err := c.Insert(Weight)
	if err != nil {
		panic(err)
	}

	return Weight, false
}
