// https://tour.golang.org/methods/16

package main

import "golang.org/x/tour/pic"

type Image struct{}

func main() {
	m := Image{}
	pic.ShowImage(m)
}


// -- Results --

// prog.go:9: cannot use m (type Image) as type image.Image in argument to pic.ShowImage:
// 	Image does not implement image.Image (missing At method)

// Program exited.
