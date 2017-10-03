package service

import (
	"fmt"
	"sort"

	"github.com/fatih/structs"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"models"
)

type (
	// ProductService represents the service for operating on the Product resource
	ProductService struct {
		session *mgo.Session
	}
)

//NewProductService creating a new ProductService instance
func NewProductService(s *mgo.Session) *ProductService {
	return &ProductService{s}
}

// GetProduct retrieves an individual product resource
func (ps ProductService) GetProduct(id string) (models.Product, bool) {

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		return models.Product{}, true
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Stub product
	p := models.Product{}

	// Fetch product
	if err := ps.session.DB("eCommerce").C("products").FindId(oid).One(&p); err != nil {
		return models.Product{}, true
	}

	return p, false
}

// GetAllProducts retrieves all product resources
func (ps ProductService) GetAllProducts() ([]models.Product, bool) {

	// Stub product
	//u := models.Product{}
	var products []models.Product
	// Fetch products
	if err := ps.session.DB("eCommerce").C("products").Find(nil).All(&products); err != nil {
		return []models.Product{}, true
	}

	return products, false
}

// CreateProduct creates a new product resource
func (ps ProductService) CreateProduct(product models.Product) (models.Product, bool) {

	// Add an Id
	product.ID = bson.NewObjectId()
	c := ps.session.DB("eCommerce").C("products")
	// Write the product to mongo
	err := c.Insert(product)
	if err != nil {
		panic(err)
	}

	return product, false
}

// RemoveProduct removes an existing product resource
func (ps ProductService) RemoveProduct(id string) bool {

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		return true
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Remove product
	if err := ps.session.DB("eCommerce").C("products").RemoveId(oid); err != nil {
		return true
	}

	return false
}

// MatchProduct matches a new product resource
func (ps ProductService) MatchProduct(product models.ProductJSON, user models.User) ([]models.Product, bool) {

	allProducts, val := ps.GetAllProducts()

	fmt.Println(val)

	var allUsers []models.User
	ps.session.DB("eCommerce").C("users").Find(nil).All(&allUsers)

	// Stub weight
	weight := models.Weight{}
	//ps.session.DB("eCommerce").C("weights").Find(bson.M{"$natural": -1}).One(&weight)

	c := ps.session.DB("eCommerce").C("weights")
	dbSize, err := c.Count()
	if err != nil {
		panic(err)
	}

	c.Find(nil).Skip(dbSize - 1).One(&weight)

	w := structs.Map(weight)
	delete(w, "ID")

	w2 := map[string]int{}

	for key, item := range w {
		w2[key] = item.(int)
	}

	type kv struct {
		Key   string
		Value int
	}
	var ss []kv
	for k, v := range w2 {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value < ss[j].Value
	})

	fmt.Println(product)
	fmt.Println(user)
	for _, kv := range ss {
		fmt.Printf("%s, %d\n", kv.Key, kv.Value)

		if kv.Key == "Color" {
			allProducts = colorFilter(product, allProducts)
		} else if kv.Key == "Gender" {
			allProducts = genderFilter(user, allProducts)
		} else if kv.Key == "AgeGroup" {
			allProducts = ageGroupFilter(user, allProducts)
		} else if kv.Key == "Size" {
			allProducts = sizeFilter(product, allProducts)
		} else if kv.Key == "EventType" {
			allProducts = eventTypeFilter(product, allProducts)
		} else if kv.Key == "ProductName" {
			allProducts = productNameFilter(product, allProducts)
		} else {
			allProducts = traditionTypeFilter(product, allProducts)
		}
	}

	return allProducts[:15], false
}

func colorFilter(product models.ProductJSON, products []models.Product) (reProducts []models.Product) {

	var prods1 []models.Product
	var prods2 []models.Product
	for i := range products {
		if products[i].Color == product.Color {
			prods1 = append(prods1, products[i])
		} else {
			prods2 = append(prods2, products[i])
		}
	}

	prods1 = append(prods1, prods2...)
	return prods1
}

func productNameFilter(product models.ProductJSON, products []models.Product) (reProducts []models.Product) {

	var prods1 []models.Product
	var prods2 []models.Product
	for i := range products {
		if products[i].ProductName == product.ProductName {
			prods1 = append(prods1, products[i])
		} else {
			prods2 = append(prods2, products[i])
		}
	}

	prods1 = append(prods1, prods2...)
	return prods1
}

func genderFilter(user models.User, products []models.Product) (reProducts []models.Product) {

	var prods1 []models.Product
	var prods2 []models.Product
	for i := range products {
		if products[i].Gender == user.Gender {
			prods1 = append(prods1, products[i])
		} else {
			prods2 = append(prods2, products[i])
		}
	}

	prods1 = append(prods1, prods2...)
	return prods1
}

func sizeFilter(product models.ProductJSON, products []models.Product) (reProducts []models.Product) {

	var prods1 []models.Product
	var prods2 []models.Product
	for i := range products {
		if products[i].Size == product.Size {
			prods1 = append(prods1, products[i])
		} else {
			prods2 = append(prods2, products[i])
		}
	}

	prods1 = append(prods1, prods2...)
	return prods1
}

func ageGroupFilter(user models.User, products []models.Product) (reProducts []models.Product) {

	var prods1 []models.Product
	var prods2 []models.Product

	var ageGroup string
	if user.Age < 18 {
		ageGroup = "15"
	} else if user.Age >= 18 && user.Age <= 22 {
		ageGroup = "20"
	} else {
		ageGroup = "25"
	}

	for i := range products {
		if products[i].AgeGroup == ageGroup {
			prods1 = append(prods1, products[i])
		} else {
			prods2 = append(prods2, products[i])
		}
	}

	prods1 = append(prods1, prods2...)
	return prods1
}

func eventTypeFilter(product models.ProductJSON, products []models.Product) (reProducts []models.Product) {

	var prods1 []models.Product
	var prods2 []models.Product
	for i := range products {
		if products[i].EventType == product.EventType {
			prods1 = append(prods1, products[i])
		} else {
			prods2 = append(prods2, products[i])
		}
	}

	prods1 = append(prods1, prods2...)
	return prods1
}

func traditionTypeFilter(product models.ProductJSON, products []models.Product) (reProducts []models.Product) {

	var prods1 []models.Product
	var prods2 []models.Product
	for i := range products {
		if products[i].TraditionType == product.TraditionType {
			prods1 = append(prods1, products[i])
		} else {
			prods2 = append(prods2, products[i])
		}
	}

	prods1 = append(prods1, prods2...)
	return prods1
}
