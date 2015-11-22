// https://tour.golang.org/methods/3

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(6)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}


// -- Results --

// Before scaling: &{X:3 Y:4}, Abs: 5
// After scaling: &{X:18 Y:24}, Abs: 30

// Program exited.
