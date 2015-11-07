// https://tour.golang.org/flowcontrol/2

package main

import "fmt"

func main() {
	sum := 1
	for ; sum < 1000; {
		fmt.Println(sum)
		sum += sum
	}
	fmt.Println(sum)
}


// -- Results --

// 1
// 2
// 4
// 8
// 16
// 32
// 64
// 128
// 256
// 512
// 1024
