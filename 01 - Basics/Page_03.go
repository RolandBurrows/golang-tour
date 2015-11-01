// https://tour.golang.org/basics/3

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.pi)
}


// -- Results --

// prog.go:9: cannot refer to unexported name math.pi
// prog.go:9: undefined: math.pi

// Program exited.


package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
}


// -- Results --

// 3.141592653589793

// Program exited.