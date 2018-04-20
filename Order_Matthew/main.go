package main

import (
    "fmt"
    "net/http"
    "encoding/json"
)

func main() {
    http.HandleFunc("/updateorder", UpdateOrderHandler)
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
