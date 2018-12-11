package main

import (
    "fmt"
    "math/rand"
    "time"
)

func add(a int) func(int) int {
    f := func(b int) int {
        return a + b
    }
    return f
}

func main() {
    seed := time.Now().UnixNano()
    r := rand.New(rand.NewSource(seed))

    gen := func() int {
        return r.Intn(100)
    }

    a := gen()
    b := gen()

    fmt.Println(add(a)(b))
}
