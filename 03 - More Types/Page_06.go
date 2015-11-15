// https://tour.golang.org/moretypes/6

package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
}


// -- Results --

// Hello World
// [Hello World]

// Program exited.
