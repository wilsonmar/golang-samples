/* latlong.go
   Manages latitudes and longitudes in a map containing a key-value structure.
   Uses Go map functionality.
   Uses float64
   Uses function to do repeated lookups
   Adapted from https://tour.golang.org/moretypes/19
*/

package main

import "fmt"

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
    // TODO: Read from a file.

	// Lookup by key:
    // key := "Bell Labs"
    key := "Taj Mahal"
	fmt.Println(key,"is at Longitude,Latitude",m[key])
	//LatLong(key)
	//LatLong("Taj Mahal")
}

// See https://github.com/mkaz/working-with-go/blob/master/04-functions.go
//func LatLong(key string) string {
//	return m[key]
// }

/* Response to go run latlong.go is the vertex value that corresponds to the key specified in the println:
Bell Labs is at Longitude,Latitude {40.68433 -74.39967}
NOTE: No commas between the two values.
*/
