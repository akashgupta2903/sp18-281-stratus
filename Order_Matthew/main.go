package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "strconv"
)

func main() {
    http.HandleFunc("/updateorder", UpdateOrderHandler)
    http.HandleFunc("/payorder", PayOrderHandler)
    http.HandleFunc("/cancelorder", CancelOrderHandler)
    fmt.Println("Go server listening on port 4000...")
    http.ListenAndServe(":4000", nil)
}

func UpdateOrderHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    if r.Method != "PUT" {
        w.Header().Set("Allow", "PUT")
        http.Error(w, http.StatusText(405), 405)
        return
    }

    var o Order
    err1 := json.NewDecoder(r.Body).Decode(&o)
    fmt.Println("Updating Order: ", o.Order_id, o.User_id, o.Timestamp, o.Status)
    if err1 != nil {
        fmt.Println("Server error: could not parse request body into Order type")
        http.Error(w, http.StatusText(400), 400)
        return
    }

    err2 := UpdateOrder(o)
    if err2 == ErrNoOrder {
        fmt.Println("Server error: no order found in database")
        http.NotFound(w, r)
        return
    } else if err2 != nil {
        fmt.Println("Server error: other error encountered")
        http.Error(w, http.StatusText(500), 500)
        return
    }

    http.StatusText(200)
}

func PayOrderHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    if r.Method != "POST" {
        w.Header().Set("Allow", "POST")
        http.Error(w, http.StatusText(405), 405)
        return
    }
    
    id := r.URL.Query().Get("id")
    if id == "" {
        fmt.Println("Error: the order ID is empty")
        http.Error(w, http.StatusText(400), 400)
        return
    }
    int_id, err := strconv.Atoi(id)
    if err != nil {
        http.Error(w, http.StatusText(400), 400)
        return
    }

    err2 := PayOrder(int_id)
    if err2 == ErrNoOrder {
        fmt.Println("Error: no order found in the database")
        http.NotFound(w, r)
        return
    } else if err2 != nil {
        fmt.Println("Error: some error(s) encountered")
        http.Error(w, http.StatusText(500), 500)
        return
    }

    http.StatusText(200)
}

func CancelOrderHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    if r.Method != "DELETE" {
        w.Header().Set("Allow", "DELETE")
        http.Error(w, http.StatusText(405), 405)
        return
    }
    
    id := r.URL.Query().Get("id")
    if id == "" {
        fmt.Println("Error: the order ID is empty")
        http.Error(w, http.StatusText(400), 400)
        return
    }
    int_id, err := strconv.Atoi(id)
    if err != nil {
        http.Error(w, http.StatusText(400), 400)
        return
    }

    err2 := CancelOrder(int_id)
    if err2 == ErrNoOrder {
        fmt.Println("Error: no order found in the database")
        http.NotFound(w, r)
        return
    } else if err2 != nil {
        fmt.Println("Error: some error(s) encountered")
        http.Error(w, http.StatusText(500), 500)
        return
    }
    
    http.StatusText(200)
}
