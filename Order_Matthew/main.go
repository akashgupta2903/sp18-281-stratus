package main

import (
    "fmt"
    "net/http"
    "strconv"
    "encoding/json"
)

func main() {
    http.HandleFunc("/updateorder", updateOrder)
    fmt.Println("Go server listening on port 4000...")
    http.ListenAndServe(":4000", nil)
}

func updateOrder(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    if r.Method != "PUT" {
        w.Header().Set("Allow", "PUT")
        http.Error(w, http.StatusText(405), 405)
        return
    }

    var o Order
    _ = json.NewDecoder(r.Body).Decode(&o)
    fmt.Println("Updating Order ", o.order_id)
    if o == nil {
        fmt.Println("Server error: could not parse request body into Order type")
        http.Error(w, http.StatusText(400), 400)
        return
    }

    err := UpdateOrder(o)
    if err == models.ErrNoOrder {
        fmt.Println("Server error: no order found in database")
        http.NotFound(w, r)
        return
    } else if err != nil {
        fmt.Println("Server error: other error encountered")
        http.Error(w, http.StatusText(500), 500)
        return
    }
    
    w.Write([]byte("OK"))
}
