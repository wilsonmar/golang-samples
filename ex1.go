// Based on https://medium.com/dev-bits/understanding-the-gorilla-mux-a-sturdy-url-router-from-the-golang-7494660f4907

package main
import (
        "github.com/gorilla/mux"
        "log"
        "net/http"
        "time"
        "fmt"
      )
func ArticleHandler(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Category is: %v\n", vars["category"])
    fmt.Fprintf(w, "ID is: %v\n", vars["id"])
}
func main(){
   // Create a new router
   r := mux.NewRouter()
   // Attach an elegant path with handler
   r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)
   srv := &http.Server{
        Handler:      r,
        Addr:         "127.0.0.1:8080",
        // Good practice: enforce timeouts for servers you create!
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }
   log.Fatal(srv.ListenAndServe())
}
