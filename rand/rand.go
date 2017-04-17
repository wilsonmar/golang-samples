/* rand.go at https://github.com/wilsonmar/golang-samples/rand
   Adapted from https://www.golang-book.com/books/intro/13#section6
   And https://www.golang-book.com/books/intro/13#section8 for command line arguments
*/

package main

import (
  "fmt"
  "flag"
  "math/rand"
)

func main() {

  // Define flags (command line arguments):
  max_p := flag.Int("max", 100, "the max value") // Define default 100 if -max=? is not specified in the run command.

  // TODO: How to change the random seed?

  // BLAH: There is no random seed with this program. All attempts with the same max input returns the same number.

  // Parse flag (command line arguments) from the go run command:
  flag.Parse()

  // BLAH: Additional non-flag arguments can be retrieved with flag.Args() which returns a []string:
  // fmt.Println(flag.Args()) // returns no blank line if no Args were defined in the run command.

  // Print the contents of pointer to the flag variable:
  fmt.Println(rand.Intn(*max_p))

}
/* Respoonse to go run sha1.go -max=50
31
*/