package main

import (
    "net/http"
    "strconv"
    "encoding/json"
    "errors"
    "github.com/mediocregopher/radix.v2/pool"
    "github.com/mediocregopher/radix.v2/redis"
    "log"
)

func main() {
    http.HandleFunc("/", pingHandler)
    http.HandleFunc("/getdetail", getDetail)
    http.HandleFunc("/like", addLike)
    http.HandleFunc("/getallproducts", getALL)
    http.HandleFunc("/popular", listPopular)
    http.ListenAndServe(":4000", nil)
}


/*

func pingHandler (w http.ResponseWriter, req *http.Request) {
        json.NewEncoder(w)
}

*/

// for load balancer's health check
func pingHandler (w http.ResponseWriter, req *http.Request) {
        json.NewEncoder(w)
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

    product, err := FindProduct(id)
    if err == ErrNoProduct {
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

    err := IncrementLikes(id)
    if err == ErrNoProduct {
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
  abs, err := FindAll()
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

}


func listPopular(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  if r.Method != "GET" {
    w.Header().Set("Allow", "GET")
    http.Error(w, http.StatusText(405), 405)
    return
  }

  abs, err := FindTopThree()
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



var db *pool.Pool

func init() {
    var err error
    // Establish a pool of 10 connections to the Redis server listening on
    // port 6379 of the local machine.
    db, err = pool.New("tcp", "13.57.12.164:6379" , 10)
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

func IncrementLikes(id string) error {
    conn, err := db.Get()
    if err != nil {
        return err
    }
    defer db.Put(conn)

    exists, err := conn.Cmd("EXISTS", "product:"+id).Int()
    if err != nil {
        return err
    } else if exists == 0 {
        return ErrNoProduct
    }
    err = conn.Cmd("MULTI").Err
    if err != nil {
        return err
    }
    err = conn.Cmd("HINCRBY", "product:"+id, "likes", 1).Err
    if err != nil {
        return err
    }
    err = conn.Cmd("ZINCRBY", "likes", 1, id).Err
    if err != nil {
        return err
    }
    err = conn.Cmd("EXEC").Err
    if err != nil {
        return err
    }
    return nil
}


func FindAll() ([]*Product, error) {
    conn, err := db.Get()
    if err != nil {
        return nil, err
    }
    defer db.Put(conn)

    for {
        // Instruct Redis to watch the likes sorted set for any changes.
        err = conn.Cmd("WATCH", "likes").Err
        if err != nil {
            return nil, err
        }
        reply, err := conn.Cmd("ZRANGE", "likes", 0, 9).List()
        if err != nil {
            return nil, err
        }

        // Use the MULTI command to inform Redis that we are starting a new
        // transaction.
        err = conn.Cmd("MULTI").Err
        if err != nil {
            return nil, err
        }
        for _, id := range reply {
            err := conn.Cmd("HGETALL", "product:"+id).Err
            if err != nil {
                return nil, err
            }
        }
        ereply := conn.Cmd("EXEC")
        if ereply.Err != nil {
            return nil, err
        } else if ereply.IsType(redis.Nil) {
            continue
        }

        areply, err := ereply.Array()
        if err != nil {
            return nil, err
        }
        abs := make([]*Product, 10)
        for i, reply := range areply {
            mreply, err := reply.Map()
            if err != nil {
                return nil, err
            }
            ab, err := populateEntry(mreply)
            if err != nil {
                return nil, err
            }
            abs[i] = ab
        }

        return abs, nil
    }
}



func FindTopThree() ([]*Product, error) {
    conn, err := db.Get()
    if err != nil {
        return nil, err
    }
    defer db.Put(conn)

    // Begin an infinite loop.
    for {
        // Instruct Redis to watch the likes sorted set for any changes.
        err = conn.Cmd("WATCH", "likes").Err
        if err != nil {
            return nil, err
        }
        reply, err := conn.Cmd("ZREVRANGE", "likes", 0, 2).List()
        if err != nil {
            return nil, err
        }
        err = conn.Cmd("MULTI").Err
        if err != nil {
            return nil, err
        }
        for _, id := range reply {
            err := conn.Cmd("HGETALL", "product:"+id).Err
            if err != nil {
                return nil, err
            }
        }
        ereply := conn.Cmd("EXEC")
        if ereply.Err != nil {
            return nil, err
        } else if ereply.IsType(redis.Nil) {
            continue
        }
        areply, err := ereply.Array()
        if err != nil {
            return nil, err
        }
        abs := make([]*Product, 3)
        for i, reply := range areply {
            mreply, err := reply.Map()
            if err != nil {
                return nil, err
            }
            ab, err := populateEntry(mreply)
            if err != nil {
                return nil, err
            }
            abs[i] = ab
        }

        return abs, nil
    }
}
