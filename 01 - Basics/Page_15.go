// https://tour.golang.org/basics/15

package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}


// -- Results --

// Hello 世界
// Happy 3.14 Day
// Go rules? true

// Program exited


// Excellent extra reading: https://blog.golang.org/constants
