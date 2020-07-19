// This is the minimal Go program, to show Go compiler is installed correctly.
package main
import( "fmt"; "time")
var now = time.Now()
func main() {
    fmt.Printf("hello, world at %v",now)
}