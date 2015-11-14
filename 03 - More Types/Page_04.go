// https://tour.golang.org/moretypes/4

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}


// -- Results --

// {1000000000 2}

// Program exited.
