// https://tour.golang.org/moretypes/13

package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("index (%d) has contents (%d)\n", i, v)
	}
}


// -- Results --

// index (0) has contents (1)
// index (1) has contents (2)
// index (2) has contents (4)
// index (3) has contents (8)
// index (4) has contents (16)
// index (5) has contents (32)
// index (6) has contents (64)
// index (7) has contents (128)
