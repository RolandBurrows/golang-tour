// https://tour.golang.org/moretypes/15

package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	proc := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		proc[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			proc[i][j] = uint8(i*j)
		}
	}
	return proc
}

func main() {
	pic.Show(Pic)
}


// -- Results --

// Image of i*j in same folder
