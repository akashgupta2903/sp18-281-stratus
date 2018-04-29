package main

import (
    "fmt"
    "time"
    "errors"
    "log"
    "strconv"
    "net/http"
    "encoding/json"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

// MongoDB config
var mongodb_server = "localhost"
var mongodb_database = "cmpe281"
var mongodb_collection = "orders"

var ErrNoOrder = errors.New("Database error: no order found")

type Order struct {
    Order_id int
    User_id int
    Timestamp string
    Status string
    Items []int
}

func main() {
    http.HandleFunc("/updateorder", UpdateOrderHandler)
    http.HandleFunc("/payorder", PayOrderHandler)
    http.HandleFunc("/cancelorder", CancelOrderHandler)
    http.HandleFunc("/placeorder", PlaceOrderHandler)
    http.HandleFunc("/getorders", GetOrdersHandler)
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

func UpdateOrder(order Order) (error) {
    session, err := mgo.Dial(mongodb_server)
    if err != nil {
        panic(err)
        return err
    }
    defer session.Close()
    session.SetMode(mgo.Monotonic, true)
    c := session.DB(mongodb_database).C(mongodb_collection)
    query := bson.M{"order_id" : order.Order_id}
    change := bson.M{"$set": bson.M{ "items" : order.Items,
        "timestamp": time.Now().Format(time.RFC822) }}
    err = c.Update(query, change)
    if err != nil {
        log.Fatal(err)
        return err
    }
    var result bson.M
    err = c.Find(bson.M{"order_id" : order.Order_id}).One(&result)
    if err != nil {
        log.Fatal(err)
        return err
    }    

    fmt.Println("Updated Order: ", result )
    return nil
}

func PayOrder(id int) (error) {
    session, err := mgo.Dial(mongodb_server)
    if err != nil {
        panic(err)
        return err
    }
    defer session.Close()
    session.SetMode(mgo.Monotonic, true)
    c := session.DB(mongodb_database).C(mongodb_collection)
    query := bson.M{"order_id" : id}
    change := bson.M{"$set": bson.M{ "status" : "Paid",
        "timestamp": time.Now().Format(time.RFC822) }}
    err = c.Update(query, change)
    if err != nil {
        log.Fatal(err)
        return err
    }
    var result bson.M
    err = c.Find(bson.M{"order_id" : id}).One(&result)
    if err != nil {
        log.Fatal(err)
        return err
    }

    fmt.Println("Paid Order: ", result)
    return nil
}

func CancelOrder(id int) (error) {
    session, err := mgo.Dial(mongodb_server)
    if err != nil {
        panic(err)
        return err
    }
    defer session.Close()
    session.SetMode(mgo.Monotonic, true)
    c := session.DB(mongodb_database).C(mongodb_collection)
    query := bson.M{"order_id" : id}
    change := bson.M{"$set": bson.M{ "status" : "Canceled",
        "timestamp": time.Now().Format(time.RFC822) }}
    err = c.Update(query, change)
    if err != nil {
        log.Fatal(err)
        return err
    }
    var result bson.M
    err = c.Find(bson.M{"order_id" : id}).One(&result)
    if err != nil {
        log.Fatal(err)
        return err
    }

    fmt.Println("Canceled Order: ", result)
    return nil
}

//GET endpoint to get all orders
func GetOrdersHandler(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")
    if r.Method != "GET" {
        w.Header().Set("Allow", "GET")
        http.Error(w, http.StatusText(405), 405)
        return
    }

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
func PlaceOrderHandler(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")
    if r.Method != "POST" {
        w.Header().Set("Allow", "POST")
        http.Error(w, http.StatusText(405), 405)
        return
    }

    session, err := mgo.Dial(mongodb_server)
    if err != nil {
        panic(err)
            return
    }

    defer session.Close()
    session.SetMode(mgo.Monotonic, true)
    c := session.DB(mongodb_database).C(mongodb_collection)
    var order Order

    if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid request payload")   
        return
    }

    // Format the timestamp
    order.Timestamp = time.Now().Format(time.RFC822)

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