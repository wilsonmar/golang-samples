/* utf8.go in https://github.com/wilsonmar/golang-samples
   Provides examples of counting and comparing Asian characters encoded in UTF-8.
   Provides an illustration of fmt.Print() vs. fmt.Println() with escapted double-quotes.
   Provides an example of using range.
   Adapted from https://coderwall.com/p/k7zvyg/dealing-with-unicode-in-go
*/

package main

import (
    "fmt"
    "unicode/utf8"
)

func CompareChars(word string) {
    s := []byte(word)
    fmt.Println("CompareChars of each character (rune) with the character on its left")
    fmt.Print("in: \"",word,"\" : ")
    for utf8.RuneCount(s) > 1 {
        r, size := utf8.DecodeRune(s)
        s = s[size:]
        nextR, size := utf8.DecodeRune(s)
        fmt.Print(r == nextR, ",")
    }
    fmt.Print("\n")
}

func CountChars(in_word string){
  // PROTIP: Output a string as individual runes:
  fmt.Print("CountChars: \"")
  for _, c := range in_word {
    fmt.Print(string(c))
  }
  fmt.Print("\"")

  // PROTIP: Count number of letters in UTF-8 using a function from the utf8 built-in library:
  fmt.Print(" has ",utf8.RuneCountInString(in_word)," UTF-8 letters.")
  fmt.Print("\n")
  // PROTIP: Extra spaces are added by fmt.Println()
}

func main(){

  hello := "Hello, 世界" // 世界 = "world" in Chinese
  CountChars(hello)

  CompareChars("你好好好") // 你好好好 = How are you? Transliterad: You good good good?

}
/* Sample results from go run utf8.go
CountChars: "Hello, 世界" has 9 UTF-8 letters.
CompareChars of each character (rune) with the character on its left
in: "你好好好" : false,true,true,
*/
