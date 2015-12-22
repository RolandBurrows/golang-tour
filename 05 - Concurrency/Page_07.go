// https://tour.golang.org/concurrency/7
// https://tour.golang.org/concurrency/8

// Exercise: Equivalent Binary Trees

package main

import (
        "fmt"
        "golang.org/x/tour/tree"
)

func walkImpl(t *tree.Tree, ch chan int) {
        if t == nil {
                return
        }
        walkImpl(t.Left, ch)
        ch <- t.Value
        walkImpl(t.Right, ch)
}

func Walk(t *tree.Tree, ch chan int) {
        walkImpl(t, ch)
        close(ch)
}

func Same(t1, t2 *tree.Tree) bool {
        w1, w2 := make(chan int), make(chan int)

        go Walk(t1, w1)
        go Walk(t2, w2)

        for {
                v1, ok1 := <-w1
                v2, ok2 := <-w2
                if !ok1 || !ok2 {
                        return ok1 == ok2
                }
                if v1 != v2 {
                        return false
                }
        }
}

func main() {
        fmt.Print("tree.New(1) == tree.New(1): ")
        if Same(tree.New(1), tree.New(1)) {
                fmt.Println("PASSED")
        } else {
                fmt.Println("FAILED")
        }

        fmt.Print("tree.New(1) != tree.New(2): ")
        if !Same(tree.New(1), tree.New(2)) {
                fmt.Println("PASSED")
        } else {
                fmt.Println("FAILED")
        }
}


// -- Results --

// tree.New(1) == tree.New(1): PASSED
// tree.New(1) != tree.New(2): PASSED

// Program exited.
