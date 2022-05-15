/*

import (
	// "fmt"
	"encoding/json"
	"fmt"
	"log"

	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shopspring/decimal"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)



var db *gorm.DB
var err error

func main() {
	dsb := "root:@/go_rest_api_crud?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsb), &gorm.Config{})
	if err != nil {
		log.Println("Connection Gagal", err)
	} else {
		log.Println("Connection Berhasil")
	}
	db.AutoMigrate(&Product{})

	handleRequest()
}

// product is a representation of a product
type Product struct {
	ID    int
	Code  string
	Name  string
	Price decimal.Decimal
}

type Result struct {
	Code    int
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func handleRequest() {
	log.Println("start the development server at http://127.0.0.1:8080")

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/api/products", createProduct).Methods("POST")
	myRouter.HandleFunc("/api/products", getProducts).Methods("GET")
	myRouter.HandleFunc("/api/products/{id}", getProduct).Methods("GET")
	myRouter.HandleFunc("/api/products/{id}", updateProduct).Methods("PUT")
	myRouter.HandleFunc("/api/products/{id}", deleteProduct).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome home!")
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	payload, _ := ioutil.ReadAll(r.Body)

	var product Product
	json.Unmarshal(payload, &product)

	// save if not exist
	db.Create(&product)

	res := Result{
		Code:    200,
		Data:    product,
		Message: "Success Create Product",
	}

	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	products := []Product{}

	db.Find(&products)

	res := Result{
		Code:    200,
		Data:    products,
		Message: "Success Get Products",
	}

	results, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["id"]

	var product Product
	db.First(&product, productID)

	res := Result{
		Code:    200,
		Data:    product,
		Message: "Success Get Products",
	}

	results, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["id"]

	payload, _ := ioutil.ReadAll(r.Body)

	var productUpdates Product
	json.Unmarshal(payload, &productUpdates)

	var product Product
	db.First(&product, productID)
	db.Model(&product).Updates(productUpdates)

	res := Result{
		Code:    200,
		Message: "Success Update Product",
		Data:    product,
	}

	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["id"]

	var product Product
	db.First(&product, productID)
	db.Delete(&product)

	res := Result{
		Code:    200,
		Message: "Success Delete Products",
	}

	results, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

*/
package main
