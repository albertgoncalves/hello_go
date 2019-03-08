package main

import (
    "fmt"
    "math/rand"
    "time"
)

func pipe(x int, fs ...func(int) int) int {
    for _, f := range fs {
        fmt.Println(x)
        x = f(x)
    }
    return x
}

func add(a int) func(int) int {
    var f = func(b int) int {
        return a + b
    }
    return f
}

func gen(r *rand.Rand, k int) int {
    return r.Intn(k)
}

func main() {
    var seed int64 = time.Now().UnixNano()
    var r *rand.Rand = rand.New(rand.NewSource(seed))
    const a, b, c, d = 100, 200, 300, 400
    y := pipe(
        gen(r, a),
        add(a),
        add(b),
        add(c),
        add(d),
    )
    fmt.Println(y)
}
