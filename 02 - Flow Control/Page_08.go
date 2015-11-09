// https://tour.golang.org/flowcontrol/8

package main

import (
	"fmt"
)

var delta float64 = 0.0000001

func Compare(a, b float64) bool {
	if ((a - b) < delta && (b - a) < delta) {
		return true
	}
	return false
}

func Newt(x, z float64) float64 {
	lim := 1
	zee := 1.0
	for lim < 100 {
		zee = (z-(((z*z)-x)/(2*z)))
		if Compare(zee, z) {
			return zee
		}
		z = zee
		lim++
	}
	return 0.0
}

func Sqrt(x float64) float64 {
	var init float64 = 1
	return Newt(x, init)
}

func main() {
	fmt.Println(Sqrt(10000))
}


// -- Results --

// 1 -> 1
// 2 -> 1.4142135623730951
// 4 -> 2.000000000000002		// TODO: Why is this not exactly 2?
// 25 -> 5
// 10000 -> 100
