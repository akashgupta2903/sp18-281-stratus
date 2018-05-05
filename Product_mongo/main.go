package main

import (
    "fmt"
    "errors"
    "net/http"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "encoding/json"
    "strconv"
)
// MongoDB config
var mongodb_server = "mongodb://52.53.171.105:27017,13.57.212.90:27017,13.56.155.229:27017,54.183.172.60:27017,13.56.191.76:27017/?replicaSet=rs01"
var mongodb_database = "cmpe281"
var mongodb_collection = "products"

var ErrNoOrder = errors.New("Database error: no order found")

type Product struct {
    Product_id int
    Name  string
    Description string
    Ingredients string
    Category string
    Price  float64
    Likes  int
}



func main() {
    http.HandleFunc("/", pingHandler)
    http.HandleFunc("/getallproducts", GetProducts)
    http.HandleFunc("/like", LikeProduct)
    fmt.Println("Go server listening on port 4000...")
    http.ListenAndServe(":4000", nil)
}

/*
func LikeProduct(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    if r.Method != "PUT" {
        w.Header().Set("Allow", "PUT")
        http.Error(w, http.StatusText(405), 405)
        return
    }
    id := r.URL.Query().Get("id")


    err2 := UpdateProduct(id)
    if err2 == ErrNoOrder {
        fmt.Println("Server error: no order found in database")
        http.NotFound(w, r)
        return
    } else if err2 != nil {
        fmt.Println("Server error: other error encountered")
        http.Error(w, http.StatusText(500), 500)
        return
    }

    http.Redirect(w, r, "/getallproducts", 303)
}


func UpdateProduct(id string) (error) {
    session, err := mgo.Dial(mongodb_server)
    if err != nil {
        panic(err)
        return err
    }
    defer session.Close()
    session.SetMode(mgo.Monotonic, true)
    c := session.DB(mongodb_database).C(mongodb_collection)
    query := bson.M{"product_id" : id}
    fmt.Println(query)
    change := bson.M{"$set": bson.M{"likes" : 1}}
    err = c.Update(query, change)
    if err != nil {
        log.Fatal(err)
        return err
    }
    var result bson.M
    err = c.Find(bson.M{"product_id" : id}).One(&result)
    if err != nil {
        log.Fatal(err)
        return err
    }

    fmt.Println("Updated Product: ", result )
    return nil
}
*/

func pingHandler (w http.ResponseWriter, req *http.Request) {
		json.NewEncoder(w)
	}

func LikeProduct(w http.ResponseWriter, r *http.Request){
  productID := r.FormValue("id")
  id, err := strconv.Atoi(productID)
  session, err :=mgo.Dial(mongodb_server)
  if err != nil {
    panic(err)
      return
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  c := session.DB(mongodb_database).C(mongodb_collection)

  pdt := Product{}

  err = c.Find( bson.M{"product_id":id}).One(&pdt)

  fmt.Println(pdt.Name)

  query := bson.M{"product_id" : id}
  change := bson.M{"$inc": bson.M{"likes": 1}}
  err = c.Update(query, change)

  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  http.Redirect(w, r, "/getallproducts", 303)
}

func GetProducts(w http.ResponseWriter, r *http.Request){
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
    var products []Product

    err = c.Find(bson.M{}).All(&products)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    respondWithJson(w, http.StatusOK, products)
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
