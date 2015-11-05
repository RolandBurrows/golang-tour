// https://tour.golang.org/flowcontrol/1

package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		fmt.Println("i =", i)
		sum += i
		fmt.Println("sum =", sum)
	}
	fmt.Println(sum)
}


// -- Results --

// i = 0
// sum = 0
// i = 1
// sum = 1
// i = 2
// sum = 3
// i = 3
// sum = 6
// i = 4
// sum = 10
// i = 5
// sum = 15
// i = 6
// sum = 21
// i = 7
// sum = 28
// i = 8
// sum = 36
// i = 9
// sum = 45
// 45
