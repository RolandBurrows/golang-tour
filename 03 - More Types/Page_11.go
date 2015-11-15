// https://tour.golang.org/moretypes/11

package main

import "fmt"

func main() {
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
}


// -- Results --

// [] 0 0
// nil!

// Program exited.
