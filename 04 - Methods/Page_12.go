// https://tour.golang.org/methods/12

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rotate *rot13Reader) Read(input []byte) (n int, err error) {
	n,err = rotate.r.Read(input)
	for i := 0; i < len(input); i++ {
		if (input[i] >= 'A' && input[i] < 'N') || (input[i] >='a' && input[i] < 'n') {
			input[i] += 13
		} else if (input[i] > 'M' && input[i] <= 'Z') || (input[i] > 'm' && input[i] <= 'z'){
			input[i] -= 13
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}


// -- Results --

// You cracked the code!
// Program exited.
