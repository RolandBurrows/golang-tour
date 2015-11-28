// https://tour.golang.org/methods/9

package main

import (
	"fmt"
)

func Sqrt(x float64) (float64, error) {
	return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}


// -- Results --

// 0 <nil>
// 0 <nil>

// Program exited.
