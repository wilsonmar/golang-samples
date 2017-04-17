/* tcp-server.go in https://github.com/wilsonmar/golang-samples
   A headless HTTP TCP protocol web server
   See https://coderwall.com/p/wohavg/creating-a-simple-tcp-server-in-go
   TODO: http://loige.co/simple-echo-server-written-in-go-dockerized/_
*/
package main

import (
    "fmt"
    "net"
    "os"
    // "log" 
)

const (
    CONN_HOST = "localhost"
    CONN_PORT = "3333"
    CONN_TYPE = "tcp"
)

func main() {
    // Listen for incoming connections.
    l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }
    // Close the listener when the application closes.
    defer l.Close()
    fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
    for {
        // Listen for an incoming connection.
        conn, err := l.Accept()
        if err != nil {
            fmt.Println("Error accepting: ", err.Error())
            os.Exit(1)
        }
        // Handle connections in a new goroutine.
        go handleRequest(conn)
    }
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
  // Make a buffer to hold incoming data.
  buf := make([]byte, 1024)
  // Read the incoming connection into the buffer.
  reqLen, err := conn.Read(buf)
  if err != nil {
    fmt.Println("Error reading:", err.Error())
  }
  // Write out the number of bytes received:
  fmt.Println(reqLen)
  // Send a response back to person contacting us.
  conn.Write([]byte("Message received."))
  // Close the connection when you're done with it.
  conn.Close()
}

/* Initiate this service: go run tcp.go
   Open another Terminal shell window To test your server.
   Send some raw data to that port:
   echo -n "test" | nc localhost 3333
   The response is a count of the number of characters.
*/
