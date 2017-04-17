/* slice.go
   Adapted from https://talks.godoc.org/github.com/davecheney/introduction-to-go/introduction-to-go.slide#69
*/

package main

import "fmt"

func main() {
    x := make([]int, 5)
    for i := 0; i < len(x); i++ {
        fmt.Println(x[i])
    }
}