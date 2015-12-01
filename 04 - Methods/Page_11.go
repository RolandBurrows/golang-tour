// https://tour.golang.org/methods/11

package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func main() {
	reader.Validate(MyReader{})
}


// -- Results --

// prog.go:10: cannot use MyReader literal (type MyReader) as type io.Reader in argument to reader.Validate:
// 	MyReader does not implement io.Reader (missing Read method)

// Program exited.
