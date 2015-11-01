// https://tour.golang.org/basics/4

package main

import "fmt"

func add(x int, y int) int {   // add(x, y int) is also valid, When two or more consecutive
	return x + y			   // named function parameters share a type
}

func main() {
	fmt.Println(add(42, 13))
}


// -- Results --

// 55

// Program exited.
