// https://tour.golang.org/methods/15

package main

import (
        "fmt"
        "image"
)

func main() {
        m := image.NewRGBA(image.Rect(0, 0, 100, 100))
        fmt.Println(m.Bounds())
        fmt.Println(m.At(0, 0).RGBA())
        fmt.Println(m.At(50, 50).RGBA())
        fmt.Println(m.At(100, 100).RGBA())
}


// -- Results --

// (0,0)-(100,100)      // Bounds are as expected, and is a square
// 0 0 0 0              // Image is perfectly white
// 0 0 0 0
// 0 0 0 0

// Program exited.
