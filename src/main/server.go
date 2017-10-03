package main

import (
	// Standard library packages
	"net/http"

	// Third party packages
	"controllers"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// Instantiate a new router
	r := httprouter.New()

	// Get a UserController instance
	uc := controllers.NewUserController()

	// Get a user resource
	r.GET("/users/:id", uc.GetUser)

	// Get aLL user resource
	r.GET("/users", uc.GetAllUser)

	r.POST("/users", uc.CreateUser)

	r.DELETE("/users/:id", uc.RemoveUser)

	///////////////////////////////////

	//Get a ProdcutController instance
	pc := controllers.NewProductController()

	// Get a product resource
	r.GET("/products/:id", pc.GetProduct)

	// Get aLL product resource
	r.GET("/products", pc.GetAllProduct)

	r.POST("/products", pc.CreateProduct)

	r.DELETE("/products/:id", pc.RemoveProduct)

	r.POST("/matches", pc.MatchProducts)

	///////////////////////////////////

	wc := controllers.NewWeightController()

	r.POST("/weights", wc.CreateWeight)

	// Fire up the server
	http.ListenAndServe("localhost:3000", r)
}
