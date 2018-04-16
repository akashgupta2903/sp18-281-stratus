package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Order struct
type Order struct {

}

// User struct
type User struct {

}

// Main function
func main() {
	// Init router
	r := mux.NewRouter()

	// Route handles & endpoints
	r.HandleFunc("/login", login).Methods("POST")
	r.HandleFunc("/signup", signup).Methods("POST")
	r.HandleFunc("/order", createOrder).Methods("POST")
	r.HandleFunc("/order/{id}", getOrder).Methods("GET")
	r.HandleFunc("/order/{id}/pay", payOrder).Methods("POST")
	r.HandleFunc("/delete/{id}", deleteOrder).Methods("POST")
	r.HandleFunc("/update/{id}", updateOrder).Methods("POST")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}

// user login
func login(w http.ResponseWriter, r *http.Request) {

}

// user signup
func signup(w http.ResponseWriter, r *http.Request) {

}

// create order
func createOrder(w http.ResponseWriter, r *http.Request) {

}

// get order
func getOrder(w http.ResponseWriter, r *http.Request) {

}

// pay order
func payOrder(w http.ResponseWriter, r *http.Request) {

}

// delete order
func deleteOrder(w http.ResponseWriter, r *http.Request) {

}

// update order
func updateOrder(w http.ResponseWriter, r *http.Request) {

}