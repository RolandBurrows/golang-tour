// https://tour.golang.org/basics/2

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems.", math.Nextafter(2, 3))
}



// -- Results --

// Now you have 2.0000000000000004 problems.
// Program exited.
