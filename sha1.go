/* sha1.go
   Adapted from https://www.golang-book.com/books/intro/13#section6
   And https://www.golang-book.com/books/intro/13#section8 for command line arguments
*/

package main

import (
  "fmt"
  "flag"
  "crypto/sha1"
)

func main() {
  // Define flags (command line arguments):
  clear_text := flag.String("text", "text", "the text to be hashed") // "some text" is the default if flag is not added with run.
  fmt.Println(clear_text)
  // Parse through command line arguments from the go run command:
  flag.Parse()

  h := sha1.New()
  h.Write([]byte("test"))
  bs := h.Sum([]byte{})
  fmt.Println(bs)

}
/* Respoonse to go run sha1.go -text="hello"
*/
