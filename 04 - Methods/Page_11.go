// https://tour.golang.org/methods/11

package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (MyReader) Read(output []byte) (int, error) {
    length := len(output)
    for i := 0; i < length; i++ {
        output[i] = 'A'
    }
    return length, nil
}

func main() {
	reader.Validate(MyReader{})
}


// -- Results --

// OK!

// Program exited.
