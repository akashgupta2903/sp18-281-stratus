package main

import (
    "fmt"
    "time"
    "errors"
    "log"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

// MongoDB Config
var mongodb_server = "localhost"
var mongodb_database = "cmpe281"
var mongodb_collection = "orders"

var ErrNoOrder = errors.New("Database error: no order found")

type Order struct {
    order_id int
    timestamp string
    status string
    items []int64
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
    query := bson.M{"order_id" : order.order_id}
    change := bson.M{"$set": bson.M{ "items" : order.items,
        "timestamp": time.Now().Format(time.RFC822) }}
    err = c.Update(query, change)
    if err != nil {
        log.Fatal(err)
        return err
    }
    var result bson.M
    err = c.Find(bson.M{"order_id" : order.order_id}).One(&result)
    if err != nil {
        log.Fatal(err)
        return err
    }        
    fmt.Println("New Order:", result )
    return nil
}