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

var ErrNoProduct = errors.New("models: no coffee found")


type Coffee struct {
    Name  string
    Ingredients string
    Price  float64
    Likes  int
}


func FindCoffee(id string) (*Coffee, error) {

}

func populateEntry(reply map[string]string) (*Coffee, error) {

}
