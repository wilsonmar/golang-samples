// hello.go
package main
import "fmt"
import "time"
var now = time.Now()
func main() {
    fmt.Printf("hello, world at %v",now)
}
