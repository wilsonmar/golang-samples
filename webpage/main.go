// This generates a HTML file, then serves that page.
// Based on

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template
func init() { // Generate HTML files from template files:
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

// func index(res http.ResponseWriter, req *http.Request) {
type handler1 int
func (m handler1) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	err := tpl.ExecuteTemplate(res, "templates/index.gohtml", nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}

func main() {

	// Handle File:
	// http.HandleFunc("/", index)
	// http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))

  var h handler

	fmt.Println("Use internet browser to view http://localhost:8080.")
	fmt.Println("Press control+C to stop server.")
	http.ListenAndServe(":8080", h)
}
