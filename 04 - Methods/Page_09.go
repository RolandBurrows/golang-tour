// https://tour.golang.org/methods/9

package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Error: can't Sqrt() a negative number: (%g)", float64(e))
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNegativeSqrt(f)
	} else if f > 0 {
        return math.Sqrt(f), nil
    } else {
		return 0, nil
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
	fmt.Println(Sqrt(0))
}


// -- Results --

// 1.4142135623730951 <nil>
// 0 Error: can't Sqrt() a negative number: (-2)
// 0 <nil>

// Program exited.
