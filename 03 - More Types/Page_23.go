// https://tour.golang.org/moretypes/23

package main

import "fmt"

func fibonacci() func() int {
	a := 0
	b := 1
	return func() int {
		c := a
		d := b
		a = d
		b = c + d
		return a
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}


// -- Results --

// 1
// 1
// 2
// 3
// 5
// 8
// 13
// 21
// 34
// 55

// Program exited.
