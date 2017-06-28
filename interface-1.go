// interface.go
// From https://www.youtube.com/watch?v=rKnDgT73v8s&t=30m40s
// by "Commander" Rob Pike of Google 30 October 2009
// CAUTION: This program is currently broken!

package main


func main() {


type Magnitude interface { // implemented de-facto.
  Abs() float; // returns a float, among other things
}
var m Magnitured;
m = x;  // x is type *Point, has method Abs()
mag := m.Abs();
type Point3 struct { X, Y, Z float }
func (p *Point3) Abs() float {
  return math.Sqrt(p.X*.pX + p.Y*p.Y + p.Z*p.Z);
}
m = &Point3{ 3, 4, 5 };
mag += m.Abs();
// On a Mac: press command + control + space at the same time:
type Polar struct { R, θ float }; // Greek upper-case Theta Unicode 03B8
func (p Polar) Abs() float { return p.R};
m = Polar { 2.0 PI/2 };
mag += mAbs();
}
