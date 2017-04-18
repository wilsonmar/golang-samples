/* latlong.go
   Manages latitudes and longitudes in a map data type containing a key-value structure.
   Uses Go map functionality.
	 Calculates a absolute value in a rough comparison of two floats using Newton's method with a tolerance.
   Uses function to do repeated lookups
   Adapted from https://tour.golang.org/moretypes/19
*/

package main

import(
	"fmt"
	"math"  // http://golang.org/pkg/math/
)

type Vertex struct {
	Lat, Long float64 // Degrees and decimal minutes (DMM)
}

// Globals:
var m map[string]Vertex

func main() {

	m = make(map[string]Vertex)
	m["Bell Labs"]   = Vertex{ 40.68433, -74.39967 }  // North and West
	m["Machu Pichu"] = Vertex{ -13.1631, -72.5450 } // 13.1631° S, 72.5450° W (negative/left of UTC)
	m["Taj Mahal"]   = Vertex{  27.1750, 78.0422 } // North and East of UTC
    // TODO: Add more values.
    // TODO: Read vertexes from a file.

	// Lookup by key:
  key := "Taj Mahal"
	fmt.Println(key,"is at Longitude,Latitude",m[key])
	//LatLong("Bell Labs")
	//LatLong("Taj Mahal")
	//LatLong("Machu Pichu")



  // Determine if two floats are roughly equal within a tolerance:
	const TOLERANCE = 0.000001
	result := 27.1750
	temp   := 27.17501
	for { // Using Newton's method - see https://coderwall.com/p/pzhz9q/comparing-floating-point-integers-in-golang
	  if diff := math.Abs(result - temp); diff < TOLERANCE { // consider equal and break
			fmt.Println("Within tolerance of",TOLERANCE)
	    break
	  } else {
			// Loop here to do something:
			fmt.Println(result,"and",temp,"not the same.")
			break
	  }
	}
}

// See https://github.com/mkaz/working-with-go/blob/master/04-functions.go
//func LatLong(key string) string {
//	return m[key]
// }

/* Response to go run latlong.go is the vertex value that corresponds to the key specified in the println:
Bell Labs is at Longitude,Latitude {40.68433 -74.39967}
NOTE: No commas between the two values.
*/
