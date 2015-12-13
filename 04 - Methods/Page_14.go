// https://tour.golang.org/methods/14

package main

import (
	"log"
	"net/http"
)

func main() {
	// your http.Handle calls here
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}


// -- Results --

// // view-source:http://localhost:4000/
// 404 page not found
