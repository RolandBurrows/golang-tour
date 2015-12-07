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

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)							// Failure point
}


// -- Results --

// prog.go:16: cannot use &r (type *rot13Reader) as type io.Reader in argument to io.Copy:
// 	*rot13Reader does not implement io.Reader (missing Read method)

// Program exited.
