// https://tour.golang.org/flowcontrol/11

package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

// Note: Time in the Go playground always appears to start at 2009-11-10 23:00:00 UTC, a value whose
// significance is left as an exercise for the reader.

// 23:00:00 UTC == "Good evening."


// -- Results --

// Good evening.

// Program exited.
