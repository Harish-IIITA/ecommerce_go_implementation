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
	// Input represents the structure of input JSON
	Input struct {
		User    models.User        `json:"user"`
		Product models.ProductJSON `json:"product"`
	}
)

type (
	// ProductController represents the controller for operating on the Product resource
	ProductController struct {
	}
)

// NewProductController creating a new ProductController instance
func NewProductController() *ProductController {
	return &ProductController{}
}

// GetProduct retrieves an individual product resource
func (pc ProductController) GetProduct(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Get a ProductService instance
	ps := service.NewProductService(util.GetSession())
	u, err := ps.GetProduct(id)
	// Verify id is ObjectId, otherwise bail
	if err {
		w.WriteHeader(404)
		return
	}

	// Marshal provided interface into JSON structure
	pj, _ := json.Marshal(u)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", pj)
}

// GetAllProduct retrieves all products resource
func (pc ProductController) GetAllProduct(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// Get a ProductService instance
	ps := service.NewProductService(util.GetSession())
	u, err := ps.GetAllProducts()
	// Verify id is ObjectId, otherwise bail
	if err {
		w.WriteHeader(404)
		return
	}

	// Marshal provided interface into JSON structure
	pj, _ := json.Marshal(u)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", pj)
}

// CreateProduct creates a new product resource
func (pc ProductController) CreateProduct(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Stub an product to be populated from the body
	pr := models.Product{}

	// Populate the product data
	json.NewDecoder(r.Body).Decode(&pr)

	// Get a ProductService instance
	ps := service.NewProductService(util.GetSession())
	product, err := ps.CreateProduct(pr)
	if err {
		w.WriteHeader(400)
		return
	}
	// Marshal provided interface into JSON structure
	pj, _ := json.Marshal(product)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", pj)
}

// RemoveProduct removes an existing product resource
func (pc ProductController) RemoveProduct(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Get a ProductService instance
	ps := service.NewProductService(util.GetSession())
	err := ps.RemoveProduct(id)
	// Verify id is ObjectId, otherwise bail
	if err {
		w.WriteHeader(400)
		return
	}

	// Write status
	w.WriteHeader(200)
}

// MatchProducts provides range of matching products
func (pc ProductController) MatchProducts(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// Stub an input to be populated from the body
	input := Input{}

	// Populate the inout data
	json.NewDecoder(r.Body).Decode(&input)

	// Stub an product to be populated from the body
	pr := input.Product

	// Stub an user to be populated from the body
	u := input.User

	fmt.Println(input)

	// Get a ProductService instance
	ps := service.NewProductService(util.GetSession())
	products, err := ps.MatchProduct(pr, u)
	if err {
		w.WriteHeader(400)
		return
	}
	// Marshal provided interface into JSON structure
	pj, _ := json.Marshal(products)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", pj)
}
