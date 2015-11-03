// https://tour.golang.org/basics/7

package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}


// -- Results --

// First guess:
// x = (17 * 4) / 9 = 7.555bar
// y = (17) - 7.555bar = 9.444bar
// print: "7.555, 9.444"


// Run Results:

// 7 10

// Program exited.


// Derp. They're ints, not floats.
