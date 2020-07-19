/* USAGE: go run challenge.go 1st 2nd 3rd 4th
// Golang program is coded with features commonly needed:
// * command-line arguments (-help, -verbose)
// * Logging
*/

package main 
  
import ( 
    "fmt"
    "os"
) 
  
func main() { 
  
    // The first argument 
    // is always program name 
    myProgramName := os.Args[0] 
  
    // this will take 4 
    // command line arguments 
    cmdArgs := os.Args[4] 
  
    // getting the arguments 
    // with normal indexing 
    gettingArgs := os.Args[2] 
  
    toGetAllArgs := os.Args[1:] 
  
    // it will display 
    // the program name 
    fmt.Println(myProgramName) 
      
    fmt.Println(cmdArgs) 
      
    fmt.Println(gettingArgs) 
      
    fmt.Println(toGetAllArgs) 
} 