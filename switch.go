/* switch/switch.go in https://github.com/wilsonmar/golang-samples
   Adapted from https://talks.godoc.org/github.com/davecheney/introduction-to-go/introduction-to-go.slide#89
*/

package main

import ( // packages:
    "fmt"
    "time" // to avoid undefined: time in time.Now
    "runtime"
)


func main() {

  // Switch cases evaluate cases from top to bottom, stopping when a case succeeds.

  t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}


  fmt.Print("Go running on ")
  // Obtain operating system (os) from the runtime module:
  switch os := runtime.GOOS; os {
    case "darwin":
      fmt.Println("MacOS.")
      // Do something on a Mac.
    case "linux":
      fmt.Println("Linux.")
      // TODO: What flavor? Debian, CoreOS, Red Hat, etc.
      // Do some Linux.
    case "windows":
      fmt.Println("Windows.") // TODO: Verify this.
      // Do some Microsoft.
    default:
      // freebsd, openbsd, plan9, windows...
      fmt.Printf("%s.", os)
  }

      // Define a bunch of prime numbers:
    primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23}
    // Add another item:
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
// A case body breaks automatically, unless it ends with a fallthrough statement.
