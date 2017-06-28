// go-server.go
// based on https://medium.com/dev-bits/understanding-the-gorilla-mux-a-sturdy-url-router-from-the-golang-7494660f4907

package main

import (
	"io"
	"net/http"
	"log"
)

// a simple web server
func SimpleServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func main() {
//	io.WriteString(w, "Starting go-server.go\n")
	http.HandleFunc("/hello", SimpleServer)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

// BLAH: I'm getting a "404" in the web page.
