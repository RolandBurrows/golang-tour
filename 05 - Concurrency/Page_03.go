// https://tour.golang.org/concurrency/3

// Buffered Channels

package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}


// -- Results --

// 1
// 2
// 3
// fatal error: all goroutines are asleep - deadlock!

// goroutine 1 [chan receive]:
// main.main()
// 	/tmp/sandbox103008279/main.go:13 +0x440

