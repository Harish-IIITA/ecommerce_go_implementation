package controllers

import (
	"encoding/json"
	"fmt"
	"models"
	"net/http"
	"util"

	"service"

	"github.com/julienschmidt/httprouter"
)

type (
	// WeightController represents the controller for operating on the Weight resource
	WeightController struct {
	}
)

// NewWeightController creating a new WeightController instance
func NewWeightController() *WeightController {
	return &WeightController{}
}

// GetWeight retrieves an individual Weight resource
func (uc WeightController) GetWeight(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Get a WeightService instance
	us := service.NewWeightService(util.GetSession())
	u, err := us.GetWeight(id)
	// Verify id is ObjectId, otherwise bail
	if err {
		w.WriteHeader(404)
		return
	}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

// CreateWeight creates a new Weight resource
func (uc WeightController) CreateWeight(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Stub an Weight to be populated from the body
	u := models.Weight{}

	// Populate the Weight data
	json.NewDecoder(r.Body).Decode(&u)

	// Get a WeightService instance
	us := service.NewWeightService(util.GetSession())
	weight, err := us.CreateWeight(u)
	if err {
		w.WriteHeader(400)
		return
	}
	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(weight)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", uj)
}
