// https://tour.golang.org/flowcontrol/10

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

// Note: Time in the Go playground always appears to start at 2009-11-10 23:00:00 UTC, a value whose
// significance is left as an exercise for the reader.

// 2009-11-10 23:00:00 UTC was a Tuesday, so, "Too far away."


// -- Results --

// When's Saturday?
// Too far away.

// Program exited.
