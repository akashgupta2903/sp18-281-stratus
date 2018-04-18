package main

import (
    "net/http"
    "go-server/models"
    "strconv"
    "encoding/json"
)

func main() {
    http.HandleFunc("/getdetail", getDetail)
    http.HandleFunc("/like", addLike)
    http.HandleFunc("/getallproducts", getALL)
    http.HandleFunc("/popular", listPopular)
    http.ListenAndServe(":4000", nil)
}

func getDetail(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    // in case wrong http method is used
    if r.Method != "GET" {
        w.Header().Set("Allow", "GET")
        http.Error(w, http.StatusText(405), 405)
        return
    }
    //in case id is empty
    id := r.URL.Query().Get("id")
    if id == "" {
        http.Error(w, http.StatusText(400), 400)
        return
    }
    if _, err := strconv.Atoi(id); err != nil {
        http.Error(w, http.StatusText(400), 400)
        return
    }

    product, err := models.FindProduct(id)
    if err == models.ErrNoProduct {
        http.NotFound(w, r)
        return
    } else if err != nil {
        http.Error(w, http.StatusText(500), 500)
        return
    }
    json.NewEncoder(w).Encode(product)
}


func addLike(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    if r.Method != "POST" {
        w.Header().Set("Allow", "POST")
        http.Error(w, http.StatusText(405), 405)
        return
    }
    id := r.URL.Query().Get("id")
    if id == "" {
        http.Error(w, http.StatusText(400), 400)
        return
    }
    if _, err := strconv.Atoi(id); err != nil {
        http.Error(w, http.StatusText(400), 400)
        return
    }

    err := models.IncrementLikes(id)
    if err == models.ErrNoProduct {
        http.NotFound(w, r)
        return
    } else if err != nil {
        http.Error(w, http.StatusText(500), 500)
        return
    }
    http.Redirect(w, r, "/getdetail?id="+id, 303)
}

func getALL(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  if r.Method != "GET" {
    w.Header().Set("Allow", "GET")
    http.Error(w, http.StatusText(405), 405)
    return
  }
  abs, err := models.FindAll()
  if err != nil {
    http.Error(w, http.StatusText(500), 500)
    return
  }
  json.NewEncoder(w).Encode(abs)
  /*
  for i, ab := range abs {
    fmt.Fprintf(w, "%d) %s by %s: £%.2f [%d likes] \n", i+1, ab.Name, ab.Ingredients, ab.Price, ab.Likes)
  }
  */


func listPopular(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  if r.Method != "GET" {
    w.Header().Set("Allow", "GET")
    http.Error(w, http.StatusText(405), 405)
    return
  }

  abs, err := models.FindTopThree()
  if err != nil {
    http.Error(w, http.StatusText(500), 500)
    return
  }
  json.NewEncoder(w).Encode(abs)
/*
  for i, ab := range abs {
    fmt.Fprintf(w, "%d) %s by %s: £%.2f [%d likes] \n", i+1, ab.Title, ab.Artist, ab.Price, ab.Likes)
  }
*/
}
