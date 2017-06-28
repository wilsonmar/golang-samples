/* slice.go
// Different ways to define and update an array.
   Adapted from https://talks.godoc.org/github.com/davecheney/introduction-to-go/introduction-to-go.slide#69
*/

package main

import "fmt"

func main() {
    // composite literal static:
    x := []int{7, 9, 42}
  	fmt.Println(x)
    // Output: [7 9 42]

    // slice literal: from https://play.golang.org/p/nGekFavi0S
    y := make([]int, 0, 100)
  	y = append(y, 8)
  	y = append(y, 12)
  	y = append(y, 43)
  	fmt.Println(y)
    // Output: [8 12 43]

    for i, v := range y {  // Range parsed rune by rune for utf-8
  		fmt.Println("Range:", i, "-", v)
      // Output: Range: 0 - 8
      // Range: 1 - 12
      // Range: 2 - 43
  	}
    // composite literal in a for loop:
    z := make([]int, 3)
    for i := 0; i < len(z); i++ {
        z[i] = i
        fmt.Println(z[i])
    }
}
