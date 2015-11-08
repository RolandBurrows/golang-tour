// https://tour.golang.org/flowcontrol/6

package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	// Returns x^n, unless that's higher than lim, in which case it return lim
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}


// -- Results --

// 9 20

// Program exited.
