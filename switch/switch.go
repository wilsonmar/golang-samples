/* switch/switch.go in https://github.com/wilsonmar/golang-samples
   Adapted from https://talks.godoc.org/github.com/davecheney/introduction-to-go/introduction-to-go.slide#89
*/

package main

import ( // packages:
    "fmt"
)

func main() {
    primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23}
    primes = append(primes, 29)

    fmt.Println("Example of a switch referencing a slice:")
    //fmt.Println("\n")
    for i, p := range primes {
        i++
        switch i {
        case 1:
            fmt.Println("The", i, "st prime is", p)
        case 2:
            fmt.Println("The", i, "nd prime is", p)
        case 3:
            fmt.Println("The", i, "rd prime is", p)
        default:
            fmt.Println("The", i, "th prime is", p)
        }
    }
}