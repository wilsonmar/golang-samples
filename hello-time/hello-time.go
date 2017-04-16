/* hello-time.go from https://github.com/wilsonmar/golang-sample
   and in Go Playpen: https://play.golang.org/p/v3BElaaAbM
   A tutorial for first-time developers of GoLang.
*/

// All Go programs must be part of a package.
// main is the default package.
package main

// import libraries:
// fmt provides basic standard out formatting functions.
import("fmt"; "time"; "strconv") // notice semicolons are used. see https://golang.org/pkg/time/
//import("fmt"; "time"; "strings") // notice semicolons are used. see https://golang.org/pkg/time/

const (
    // For use with http://golang.org/pkg/time/#Parse
    timeLayout = "2006-01-02 15:04 MST"
)

// The one of a few reasons for a global variable:
var runtime_start  = time.Now() // See: http://golang.org/pkg/time/

// A function must match the package name to
// provide an entry point which gets called when run:
func main() {
    // Four character indent.

  	// Go requires a time zone when specifying a date, so build one:
  	milestone_tz, _ := time.LoadLocation("MST") // You can also use time.UTC constant
  	milestone_date := time.Date(2009, 11, 10, 23, 15, 4, 5, milestone_tz)
    // Format by example does not recognize use of %v
    fmt.Println("Go became an open source project on",milestone_date.Format("Mon, Jan 2, 2006 at 3:04pm"))
    // No space needs to be added in the text above.
    // fmt.Printf("\n")  // blank line

    // Code in the Go Playground http://play.golang.org/p/m5AWasfhQy
    // always returns timestamps at 1257894000 (Tue, 10 Nov 2009 23:00:00),
    // the date Go became a public open source project.

    now := time.Now()
    epoch_start := now.Unix()  // Alternate way to define variable

  	// BLAH: Duration objections do not have days due to timezones
  	// and daylight savings time (which we think is a cop out but whatever)
    // Time Addition/Subtraction returns a Duration object
  	diff := now.Sub(milestone_date)
  	diff_days := int(diff.Hours() / 24)
  	fmt.Printf("That's %d days ago\n", diff_days)


    // The Println method from fmt package writes to standard out:
    fmt.Printf("from now at %v (default format)",now)
    fmt.Printf("\n") // Blank line needs to be specified.

    // Use built-in time constants for date formatting
  	// See Time constans for more: http://golang.org/pkg/time/#pkg-constants
  	fmt.Println("RFC850:  ",now.Format(time.RFC850))
  	fmt.Println("RFC1123:  ",now.Format(time.RFC1123))

    // Double quotes within double quotes:
    fmt.Printf("That's %v seconds from midnight, January 1, 1970 (\"epoch time\")",epoch_start)
    // PROTIP: Put blank lines in a separate line for easier refactoring:
    fmt.Printf("\n")


    // Construct date of next Christmas:
    // TODO: Convert milestone_tz (from above) to string
    christmas_text := strconv.Itoa(now.Year()) + "-12-25 00:00 MST" // 0:0:0 won't dd.
    //christmas_text := "2017-12-26 12:00 MST" // + milestone_tz
  	//then, err := time.Parse("2006-01-02 15:04:05", christmas_text)
    then, err := time.Parse(timeLayout, christmas_text)
  	if err != nil {
  		 fmt.Println(">>> ",err)
       // return  // due to err
  	}
    // Calculate difference in dates from now, made positive:
    diff_now := time.Since(then) * -1
    // Go does not allow re-definition of the same variable.
    // BLAH: Duration objections do not have Days() due to timezones and daylight savings time.
    // Time Addition/Subtraction returns a Duration object
    diff_xmas := int(diff_now.Hours() / 24)
  	fmt.Println("And",diff_xmas,"days until next Midnight Christmas on", then.Format("Monday, Jan 2, 2006"))
    // See https://golang.org/pkg/fmt/

    /*
  	twodaysDiff := time.Hour * 24 * 2
  	twodays := now.Add(twodaysDiff)
  	fmt.Println("Two Days: ", twodays.Format(time.ANSIC))

    // Time Conditionals: Use Before() After() Equal() to compare dates:
  	if milestone_date.Before(now) {
  		fmt.Println("That is before Now ")
  	}
    */


    runtime_elapsed := time.Since(runtime_start)
    fmt.Printf("Execution time: %s microseconds", runtime_elapsed)
    // See https://coderwall.com/p/cp5fya/measuring-execution-time-in-go

    fmt.Printf("\n") // added so new prompt begins from left margin.
}
// Sample output: $ go run hello-time.go
/*
Go became an open source project on Tue, Nov 10, 2009 at 11:15pm
That's 2713 days ago
from now at 2017-04-16 06:56:01.397340297 -0600 MDT (default format)
RFC850:   Sunday, 16-Apr-17 06:56:01 MDT
RFC1123:   Sun, 16 Apr 2017 06:56:01 MDT
That's 1492347361 seconds from midnight, January 1, 1970 ("epoch time")
Add 252 days until next Midnight Christmas on Monday, Dec 25, 2017
Execution time: 312.909µs microseconds
*/
// EOF
