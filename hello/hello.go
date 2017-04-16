/* hello.go from https://github.com/wilsonmar/golang-sample
   Sample output: "hello, world at 2017-04-15 17:32:52.470348041 -0600 MDT
*/
package main
import("fmt"; "time") // notice semicolons are used. see https://golang.org/pkg/time/
var now = time.Now()
func main() {
    fmt.Printf("hello, world at %v.",now)
    fmt.Printf("\n")
}
