package models

import (
    "errors"
    "github.com/mediocregopher/radix.v2/pool"
    "log"
    "strconv"
)

// Declare a global db variable to store the Redis connection pool.
var db *pool.Pool

func init() {
    var err error
    // Establish a pool of 10 connections to the Redis server listening on
    // port 6379 of the local machine.
    db, err = pool.New("tcp", "localhost:6379", 10)
    if err != nil {
        log.Panic(err)
    }
}

var ErrNoProduct = errors.New("models: no product found")


type Product struct {
    Product_id int
    Name  string
    Description string
    Ingredients string
    Category string
    Price  float64
    Likes  int
}


func FindProduct(id string) (*Product, error) {
    // Use the connection pool's Get() method to fetch a single Redis
    // connection from the pool.
    conn, err := db.Get()
    if err != nil {
        return nil, err
    }
    defer db.Put(conn)

    reply, err := conn.Cmd("HGETALL", "product:"+id).Map()
    if err != nil {
        return nil, err
    } else if len(reply) == 0 {
        return nil, ErrNoProduct
    }

    return populateEntry(reply)
}

func populateEntry(reply map[string]string) (*Product, error) {
    var err error
    product := new(Product)
    //product.Product_id = reply["product_id"]
    product.Name = reply["name"]
    product.Description = reply["description"]
    product.Ingredients = reply["ingredients"]
    product.Category = reply["category"]
    product.Price, err = strconv.ParseFloat(reply["price"], 64)
    if err != nil {
        return nil, err
    }
    product.Likes, err = strconv.Atoi(reply["likes"])
    if err != nil {
        return nil, err
    }
    product.Product_id, err = strconv.Atoi(reply["product_id"])
    if err != nil {
        return nil, err
    }
    return product, nil
}
