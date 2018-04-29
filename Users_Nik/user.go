package main

import (
    
   
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
   
)

// MongoDB Config
var mongodb_server = "localhost"
var mongodb_database = "cmpe281"
var mongodb_collection = "users"


type User struct {
    User_id string `bson:"_id" json:"id"`
    UserName string `bson:"username" json:"username"`
    Email  string `bson:"email" json:"email"`
    Password string `bson:"password" json:"password"`
}

var db *mgo.Database

func connect() *mgo.Database{
    // Create a session which maintains a pool of socket connections
	// to our MongoDB.
	session, err := mgo.Dial(mongodb_server)
    if err != nil {
        panic(err)

    }
    defer session.Close()
    session.SetMode(mgo.Monotonic, true)
 
   db=session.DB(mongodb_database)


    return db
}

func collection() *mgo.Collection {

  db=connect()

	 return db.C(mongodb_collection)
}




// Select * from users where username=""
func FindUser(username string) (User,error) {
  u := User{}

  c := collection()

  err := c.Find(bson.M{"username":username}).One(&u)
  
  if  err != nil {
    return u, err
  }
  return u, nil

}





