// https://tour.golang.org/moretypes/2

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}


// -- Results --

// {1 2}

// Program exited.
