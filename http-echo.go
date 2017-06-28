/* http-echo.go
   Adapted from https://mkaz.tech/geek/testing-clients-to-an-http-api-in-go/
   This sets up a test server to just echo back whatever parameter gets passed in, to confirm what was passed in.
   This makes use of Go's standard library net/http/httptest which creates a test HTTP server,
   similar to Go’s normal HTTP server. The test server creates a server that listens locally on a random port.
   https://github.com/Automattic/go/tree/master/jaguar
*/
package main

    func TestGetParams(t *testing.T) {

        // echoHandler, passes back form parameter p
        echoHandler := func( w http.ResponseWriter, r *http.Request) {
            fmt.Fprint(w, r.FormValue("p"))
        }

        // create test server with handler
        ts := httptest.NewServer(http.HandlerFunc(echoHandler))
        defer ts.Close()

        // call library I want to test, using test server ts
        f := fetcher.NewFetcher()
        f.Params.Add("p", "hello")
        result, err := f.Fetch(ts.URL,"GET")
        if err != nil {
            t.Errorf("Error: %v", err)
        }

        // confirm result
        if result != "hello" {
            t.Errorf("Unexpected result: %v", result)
        }
    }
