/* rand.go at https://github.com/wilsonmar/golang-samples/rand
   Adapted from https://www.golang-book.com/books/intro/13#section8 for command line arguments
   And https://www.golang-book.com/books/intro/13#section6 which doesn't have a randome seed.
   And https://stackoverflow.com/questions/12321133/golang-random-number-generator-how-to-seed-properly
*/

package main

import (
  "fmt"
  "time" // for random seed generation
  "flag" // to obtain arguments from command line
  "math/rand"
)

func main() {

  // Define flags (command line arguments):
  max_p := flag.Int("max", 100, "the max value") // Define default 100 if -max=? is not specified in the run command.
  //repeats := flag.Int("repeat", 1, "the number of outputs")

  // So each run returns a different number, get a random seed based on the time :
  rand.Seed(time.Now().UTC().UnixNano()) // value is a epoch date/time stamp like 1492421032295856507

  // Parse flag (command line arguments) from the go run command:
  flag.Parse()

  // BLAH: Additional non-flag arguments can be retrieved with flag.Args() which returns a []string:
  // fmt.Println(flag.Args()) // returns no blank line if no Args were defined in the run command.

  // TODO: Send to a file.

  // TODO: Loop the number of repeats specified:
  //for i := 0; i < repeats; i++ {
      // BLAH: ./rand.go:33: invalid operation: i < repeats (mismatched types int and *int)
      // Intn gets the next random integer. Then Print the contents of pointer to the flag variable:
      fmt.Println(rand.Intn(*max_p))
  //}
}
/* Respoonse to go run sha1.go -max=50
31
*/
