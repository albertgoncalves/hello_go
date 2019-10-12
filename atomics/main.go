package main

import (
    "fmt"
    "sync"
    "sync/atomic"
    "time"
)

var GLOBAL uint32 = 0

func work(value uint32, wg *sync.WaitGroup) {
    time.Sleep(1 * time.Second)
    atomic.AddUint32(&GLOBAL, value)
    wg.Done()
}

func main() {
    var wg sync.WaitGroup
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        value := uint32(i)
        go work(value, &wg)
    }
    wg.Wait()
    fmt.Println(GLOBAL)
}
