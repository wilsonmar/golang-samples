// api.go
// Adapted from https://thenewstack.io/make-a-restful-json-api-go/

package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)
// Before running this - go get github.com/gorilla/mux

func main() {

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    log.Fatal(http.ListenAndServe(":8080", router))
}
func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
    todos := Todos{
        Todo{Name: "Write presentation"},
        Todo{Name: "Host meetup"},
    }

    json.NewEncoder(w).Encode(todos)
}
