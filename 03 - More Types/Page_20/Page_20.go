// https://tour.golang.org/moretypes/20

package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	fieldz := strings.Fields(s)
	numz := len(fieldz)
	mapz := make(map[string]int)
	for y := 0; y < numz; y++ {
		(mapz[fieldz[y]])++
	}
	return mapz
}

func main() {
	wc.Test(WordCount)
}


// -- Results --

// PASS
//  f("I am learning Go!") = 
//   map[string]int{"I":1, "am":1, "learning":1, "Go!":1}
// PASS
//  f("The quick brown fox jumped over the lazy dog.") = 
//   map[string]int{"The":1, "fox":1, "lazy":1, "over":1, "the":1, "dog.":1, "quick":1, "brown":1, "jumped":1}
// PASS
//  f("I ate a donut. Then I ate another donut.") = 
//   map[string]int{"donut.":2, "Then":1, "another":1, "I":2, "ate":2, "a":1}
// PASS
//  f("A man a plan a canal panama.") = 
//   map[string]int{"a":2, "plan":1, "canal":1, "panama.":1, "A":1, "man":1}

