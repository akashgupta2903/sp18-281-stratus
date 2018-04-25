package main

import (
	"fmt"
    "net/http"
    "go-server/models"
    "strconv"
    "encoding/json"
)

func main() {
    http.HandleFunc("/payorder", PayOrderHandler)
    http.HandleFunc("/cancelorder", CancelOrderHandler)
    fmt.Println("Go server is now listening on the port 4000...")
    http.ListenAndServe(":4000", nil)
}

func PayOrderHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    // the case when the http input method is wrong
    if r.Method != "POST" {
        w.Header().Set("Allow", "POST")
        http.Error(w, http.StatusText(405), 405)
        return
    }
    
    id := r.URL.Query().Get("id") //// ???? order id, product id????
    // the case when the id is empty
	if id == "" {
		fmt.Println("Error: the order ID is empty")
        http.Error(w, http.StatusText(400), 400)
        return
    }
    if _, err := strconv.Atoi(id); err != nil {
        http.Error(w, http.StatusText(400), 400)
        return
    }
    

    //update the status from "placed" to "paid" referencing to the order ID
    order, err := models.PayOrder(id)
    if err == models.ErrNoOrder {
    	fmt.Println("Error: no order found in the database")
        http.NotFound(w, r)
        return
    } else if err != nil {
    	fmt.Println("Error: some error(s) encountered")
        http.Error(w, http.StatusText(500), 500)
        return
    }
    json.NewEncoder(w).Encode(order)
}

func CancelOrderHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    // the case when the http input method is wrong
    if r.Method != "DELETE" {
        w.Header().Set("Allow", "DELETE")
        http.Error(w, http.StatusText(405), 405)
        return
    }
    
    id := r.URL.Query().Get("id") //// ???? order id, product id????
    // the case when the id is empty
	if id == "" {
		fmt.Println("Error: the order ID is empty")
        http.Error(w, http.StatusText(400), 400)
        return
    }
    if _, err := strconv.Atoi(id); err != nil {
        http.Error(w, http.StatusText(400), 400)
        return
    }
    
    //Deleting the order referencing to the order ID
    order, err := models.DeleteOrder(id)
    if err == models.ErrNoOrder {
    	fmt.Println("Error: no order found in the database")
        http.NotFound(w, r)
        return
    } else if err != nil {
    	fmt.Println("Error: some error(s) encountered")
        http.Error(w, http.StatusText(500), 500)
        return
    }
    json.NewEncoder(w).Encode(order)
    }
    
