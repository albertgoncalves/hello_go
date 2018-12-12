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
    var seed int64 = time.Now().UnixNano()
    r := rand.New(rand.NewSource(seed))

    gen := func(k int) int {
        return r.Intn(k)
    }

    const a int = 100
    b := gen(a)

    fmt.Println(add(a)(b))
}
