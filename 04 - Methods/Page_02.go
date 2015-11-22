// https://tour.golang.org/methods/2

package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	g := MyFloat(math.Sqrt2)
	fmt.Println(f.Abs(), g.Abs())
}


// -- Results --

// 1.4142135623730951 1.4142135623730951

// Program exited.
