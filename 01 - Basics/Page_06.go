// https://tour.golang.org/basics/6

package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}


// -- Results --

// world hello

// Program exited.