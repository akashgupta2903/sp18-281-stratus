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
    http.Redirect(w, r, "/getdetail?id="+id, 303)
}