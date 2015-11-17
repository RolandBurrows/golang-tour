// https://tour.golang.org/moretypes/15

package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// Failure point
}

func main() {
	pic.Show(Pic)
}


// -- Results --

// prog.go:6: missing return at end of function

// Program exited.
