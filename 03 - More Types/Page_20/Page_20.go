// https://tour.golang.org/moretypes/20

package main

import (
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	return map[string]int{"x": 1}
}

func main() {
	wc.Test(WordCount)
}


// -- Results --

// FAIL
//  f("I am learning Go!") =
//   map[string]int{"x":1}
//  want:
//   map[string]int{"Go!":1, "I":1, "am":1, "learning":1}
// Program exited.
