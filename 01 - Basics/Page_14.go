// https://tour.golang.org/basics/14

package main

import "fmt"

func main() {
	v := 42 // change me!
	fmt.Printf("v is of type %T\n", v)
}


// -- Results --

// v is of type int

// Program exited.


// For v := 3.14
// v is of type float64

// For v := 0.5i
// v is of type complex128

// For v := "dog"
// v is of type string

// For v := true
// v is of type bool