/* rand.go at https://github.com/wilsonmar/golang-samples/rand
   Option to output a flat range as an integer or a floating point number.

   Adapted from https://www.golang-book.com/books/intro/13#section8 for command line arguments
   And https://www.golang-book.com/books/intro/13#section6 which doesn't have a random seed.
   And https://stackoverflow.com/questions/12321133/golang-random-number-generator-how-to-seed-properly
*/

package main

import (
  "fmt"
  "time" // for random seed generation
  "flag" // to obtain arguments from command line
  "math/rand"  // http://golang.org/pkg/math/rand/
)

func main() {

  // Define flags (command line arguments):
  max_p   := flag.Int("max", 100, "the max value") // Define default 100 if -max=? is not specified in the run command.
  repeats := flag.Int("repeat", 1, "the number of outputs")
  distribution := flag.Int("distribution", 1, "the type of distribution")

  // So each run returns a different number, get a (deterministic) random seed based on a precise time :
  rand.Seed(time.Now().UTC().UnixNano()) // value is a epoch date/time stamp like 1492421032295856507

  // Parse flag (command line arguments) from the go run command:
  flag.Parse()

  // BLAH: Additional non-flag arguments can be retrieved with flag.Args() which returns a []string:
  // fmt.Println(flag.Args()) // returns no blank line if no Args were defined in the run command.

  // TODO: Send to a file.

  // TODO: Loop the number of repeats specified:
  // PROTIP: Reference argument flag integers as pointers:
  fmt.Println("Repeats flag:", *repeats)
  fmt.Println("Distribution flag:", *distribution)

  for i := 0; i < *repeats; i++ {
      // BLAH: ./rand.go:33: invalid operation: i < repeats (mismatched types int and *int)

      // There are several types of random numbers:

      // distribution == 1 (default)

         // Intn (integer number) gets the next random integer. Then Print the contents of pointer to the flag variable:
         fmt.Println("Of", *max_p, ":", rand.Intn(*max_p))

         // Flat chance:
      	 fmt.Println("Random Int:", rand.Int())

      	 fmt.Println("Random Float:", rand.Float32())

      // distribution == 2
       	 // A not so random function in the math/rand package
       	 // The NormFloat64 method returns a number based on Normal distrib
        	// mean=0, stddev=1 -- if called enough and plotted will give bell curve
        fmt.Println("Normalized:", rand.NormFloat64())

        // fmt.Println("Random Permutation of N ints:", rand.Perm(8))
  }
}
/* Respoonse to go run sha1.go -max=50
31
*/
