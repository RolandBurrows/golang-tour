// https://tour.golang.org/basics/1

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}


// -- Results --

// My favorite number is 1

// Program exited.
