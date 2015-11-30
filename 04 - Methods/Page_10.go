// https://tour.golang.org/methods/10

package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader! I love you!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}


// -- Results --

// b[:n] = "Hello, R"
// n = 8 err = <nil> b = [101 97 100 101 114 33 32 73]
// b[:n] = "eader! I"
// n = 8 err = <nil> b = [32 108 111 118 101 32 121 111]
// b[:n] = " love yo"
// n = 2 err = <nil> b = [117 33 111 118 101 32 121 111]
// b[:n] = "u!"
// n = 0 err = EOF b = [117 33 111 118 101 32 121 111]
// b[:n] = ""
