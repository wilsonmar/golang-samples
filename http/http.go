/* http.go
   Adapted from https://talks.godoc.org/github.com/davecheney/introduction-to-go/introduction-to-go.slide#10
   go run http.go
   Then open a new browser to http://localhost:7777/
*/

package main

import (
    "html/template"
    "log"
    "net/http"
)

const templ = `<html><head><title>Hello</title></head><body>
Hello {{ .RemoteAddr }}
You sent me a {{ .Method }} request for {{ .URL }}
</body></html>`

func HelloServer(w http.ResponseWriter, req *http.Request) {
    log.Println(req.URL)
    t := template.Must(template.New("html").Parse(templ))
    t.Execute(w, req)
}

func main() {
    log.Println("please connect to http://localhost:7777/")
    http.HandleFunc("/", HelloServer)
    log.Fatal(http.ListenAndServe(":7777", nil))
}