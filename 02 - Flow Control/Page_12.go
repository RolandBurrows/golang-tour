// https://tour.golang.org/flowcontrol/12

package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}


// -- Results --

// hello
// world

// Program exited.
