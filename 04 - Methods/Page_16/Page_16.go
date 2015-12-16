// https://tour.golang.org/methods/16

package main

import (
        "image"
        "image/color"
        "code.google.com/p/go-tour/pic"
)

type Image struct {
        Height, Width int
}

func (m Image) Bounds() image.Rectangle {
        return image.Rect(0, 0, m.Height, m.Width)
}

func (m Image) ColorModel() color.Model {
        return color.RGBAModel
}

func (m Image) At(x, y int) color.Color {
        c := uint8(x ^ y)
        return color.RGBA{c, c, 255, 255}
}

func main() {
        m := Image{256, 256}
        pic.ShowImage(m)
}


// -- Results --

// See image file in same dir
