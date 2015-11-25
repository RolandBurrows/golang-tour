// https://tour.golang.org/methods/7

package main

import "fmt"

type IPAddr [4]byte

func (p IPAddr) String() string {
    return fmt.Sprintf("%v.%v.%v.%v",p[0], p[1], p[2], p[3])
}

func main() {
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a)
	}
}


// -- Results --

// loopback: 127.0.0.1
// googleDNS: 8.8.8.8

// Program exited.
