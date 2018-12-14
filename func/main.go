package main

import (
    "fmt"
    "math/rand"
    "time"
)

func add(a int) func(int) int {
    var f = func(b int) int {
        return a + b
    }
    return f
}

var gen = func(r *rand.Rand, k int) int {
    return r.Intn(k)
}

func main() {
    var seed int64 = time.Now().UnixNano()
    var r *rand.Rand = rand.New(rand.NewSource(seed))

    const a int = 100
    var b int = gen(r, a)
    var c int = gen(r, b)
    var d int = add(a)(b)
    var e int = add(c)(d)

    fmt.Println(e)
}
