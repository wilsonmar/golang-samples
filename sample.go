// sample.go 
// USAGE: go run sample.go add -a=42 -b=23
//        go run sample.go mul -a=25 -b=4

package main

import( "fmt"; "time"; "os"; "errors"; "flag"; "math"; "strings")

var now = time.Now()  // wall clock

func circleArea(radius float32) (float32, error) {
    if radius < 0 {
        return 0, errors.New("ERROR: In call of circleArea: radius should be a positive value.")
    } else {
        return math.Pi * radius * radius, nil
    }
}

func main() {
    start_time := time.Now()   // monotonic clock
    args := os.Args[0]
    f := strings.Fields(args)
    for _, v := range f {
       fmt.Println(v)  // TODO: output in one
    }
    fmt.Printf("starting at %v \n", now)  // https://golang.org/pkg/time/#Time.Format

       argLength := len(os.Args[1:])
    if argLength > 0 {
        for i, a := range os.Args[1:] {
        	// TODO: Use defer to concatenate string parts?
            fmt.Printf("Arg %d is %s\n", i+1, a) 
        }
    }

    // Adapted from https://golangdocs.com/flag-package-golang
    // TODO: Add -verbosity
    addcmd := flag.NewFlagSet("add", flag.ExitOnError)
    a_add := addcmd.Int("a", 0, "The value of a")
    b_add := addcmd.Int("b", 0, "The value of b")
 
    mulcmd := flag.NewFlagSet("mul", flag.ExitOnError)
    a_mul := mulcmd.Int("a", 0, "The value of a")
    b_mul := mulcmd.Int("b", 0, "The value of b")
 
    switch os.Args[1] {
    case "add":
        addcmd.Parse(os.Args[2:])
        fmt.Println(*a_add + *b_add)
    case "mul":
        mulcmd.Parse(os.Args[2:])
        fmt.Println(*(a_mul) * (*b_mul))
    default:
        fmt.Println("expected add or mul command")
        os.Exit(1)
    }

    // Known error (-3 radius input):
    if area2, err := circleArea(-3); err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(area2)
    }

    t := time.Now()
    elapsed := t.Sub(start_time)
    fmt.Printf("Elapsed: %v\n", elapsed)
    fmt.Printf("Seconds: %f\n", elapsed.Seconds())
    fmt.Printf("Minutes: %f\n", elapsed.Minutes())

    defer func(msg string) {
        if r := recover(); r != nil {
            fmt.Println("recovered")	
        }
        fmt.Println(msg)
    }("Done.")

//    panic("Die!")  // to test defer function to recover
}
