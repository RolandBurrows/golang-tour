// https://tour.golang.org/methods/14

package main

import (
        "fmt"
        "log"
        "net/http"
)

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, s)
}

type Struct struct {
        Greeting string
        Punct    string
        Who      string
}

func (s *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "%s%s %s", s.Greeting, s.Punct, s.Who)
}

func main() {
        http.Handle("/string", String("I'm a frayed knot."))
        http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
        err := http.ListenAndServe("localhost:4000", nil)
        if err != nil {
                log.Fatal(err)
        }
}  


// -- Results --

// // view-source:http://localhost:4000
// 404 page not found

// // view-source:http://localhost:4000/string
// I'm a frayed knot.

// // view-source:http://localhost:4000/struct
// {Hello : Gophers!}

