package main

import (
	"fmt"	
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
)

type Order struct {
    Order_id int
    User_id int
    Timestamp string
    Status string
    Items []int
}

var mongodb_server = "localhost"
var mongodb_database = "cmpe281"
var mongodb_collection = "orders"

//GET endpoint to get all orders
func AllOrdersEndpoint(w http.ResponseWriter, r *http.Request){
	session, err := mgo.Dial(mongodb_server)
	if err != nil {
		panic(err)
	        return
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodb_database).C(mongodb_collection)
	var orders []Order
	err = c.Find(bson.M{}).All(&orders)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, orders)	
}

//POST endpoint to insert a new order
func CreateOrderEndpoint(w http.ResponseWriter, r *http.Request){
	session, err := mgo.Dial(mongodb_server)
	if err != nil {
		panic(err)
	        return
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodb_database).C(mongodb_collection)
	var order Order
	fmt.Printf("reached here")
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		fmt.Printf("in error")		
		return
	}
	//order.Order_id = bson.NewObjectId()
	err = c.Insert(&order)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, order)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}


func main() {
	r := mux.NewRouter()
	r.HandleFunc("/orders", AllOrdersEndpoint).Methods("GET")
	r.HandleFunc("/orders", CreateOrderEndpoint).Methods("POST")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}


