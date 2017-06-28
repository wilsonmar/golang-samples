/* struct.go
   A structure of multiple fields 
*/

package main

import (
  "fmt"
)

func main() {

    p1 := struct {
      fname string
      lname string
    }{
      "James",
      "Bond",  // notice the comma even on the last line. Thank you.
    }
   fmt.Println(p1)  // output: {James Bond}
}
