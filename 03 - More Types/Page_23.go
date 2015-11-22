// https://tour.golang.org/moretypes/23

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
}										// Fails here

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}


// -- Results --

// prog.go:8: missing return at end of function

// Program exited.
