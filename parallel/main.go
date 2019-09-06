package main

import (
    "fmt"
    "runtime"
    "sync"
)

func work(n int, wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < 10; i++ {
        var k int = 0
        for j := 0; j < 1e8; j++ {
            k += 1
        }
        fmt.Println(n, i, k)
    }
}

func main() {
    fmt.Println("Runtime    :", runtime.Version())
    fmt.Println("NumCPU     :", runtime.NumCPU())
    fmt.Println("GOMAXPROCS :", runtime.GOMAXPROCS(0))
    var n int = runtime.NumCPU() + 1
    var wg sync.WaitGroup
    wg.Add(n)
    for i := 0; i < n; i++ {
        go work(i+1, &wg)
    }
    wg.Wait()
}
