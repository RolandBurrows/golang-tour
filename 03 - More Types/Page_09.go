// https://tour.golang.org/moretypes/9

package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s ==", s)					// [2, 3, 5, 7, 11, 13] ?
	fmt.Println("s[1:4] ==", s[1:4])		// [3, 5, 7, 11] ?

	// missing low index implies 0
	fmt.Println("s[:3] ==", s[:3])			// [2, 3, 5, 7] ?

	// missing high index implies len(s)
	fmt.Println("s[4:] ==", s[4:])			// [11, 13] ?
}


// -- Results --

// s == [2 3 5 7 11 13]
// s[1:4] == [3 5 7]
// s[:3] == [2 3 5]
// s[4:] == [11 13]

// Program exited.


// From Notes:
// "from lo through hi-1, inclusive" - ah, right.
